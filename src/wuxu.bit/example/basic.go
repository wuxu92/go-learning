package main

import (
  "fmt"
  "time"
  "math"
  "math/rand"
  "strconv"
  "sort"
  // "strconv"
)

func values() {
  _ = fmt.Println
  fmt.Println("values funciton")
}

func variables(){
  var a string = "initial"
  b := "initailed"
  var d = true

  var i interface{}
  i = nil

  fmt.Println(a, b, d, i)
}

func constants(){
  // 声明const变量不需要:=
  const n = 1013
  fmt.Println(n)

  const e = 3e5/n
  fmt.Println(e)

  // 类型转换
  // 不能直接把const类型转到到其他类型，需要新建一个变量来转换
  fmt.Println((float64(e)))
  var xi float64 = e;
  fmt.Println(int64(xi))

  fmt.Println(math.Sin(n))
}

func forThins() {
  i := 0
  for i<3 {
    fmt.Println(i)
    i = i+1
  }
}

func ifElse() {
  
}

func switchThings() {
  // switch 没有fallthrough
  i := 2
  switch i {
    case 1: fmt.Println("i is 1")
    case 2: fmt.Println("i is 2")
    case 3: fmt.Println("i is 3")
    default: fmt.Println("i is default")
  }

  // 一个case可以使用,分隔多个条件
  t := time.Now().Weekday()
  switch t {
    case time.Saturday, time.Sunday: fmt.Println("today is weekend")
    default: fmt.Println("today is weekday")
  }
}


func arrays() {
   var a [5]int    // 声明一个数组
   fmt.Println("array", a)

   a[3] = 1;
   fmt.Println("set 3", a)

   // 使用字面量声明数组，不需要显示声明长度
   b := []int{1,2,3,4}
   fmt.Println("b:", b)

   // 二维数组
   var c [2][3]int
   for i:=0;i<2;i++ {
     for j:=0; j<3; j++ {
      c[i][j] = i+j
     }
   }
   fmt.Println("two dimension:", c)
}

func slices() {
  // slice 使用make初始化
  s := make([]string, 3)
  fmt.Println("empty slice: ", s)

  // 想数组一样使用切片
  s[0] = "999"
  s[2] = "ooo"
  fmt.Println("set slice: ", s, " at len;", len(s))

  // 切片操作
  sc := s[0:1]  // 获得第一个元素
  fmt.Println("slice 0 1:", sc)

  // 二维非定长切片
  d2s := make([][]int, 4)
  for i:=0; i<4; i++ {
    d2s[i] = make([]int, i+1)
    for j:=0; j<i+1; j++ {
      d2s[i][j] = i+j
    }
  }
  fmt.Println("d2s", d2s)
}


func maps() {
  // map是关联数据类型
  // slice没有专门的关键字，但是map是使用map关键字的 
  m := make(map[string]string)

  m["today"] = "sunday"
  m["tomorrow"] = "monday"
  fmt.Println(m)

  delete(m, "yesterday")

  nm := map[string]string{"today": "haha", "yesterday": "nice"}
  _, pr := nm["notExist"]
  fmt.Println("select a not exist index", pr)
}


func rangeThings() {
  // range 用于slice和map的遍历
}

type Person struct {
  name string
  age int64
  sex bool
}

func structs() {
  fmt.Println(Person{"wuxu", 22, true});

  fmt.Println(Person{name: "wuxu", age: 22, sex: true})

  me := Person{"wuxu", 22, true}
  fmt.Println("i am "+ me.name)

  // 结构体指针
  mePtr := &me
  fmt.Println("i am " , mePtr.age , " years old")
  mePtr.age = 23
  fmt.Println("i will be " , (mePtr.age) , " years old next year")

  me.age = 35;
  fmt.Println("new age", mePtr.age)
  // wrong: mePtr = Person("uu", 22, true)
  // follow is right
  *mePtr = Person{"wuxx", 33, true}
  fmt.Println(me.name)
}

// method part, add method to struct
// @see https://gobyexample.com/methods
func (p Person) say(words string) {
  fmt.Println(p.name, " says: ", words)
  p.age = p.age+1
  fmt.Println("age after say: ", p.age)
}

func (p Person) run(dist int) {
  fmt.Println(p.name, " runs away for ", dist, " m")
  p.name = p.name+ "run"
}

func methods() {
  me := Person{"wuxu", 22, true}
  me.say(" wtf! ")
  me.run(100)
  fmt.Println(me.name, me.age)
}

// interface part
type animal interface {
  run(int)
}

type Pig struct {
  name string
  age int
}

func (p Pig) run(d int) {
  fmt.Println("Pig runs away ", d)
}

func runAway(an animal) {
  an.run(100);
}

func inters() {
  runAway(Person{"TM", 22, false})
  runAway(Pig{"wangcai", 1})
}

// errors
type myErr struct {
  arg int
  prob string
}

func (e myErr) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func (p Person)buy(item string) (bool, error) {
  if p.age < 18 {
    return false, &myErr{0, "too yong"}
  }
  return true, nil
}

func errors() {
  res,e := Person{"wuxu", 14, true}.buy("shirt")
  if me,ok := e.(*myErr); ok {
    fmt.Println("error occurs", me.Error())
  }
  fmt.Println("result", res)
}

//////////////////////////////////
//   goroutines
//   channel
//////////////////////////////////

func channels() {
  msgChan := make(chan string, 2)
  msgChan <- "hello"
  msgChan <- "world"

  msg := <-msgChan
  fmt.Println("message is ", msg)
  fmt.Println("second msg", <-msgChan)
}

func working(flagChan chan bool) {
  fmt.Println("working on it")
  time.Sleep(2 * time.Second)
  fmt.Println("work done")

  flagChan <- true
}

func chanSync() {
  flagChan := make(chan bool, 1)
  go working(flagChan)

  done := <- flagChan
  fmt.Println("sync goroutine output", done)
}

// select part
func gr1(c chan string) {
  dur := rand.Intn(300)
  time.Sleep(time.Millisecond * time.Duration(dur))
  c <- "gr1 sleep done " + strconv.Itoa(dur)
}

func gr2(c chan string) {
  dur := rand.Intn(400)
  time.Sleep(time.Millisecond * time.Duration(dur))
  c <- "gr2 sleep done " + strconv.Itoa(dur)
}

func selects() {
  c1 := make(chan string)
  c2 := make(chan string)
  go gr1(c1);
  go gr1(c1);
  go gr2(c2);
  go gr2(c2);

  for i:=0; i<4; i++ {
    select {
    case msg := <-c1:
      fmt.Println(msg)
    case msg2 := <-c2:
      fmt.Println(msg2)
    // default:
      // non-blocking channel
    //    fmt.Println("select default")
    }
  }
}

func rangeChannel() {
  c := make(chan string, 2)
  c <- "tomorrow is monday"
  c <- "fee out"
  close(c)
  for str := range c {
    fmt.Println(str)
  }
}

func sorts() {
  strs := []string{"today", "tomorrow", "yesterday", "am who", "make"}
  sort.Strings(strs)
  fmt.Println("sorted strings:", strs)
  fmt.Println(sort.StringsAreSorted(strs))
}


