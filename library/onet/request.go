package onet

import (
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type Request struct {
	transport *http.Transport
	client    *http.Client
	timeout   time.Duration
	headers   map[string]string
	writer    *multipart.Writer
}

type MultipartParam struct {
	FieldName string
	FileName  string
	FileBody  io.Reader
}
