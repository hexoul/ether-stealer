// Package json is a JSON parser for request & response body used in JSON-RPC
package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RPCRequest is a interface for JSON-RPC request
type RPCRequest struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int32    `json:"id"`
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

var (
	errGeneral = fmt.Errorf("Failed")
)

// GetRPCRequestFromRaw returns RPCRequest struct from raw byte array
func GetRPCRequestFromRaw(msg []byte) RPCRequest {
	var data RPCRequest
	json.Unmarshal(msg, &data)
	return data
}

func (r *RPCRequest) Byte() []byte {
	if ret, err := json.Marshal(r); err == nil {
		return ret
	}
	return nil
}

// GetRPCResponse returns RPCResponse through RPCRequest
func GetRPCResponse(url string, rpcRequest RPCRequest) (*RPCResponse, error) {
	resp, err := http.Post(url, "application/json", bytes.NewReader(rpcRequest.Byte()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := GetRPCResponseFromRaw(respBody)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRPCResponseFromRaw returns RPCResponse from raw byte array
func GetRPCResponseFromRaw(msg []byte) (*RPCResponse, error) {
	data := RPCResponse{}
	err := json.Unmarshal(msg, &data)
	if err == nil {
		return &data, nil
	}
	return nil, err
}

func (r *RPCResponse) String() string {
	if ret, err := json.Marshal(r); err == nil {
		return string(ret)
	}
	return ""
}
