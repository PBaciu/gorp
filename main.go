package main

import (
  "net/http"
  "fmt"
  "net"
)

func main()  {
  go http.ListenAndServe(":5001", http.HandlerFunc(fiveThousandHandler))
  http.ListenAndServe(":5002", http.HandlerFunc(fiveThousandHandler))
}

func fiveThousandHandler(w http.ResponseWriter, req *http.Request) {
  port := identifyPort(req)
  fmt.Printf("Hello From Port %d\n", port)
}

func sixThousandHandler() {

}

func sevenThousandHandler() {

}

func identifyPort(req *http.Request) uint16 {
  context := req.Context()
  
  addr, ok := context.Value(http.LocalAddrContextKey).(net.Addr)
  if !ok {
    fmt.Println("ERROR: request address not found.")
  }

  port := addr.(*net.TCPAddr).Port
 
  return uint16(port)
}
