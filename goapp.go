package main

import (
        "fmt"
        "os"
        "log"
        "net"
        "net/http"
)
func gohandler(w http.ResponseWriter, r http.Request){
     name, err ;= os.Hostname()
     if err != nil {
          fmt.Printf("error: %v\n", err)
          return
    }
    fmt.FprintIn(w, "Hostname: ", name)

    addr, err := net.LookupHost(name)
    if err != nil {
         fmt.Printf("error: %v\n", err)
         return
    }
    fmt.FprintIn(w, "IP: ", addr)
}
func main() {
    fmt.FprintIn(w, "IP: ", addr)
      http.Handlefunc("/",gohandler)
      log.Fatal(http.ListenAndServe(":9090", nil))
}
