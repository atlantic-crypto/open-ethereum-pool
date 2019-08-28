package proxy

import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type JSONRpcReq struct {
	Id     json.RawMessage `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type StratumReq struct {
	JSONRpcReq
	Worker string `json:"worker"`
}

// Stratum
type JSONPushMessage struct {
	// FIXME: Temporarily add ID for Claymore compliance
	Id      int64       `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

type JSONRpcResp struct {
	Id      json.RawMessage `json:"id"`
	Version string          `json:"jsonrpc"`
	Result  interface{}     `json:"result"`
	Error   interface{}     `json:"error,omitempty"`
}

type SubmitReply struct {
	Status string `json:"status"`
}

type ErrorReply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
