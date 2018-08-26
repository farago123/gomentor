package main
import (
  "net/http"
  "io/ioutil"
  "strings"
  "fmt"
  "encoding/json"
  "os/exec"
)

type Message struct {
    Body string
}

func fullText(w http.ResponseWriter, r *http.Request) {

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

    matchingLines := []string{}

    for _, item := range temp {
      if strings.Contains(item, word) {
         matchingLines = append(matchingLines, item)
      }
    }

    b, err := json.Marshal(matchingLines)
    if err != nil {
        fmt.Println(err)
    }
    w.Write(b)
 
}

func nlp(w http.ResponseWriter, r *http.Request) {

    data, err := ioutil.ReadFile("text.txt")

    if err != nil {
      fmt.Println(err)
    }
    
    file := string(data)
    temp := strings.Split(file, ". ")
    firstSentence := temp[0]

    m := Message{"First Sentence: " + firstSentence}
    b, err := json.Marshal(m)
    if err != nil {
        fmt.Println(err)
    }
    w.Write(b)
    w.Write([]byte("\n"))

    for _, item := range temp[1:] {
      
        input1 := strings.Replace(item, "\n", "", -1)
        input2 := strings.Replace(firstSentence, "\n", "", -1)
        s1 := strings.Split(computeSemanticSimilarity(input1, input2), "\n")
        m := Message{item + " - Semantic Similarity to First Sentence: " + s1[6]}
        b, err := json.Marshal(m)
        if err != nil {
            fmt.Println(err)
        }
        w.Write([]byte("\n"))
        w.Write(b)

    }

}

func computeSemanticSimilarity(sentence1 string, sentence2 string) string {

    cmd := exec.Command("python",  "-c", "import SentenceSimilarity; print SentenceSimilarity.getSemanticSimilarity('" + sentence1 + "', '" + sentence2 + "')")
    out, err := cmd.CombinedOutput()

    if err != nil { fmt.Println(err); }   

    return string(out)

}

func main() {

  http.HandleFunc("/nlp", nlp)
  http.HandleFunc("/search", search)
  http.HandleFunc("/hi", sayHi)
  http.HandleFunc("/", fullText)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
