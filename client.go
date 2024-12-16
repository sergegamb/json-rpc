package main

import (
 "bytes"
 "fmt"
 "io"
 "net/http"

 "github.com/tidwall/gjson"
 "github.com/tidwall/sjson"
)

func main() {
 // Create JSON-RPC request using sjson
 request := "{}"
 request, _ = sjson.Set(request, "jsonrpc", "2.0")
 request, _ = sjson.Set(request, "method", "subtract")
 request, _ = sjson.Set(request, "params.minuend", 10)
 request, _ = sjson.Set(request, "params.subtrahend", 4)
 request, _ = sjson.Set(request, "id", 1)

 // Send HTTP POST request
 resp, err := http.Post("http://localhost:8080/rpc", "application/json", bytes.NewBufferString(request))
 if err != nil {
  fmt.Println("Error making request:", err)
  return
 }
 defer resp.Body.Close()

 // Read the response body
 body, err := io.ReadAll(resp.Body)
 if err != nil {
  fmt.Println("Error reading response:", err)
  return
 }

 // Parse the result using gjson
 result := gjson.GetBytes(body, "result")
 fmt.Println("Result:", result)
}
