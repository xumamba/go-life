package onet

import (
	"net/url"
)

type IHTTPClient interface {
	Get(url string) (statusCode int, respBody []byte, err error)
	Post(url string, data []byte) (statusCode int, respBody []byte, err error)
	PostForm(url string, data url.Values) (statusCode int, respBody []byte, err error)
	PostJson(url string, data []byte) (statusCode int, respBody []byte, err error)
	PostJsonObj(url string, req, resp interface{}) (err error)
}
