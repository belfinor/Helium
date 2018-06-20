package jsonrpc2

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06

import (
	"encoding/json"
	"github.com/belfinor/Helium/log"
	"reflect"
)

var callback map[string]interface{} = map[string]interface{}{}

func RegisterMethod(method string, fn interface{}) {
	callback[method] = fn
}

func MakeError(rec *Request, code int64, message string) *Response {
	return &Response{Id: rec.Id, JsonRPC: "2.0", Error: &Error{Code: code, Message: message}}
}

func Exec(rec *Request) *Response {
	if rec.JsonRPC != "2.0" {
		log.Error("invalid protocol version: " + rec.JsonRPC)
		if rec.Id != nil {
			return MakeError(rec, -32600, "Invalid Request")
		}
		return nil
	}

	fn, has_fn := callback[rec.Method]
	if !has_fn {
		log.Error("method not found: " + rec.Method)
		if rec.Id != nil {
			return MakeError(rec, -32601, "Method not found")
		}
		return nil
	}

	if reflect.TypeOf(fn).NumIn() != 1 {
		log.Error("inavlid number of arguments in handler function")
		if rec.Id != nil {
			return MakeError(rec, -32603, "Internal error")
		}
		return nil
	}

	arg_type := reflect.TypeOf(fn).In(0)
	arg := reflect.New(arg_type)
	arg_i := arg.Interface()

	data, _ := rec.Params.MarshalJSON()

	if err := json.Unmarshal(data, &arg_i); err != nil {
		log.Error("invalid params")
		if rec.Id != nil {
			return MakeError(rec, -32602, "Invalid params")
		}
		return nil

	}

	rets := reflect.ValueOf(fn).Call([]reflect.Value{arg.Elem()})

	if rec.Id == nil {
		return nil
	}

	if !rets[1].IsNil() {
		return &Response{Id: rec.Id, JsonRPC: rec.JsonRPC, Error: rets[1].Interface()}
	}

	return &Response{Id: rec.Id, JsonRPC: rec.JsonRPC, Result: rets[0].Interface()}
}
