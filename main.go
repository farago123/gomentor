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
      temp := strings.Split(file, "\n")

      for _, item := range temp {
           m := Message{item}
           b, err := json.Marshal(m)
           if err != nil {
            fmt.Println(err)
           }
           w.Write(b)

           // w.Write([]byte(item))
      }
      
  }

}

func sayHi(w http.ResponseWriter, r *http.Request) {

    // u, err := url.Parse("https://example.org" + r.URL.String())

    // if err != nil {
    //   fmt.Println(err)
    // }

    // q := u.Query()
    // w.Write([]byte("hi " + q["firstname"][0] + " " + q["lastname"][0]))

    firstName := r.URL.Query()["firstname"][0]
    lastName := r.URL.Query()["lastname"][0]

    // w.Write([]byte("hi " + string(firstName) + " " + string(lastName)))

    m := Message{"hi " + firstName + " " + lastName}
    b, err := json.Marshal(m)
    if err != nil {
      fmt.Println(err)
    }
    w.Write(b)

}

func search(w http.ResponseWriter, r *http.Request) {

    // word := strings.TrimPrefix(r.URL.String(), "/search?")

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
         // w.Write([]byte(item))
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
