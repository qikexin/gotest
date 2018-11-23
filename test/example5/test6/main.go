package test6

import (
	"net/http"
)

type Request struct {
	//HTTP请求
	httpReq *http.Request
	//请求的深度
	depth uint32
}
//用于创建一个新的请求实例
func NewRequest(httpReq *http.Request,depth uint32) *Request {
	return &Request{httpReq: httpReq,depth: depth}
}
//用于获取HTTP请求
func (req *Request) HTTPReq() *http.Request{
	return req.httpReq
}
//用于获取请求的深度
func (req *Request) Depth() uint32{
	return req.depth
}

type Response struct {
	httpResp *http.Response
	depth uint32
}
//创建一个新的响应实例
func NewResponse(httpResp *http.Response,depth uint32) *Response{
	return &Response{httpResp: httpResp,depth: depth}
}

func (resp *Response) HTTPResp() *http.Response{
	return resp.httpResp
}
func (resp *Response)Depth() uint32 {
	return resp.depth
}

type Item map[string]interface{}

type Data interface {
	//用于判断数据是否有效
	Valid() bool
}

//判断请求是否有效
func (req *Request) Valied() bool  {
	return req.httpReq != nil &&req.httpReq.URL != nil
}

//用于判断响应是否有效
func (resp *Response) Valied() bool{
	return resp.httpResp !=nil && resp.httpResp.Body != nil
}

//用于判断条目是否有效
func (item Item) Valied()bool {
	return item != nil
}

type CrawlerError interface {
	Type() ErrorType
	Error() string
}