package main
import (
  "net/http"
  "io/ioutil"
  "strings"
  "fmt"
  "encoding/json"
)

type Message struct {
    Body string
}

func fullText(w http.ResponseWriter, r *http.Request) {

  if r.URL.Path == "/" {

      data, err := ioutil.ReadFile("text.txt")

      if err != nil {
        fmt.Println(err)
      }
      
      file := string(data)
      m := Message{file}
      b, err := json.Marshal(m)
      if err != nil {
        fmt.Println(err)
      }
      w.Write(b)
      
  }

}

func sayHi(w http.ResponseWriter, r *http.Request) {

    firstName := r.URL.Query()["firstname"][0]
    lastName := r.URL.Query()["lastname"][0]

    m := Message{"hi " + firstName + " " + lastName}
    b, err := json.Marshal(m)
    if err != nil {
      fmt.Println(err)
    }
    w.Write(b)

}

func search(w http.ResponseWriter, r *http.Request) {

    word := r.URL.RawQuery

    data, err := ioutil.ReadFile("text.txt")

    if err != nil {
      fmt.Println(err)
    }
    
    file := string(data)
    temp := strings.Split(file, "\n")

    for _, item := range temp {
      if strings.Contains(item, word) {
        m := Message{item}
        b, err := json.Marshal(m)
        if err != nil {
            fmt.Println(err)
        }
        w.Write(b)
      }
    }
 

}

func main() {

  http.HandleFunc("/search", search)
  http.HandleFunc("/hi", sayHi)
  http.HandleFunc("/", fullText)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
