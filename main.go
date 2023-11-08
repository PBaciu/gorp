package main

import (
  "net/http"
  "fmt"
  "net"
  "slices"
)

var registry = make(map[uint16][]string)

func main()  {
  registerHost(5001, "https://google.com")
  registerHost(5001, "https://cnn.com")
  
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

// Registers host as a potential target for incoming request from port.
func registerHost(port uint16, host string) {
  if !slices.Contains(registry[port], host) {
    registry[port] = append(registry[port], host)
  }
}
