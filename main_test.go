package main
import (
    "net/http"
    "testing"
    "net/http/httptest"
    "fmt"
    "io/ioutil"
    "log"
)

func TestFullTextHandler(t *testing.T) {
    
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(fullText)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Body":"As of 2017, text messages are used by youth and adults for personal, family and social purposes and in\nbusiness. Governmental and non-governmental organizations use text messaging for communication between\ncolleagues. As with emailing, in the 2010s, the sending of short informal messages has become an accepted\npart of many cultures.[1] This makes texting a quick and easy way to communicate with friends and colleagues,\nincluding in contexts where a call would be impolite or inappropriate (e.g., calling very late at night or\nwhen one knows the other person is busy with family or work activities). Like e-mail and voice mail, and\nunlike calls (in which the caller hopes to speak directly with the recipient), texting does not require the\ncaller and recipient to both be free at the same moment; this permits communication even between busy\nindividuals. Text messages can also be used to interact with automated systems, for example, to order\nproducts or services from e-commerce websites, or to participate in online contests. Advertisers and service\nproviders use direct text marketing to send messages to mobile users about promotions, payment due dates,\nand other notifications instead of using postal mail, email, or voicemail."}`

    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSayHiHandler1(t *testing.T) {
    
    req, err := http.NewRequest("GET", "http://localhost:8080/hi?firstname=joe&lastname=ellis", nil)

    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(sayHi)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Body":"hi joe ellis"}`

    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSayHiHandler2(t *testing.T) {
    
    req, err := http.NewRequest("GET", "http://localhost:8080/hi?firstname=peter&lastname=farago", nil)

    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(sayHi)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Body":"hi peter farago"}`

    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSayHiHandler3(t *testing.T) {
    
    req, err := http.NewRequest("GET", "http://localhost:8080/hi?firstname=jared&lastname=scheib", nil)

    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(sayHi)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Body":"hi jared scheib"}`

    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSearchHandler1(t *testing.T) {
    
    req, err := http.NewRequest("GET", "http://localhost:8080/search?the", nil)

    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(search)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Body":"colleagues. As with emailing, in the 2010s, the sending of short informal messages has become an accepted"}{"Body":"when one knows the other person is busy with family or work activities). Like e-mail and voice mail, and"}{"Body":"unlike calls (in which the caller hopes to speak directly with the recipient), texting does not require the"}{"Body":"caller and recipient to both be free at the same moment; this permits communication even between busy"}{"Body":"and other notifications instead of using postal mail, email, or voicemail."}`

    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSearchHandler2(t *testing.T) {
    
    req, err := http.NewRequest("GET", "http://localhost:8080/search?where", nil)

    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(search)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Body":"including in contexts where a call would be impolite or inappropriate (e.g., calling very late at night or"}`
    
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSearchHandler3(t *testing.T) {
    
    req, err := http.NewRequest("GET", "http://localhost:8080/search?call", nil)

    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(search)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Body":"including in contexts where a call would be impolite or inappropriate (e.g., calling very late at night or"}{"Body":"unlike calls (in which the caller hopes to speak directly with the recipient), texting does not require the"}{"Body":"caller and recipient to both be free at the same moment; this permits communication even between busy"}`
    
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func IntegrationTestServer(t *testing.T) {

    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, client")
    }))
    defer ts.Close()

    res, err := http.Get(ts.URL)

    if err != nil {
        log.Fatal(err)
    }

    greeting, err := ioutil.ReadAll(res.Body)
    res.Body.Close()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s", greeting)

}

func IntegrationTestFullText(t *testing.T) {

    ts := httptest.NewServer(http.HandlerFunc(fullText))
    defer ts.Close()

    res, err := http.Get(ts.URL)

    if err != nil {
        log.Fatal(err)
    }

    text, err := ioutil.ReadAll(res.Body)
    res.Body.Close()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s", text)

}

func IntegrationTestSayHi(t *testing.T) {

    ts := httptest.NewServer(http.HandlerFunc(sayHi))
    defer ts.Close()

    res, err := http.Get(ts.URL)

    if err != nil {
        log.Fatal(err)
    }

    text, err := ioutil.ReadAll(res.Body)
    res.Body.Close()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s", text)

}

func IntegrationTestSearch(t *testing.T) {

    ts := httptest.NewServer(http.HandlerFunc(search))
    defer ts.Close()

    res, err := http.Get(ts.URL)

    if err != nil {
        log.Fatal(err)
    }

    text, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s", text)

}

