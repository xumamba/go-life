package onet

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
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
	var resp *http.Response
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return
	}
	// if isJson {
	// 	req.Header.Set("Content-Type", JSON)
	// } else {
	// 	req.Header.Set("Content-Type", TEXT)
	// }
	// increase the max connection per host to prevent error "no free connection available" error while sending more requests.
	hc.client.Transport.(*http.Transport).MaxIdleConnsPerHost = 512 * 20
	resp, err = hc.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	respBody, err = ioutil.ReadAll(resp.Body)
	return
}

func (hc *HTTPClient) PostForm(url string, data url.Values) (statusCode int, respBody []byte, err error) {
	resp, err := hc.client.PostForm(url, data)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	respBody, err = ioutil.ReadAll(resp.Body)
	return

}

func (hc *HTTPClient) PostJson(url string, data []byte) (statusCode int, respBody []byte, err error) {
	var resp *http.Response
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	// increase the max connection per host to prevent error "no free connection available" error while sending more requests.
	hc.client.Transport.(*http.Transport).MaxIdleConnsPerHost = 512 * 20
	resp, err = hc.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	respBody, err = ioutil.ReadAll(resp.Body)
	return
}

func (hc *HTTPClient) PostJsonObj(url string, req, resp interface{}) (err error) {
	var response *http.Response
	if req != nil {
		b, err := json.Marshal(req)
		if err != nil {
			return err
		}
		response, err = hc.client.Post(url, "application/json;charset=utf-8", bytes.NewReader(b))
	} else {
		response, err = hc.client.Post(url, "application/json;charset=utf-8", nil)
	}

	if err != nil {
		return
	}

	if response != nil {
		defer response.Body.Close()
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("http response status code !=200,code:%d", response.StatusCode)
	}
	return json.NewDecoder(response.Body).Decode(resp)

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
