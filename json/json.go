// Package json is a JSON parser for request & response body used in JSON-RPC
package json

import (
	"encoding/json"
)

// RPCRequest is a interface for JSON-RPC request
type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int32         `json:"id"`
}

// RPCError is a interface for JSON-RPC error
type RPCError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// RPCResponse is a interface for JSON-RPC response
type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int32       `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
}

// GetRPCRequestFromJSON returns RPCRequest struct from JSON
func GetRPCRequestFromJSON(msg []byte) RPCRequest {
	var data RPCRequest
	json.Unmarshal(msg, &data)
	return data
}

func (r *RPCRequest) String() string {
	if ret, err := json.Marshal(r); err == nil {
		return string(ret)
	}
	return ""
}

// GetRPCResponseFromJSON returns RPCRequest struct from JSON
func GetRPCResponseFromJSON(msg []byte) RPCResponse {
	var data RPCResponse
	json.Unmarshal(msg, &data)
	return data
}

func (r *RPCResponse) String() string {
	if ret, err := json.Marshal(r); err == nil {
		return string(ret)
	}
	return ""
}
