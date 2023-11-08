package main

import (
  "net/http"
  "fmt"
  "net"
  "slices"
)

// TODO consider using set instead of slices
var registry = make(map[uint16][]string)

func main()  {
  initializeRegistry()

  go http.ListenAndServe(":5001", http.HandlerFunc(requestHandler))
  http.ListenAndServe(":5002", http.HandlerFunc(requestHandler))
}

func initializeRegistry() {
  registerHost(5001, "https://google.com")
  registerHost(5001, "https://cnn.com")
  registerHost(5002, "https://google.com")
  registerHost(5002, "https://bloomberg.com")
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
  port := identifyPort(req)

  // TODO add some logic to forward a request to a host given host availability
  candidateHosts := registry[port]
  if len(candidateHosts) == 0 {
    fmt.Println("ERROR: no candidate hosts for port")
  }
  fmt.Printf("Candidates for port %d: %v\n", port, candidateHosts)
}

// Identifies the port of the incoming request
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
