package gale

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type Req struct {
	context *fasthttp.RequestCtx
	body    []byte
}

type Body struct {
}

func (r Req) Send(s string) {
	r.context.SetContentType("text/plain; charset=utf-8")
	r.context.SetStatusCode(fasthttp.StatusOK)
	r.context.WriteString(s)
}

func (r Req) ByteBody() []byte {
	return r.body
}

func (r Req) StringBody() string {
	return string(r.body)
}

func (r Req) JsonBody() (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(r.body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling")
		return nil, err
	}
	return data, nil
}
