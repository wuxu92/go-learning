package main

import (
  "fmt"
  "time"
  "math/rand"
  "sync/atomic"
  "sync"
  "runtime"
)

type iMilk interface {
  getName() string
}

type Milk struct {
  name string
}

func (m Milk) getName() string {
  return m.name
}

type Logger struct {
  content string
}

func (l Logger) Log(s string) {
  l.content += s
  fmt.Println(s)
}

type infoLog struct {
  count int
  *Logger
}

func main() {
  var m interface {}

  m = Milk{"yili"}
  if _,ok := m.(iMilk); ok {
    fmt.Printf("value %v has implement iMilk\n ", m)
  } else {
    fmt.Println("has not implememt")
  }

  var logger = infoLog{0, &Logger{"sss"} }
  logger.Log("embedding function invoke")

  t1 := time.NewTimer(time.Second * 2)

  go func() {
    <- t1.C
    fmt.Println("t1 expired? yes")
  }()
  st := t1.Stop()
  if st {
    fmt.Println("t1 has stopped, will not expired")
  }

  t2 := time.NewTicker(time.Millisecond * 500)
  go func() {
    for t:= range t2.C {
      fmt.Println("tick tick at ", t)
    }
  }()
  time.Sleep(time.Second * 3)
  t2.Stop()
  fmt.Println("tick tick stoped")
  

  // go worker pool
  jobs := make(chan int, 100)
  result := make(chan int, 100)

  for w:=0; w<3; w++ {
    go func(w int, jobs <-chan int, result chan<- int) {
      for j:= range jobs {
        fmt.Println("worker", w, "processing on", j)
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(600)))
        result <- j * 2
      }
    }(w, jobs, result)
  }

  // insert job
  for j:=0; j<10; j++ {
    jobs <- j
  }
  close(jobs)
  // do some others
  // sync
  for a :=0; a<10; a++ {
    <- result
  }

  // rate-limiting
  burstChan := make(chan time.Time, 3)
  for i:=0; i<3; i++ {
    burstChan <- time.Now()
  }

  // send value per 200 millisecond
  go func() {
    for t := range time.Tick(time.Millisecond * 200){
      burstChan <- t
    }
  }()

  burstReq := make(chan int, 5)
  for i:=0; i< 5; i++ {
    burstReq <- i
  }
  close(burstReq)
  for r:= range burstReq {
    ht := <- burstChan
    fmt.Println("req", r, "handle at", ht)
  }


  // atomic counters
  var counter uint64= 0;
  for ci:=0; ci<50; ci++ {
    go func() {
      for {
        atomic.AddUint64(&counter, 1)
        cr := atomic.LoadUint64(&counter)
        if cr > 1000000 {
          break
        }
        runtime.Gosched()
      }
    }()
  }
  time.Sleep(time.Second * 1)
  cResult := atomic.LoadUint64(&counter)
  fmt.Println("counter result is ", cResult)

  cResult = atomic.LoadUint64(&counter)
  fmt.Println("counter result is ", cResult)

  _ = sync.Mutex{}
  
  // goroutine stateful
  type read struct {
    key int
    req chan int
  }
  type write struct {
    key int
    value int
    req chan bool
  }
  // shared channels
  reads := make(chan *read)
  writes := make(chan *write)

  // state manage gr
  go func() {
    state := make(map[int]int)
    for {
      select {
        case r:= <-reads:
          r.req <- state[r.key]
        case w:= <-writes:
          state[w.key] = w.value
          w.req <- true
      }
    }
  }()
  var ops int64 = 0
  // read
  for ri := 0; ri<20; ri++ {
    go func() {
      for {
        r := &read {
          key: rand.Intn(10),
          req: make(chan int),
        }
        reads <- r
        res := <-r.req
        _ = res
        atomic.AddInt64(&ops, 1)
      }
    }()
  }
  // write
  for wi:=0; wi<10; wi++ {
    go func() {
      for {
        w := &write{
          key: rand.Intn(10),
          value: rand.Intn(100),
          req: make(chan bool),
        }
        writes <- w
        res := <- w.req
        _ = res
        atomic.AddInt64(&ops, 1)
      }
    }()
  }

  time.Sleep(time.Second * 1)
  opsFinal := atomic.LoadInt64(&ops)
  fmt.Println("total ops:", opsFinal)
}
