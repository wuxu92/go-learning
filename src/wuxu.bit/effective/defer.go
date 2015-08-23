package main

import (
  "fmt"
  "io"
  "os"
  "net/http"
)

var (
  Enon = 2
  Eio = 1
)

type ByteSize float64
const (
  _ = iota
  KB ByteSize = 1 << (1*iota)
  MB
  BG
)

func Contents(filename string) (string, error) {
  f, err := os.Open(filename)
  if err != nil {
    return "", err
  }

  // use defer to close a file
  // you will not forget to close your file
  defer f.Close()
  var result []byte

  buf := make([]byte, 100)
  for {
    n, err := f.Read(buf[0:])
    result = append(result, buf[0:n]...) // what does these 3 dots means
    if err != nil {
      if err == io.EOF {
        break
      }
      return "", err
    }
  }

  return string(result), nil
}

// defer run as LIFO order
func deferRun() {
  for i:=0; i<5; i++ {
    defer fmt.Println("i is ", i)
  }
}

// literal initail a slice
func literal() {
  // a := [...]string {Enon: "no error", Eio:"error io"}
  b := []string {"no error", "error io"}
  m := map[int]string {Enon: "no error", Eio: "error io"}
  fmt.Println( b, m)
}


// array
func arrayCopy(arr []int) []int {
  var newArr = arr;
  for i:=0; i<len(arr); i++ {
    newArr[i]++
  }
  return newArr
}

// twoDimenSlice
func twoDimenSlice() [][]int8 {
  picture := make([][]int8, 10)
  for i := range picture {
    picture[i] = make([]int8, 10)
  }

  picture2 := make([][]int8, 10)
  pixels := make([]int8, 10*10)

  for i:= range picture2 {
    picture2[i], pixels = pixels[:10], pixels[10:]
  }
  fmt.Println(picture, picture2)
  return picture2
}

func mapOp() {
  var weekDay = map[int]string {
    1: "monday",
    2: "tuesday",
    3: "wednesday",
  }

  day2 := weekDay[2]
  fmt.Println(day2)
  day7, ok := weekDay[7]
  fmt.Println(ok, day7)

  // delete
  delete(weekDay, 3)
  // _, exist := weekDay[3]
  fmt.Println(weekDay)

}


func constOp() {
  fmt.Println(KB, MB)
}

var (
  home = os.Getenv("HOME")
  user = os.Getenv("USER")
)
func varOp() {
  fmt.Println(home, user)
}

func init() {
  fmt.Println("init done")
}

func init() {
  f := 0.64
  c := 4
  fmt.Println(f*float64(c))
  fmt.Println("2 init-funcz")
}

type Book struct {
  price float64
}
func (p Book) Discount(c int) float64 {
  return p.price * float64(c)
}

type iMilk interface {
  String() string
  Price() float64
}

type Milk struct {
  name string
  price float64
}

func (m Milk) String() string {
  return m.name
}

func (m Milk) Price() float64 {
  return m.price
}

func interfaceOp() {
  var value = Milk{"helan", 38.0}
  var iv interface{}
  iv = value
  switch str := iv.(type) {
  case string:
    fmt.Println(str)
  case Milk:
    fmt.Println("milk", str)
  default:
    fmt.Println("interface type:", str)
  }

  // fmt.Println(iv)
  if str, ok := iv.(string); ok {
    fmt.Sprint("string is %q\n", str)
  } else if str, ok := iv.(Milk); ok{
    fmt.Sprint("type is: %q, Milk is %q\n", str, iv)
  }
}

// define a handler on func
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  f(w, req)
}

func ArgServer(w http.ResponseWriter, req *http.Request) {
  fmt.Fprint(w, "use func as receiver")
}

// add handle
func addFuncOp() {
  http.Handle("/args", http.HandlerFunc(ArgServer))
  http.ListenAndServe(":7777", nil)
}
