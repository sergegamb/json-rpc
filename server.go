package main

import (
 "fmt"
 "io"
 "net/http"

 "github.com/tidwall/gjson"
 "github.com/tidwall/sjson"
)

func rpcHandler(w http.ResponseWriter, r *http.Request) {
 // Read the request body
 body, err := io.ReadAll(r.Body)
 if err != nil {
  http.Error(w, "Invalid request body", http.StatusBadRequest)
  return
 }

 // Parse JSON-RPC fields
 method := gjson.GetBytes(body, "method").String()
 id := gjson.GetBytes(body, "id").Int()
 params := gjson.GetBytes(body, "params")

 var result interface{}
 var response string

 switch method {
 case "subtract":
  // Extract parameters using gjson
  minuend := params.Get("minuend").Float()
  subtrahend := params.Get("subtrahend").Float()
  result = minuend - subtrahend

  // Construct response using sjson
  response, _ = sjson.Set(response, "jsonrpc", "2.0")
  response, _ = sjson.Set(response, "result", result)
  response, _ = sjson.Set(response, "id", id)

 default:
  // Handle method not found
  response, _ = sjson.Set(response, "jsonrpc", "2.0")
  response, _ = sjson.Set(response, "error", map[string]interface{}{
   "code":    -32601,
   "message": "Method not found",
  })
  response, _ = sjson.Set(response, "id", id)
 }

 // Set the response headers
 w.Header().Set("Content-Type", "application/json")
 w.Write([]byte(response))
}

func main() {
 http.HandleFunc("/rpc", rpcHandler)
 fmt.Println("Server running on http://localhost:8080/rpc")
 http.ListenAndServe(":8080", nil)
}
