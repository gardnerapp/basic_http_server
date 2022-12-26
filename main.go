package main

import (
  "net/http"
  "io"
  "log"
  "fmt"
  "time"
)

func main()  {

  // you can also register handlerfunctions ananoymously
  anon := func(w http.ResponseWriter, r *http.Request)  {
    io.WriteString(w, "Congrats you ust registered an anonymous handler function\n")
    LogReq(r)
  }

  //http.handle func takes the path and the func you
  //want to exe when the path is called
  http.HandleFunc("/", index)
  http.HandleFunc("/meow", meow)
  http.HandleFunc("/anon", anon)
  // ^^ registers w mux

  fmt.Println("Starting HTTP Server !")
  log.Fatal(http.ListenAndServe(":3000", nil)) // nil == server middleware def to mux

}

// create an http handler
func index(w http.ResponseWriter, r *http.Request)  {
  io.WriteString(w, "You just called the index function\n")
  LogReq(r)
}

func meow(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "RaWR XD\n")
  LogReq(r)
}


func LogReq(req *http.Request) {
  fmt.Printf("%s\n", resp)
  fmt.Printf("%s --- %s\n", time.Now(), req.RemoteAddr)
  fmt.Printf("%s %s %s\n", req.Method, req.URL, req.Proto)
  fmt.Println()
}
