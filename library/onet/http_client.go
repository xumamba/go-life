package onet

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

func (hc *HTTPClient) Get(url string) (statusCode int, respBody []byte, err error) {
	response, err := hc.client.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return
	}
	statusCode = response.StatusCode
	respBody, err = ioutil.ReadAll(response.Body)
	return
}

func (hc *HTTPClient) Post(url string, data []byte) (statusCode int, respBody []byte, err error) {
	panic("implement me")
}

func (hc *HTTPClient) PostForm(url string, data url.Values) (statusCode int, respBody []byte, err error) {
	panic("implement me")
}

func (hc *HTTPClient) PostJson(url string, data []byte) (statusCode int, respBody []byte, err error) {
	panic("implement me")
}

func (hc *HTTPClient) PostJsonObj(url string, req, resp interface{}) (err error) {
	panic("implement me")
}

func NewHTTPClient() IHTTPClient {
	httpClient := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   false,
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 1000,
			IdleConnTimeout:     30 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				KeepAlive: 60 * time.Second,
				Timeout:   60 * time.Second,
			}).DialContext,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	return &HTTPClient{client: httpClient}

}

func (hc *HTTPClient) SetTimeout(timeout int) {
	hc.client.Timeout = time.Duration(timeout) * time.Millisecond
}
