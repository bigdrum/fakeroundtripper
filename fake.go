package common

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/golang/glog"
)

// FakeRoundTripper for unitests.
type FakeRoundTripper struct {
	urlToFileMap    map[string]string
	urlToContentMap map[string]string
	urlToHandlerMap map[string]http.Handler
}

// RoundTrip implementation.
func (f *FakeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	glog.Infof("%s %s", req.Method, req.URL.String())
	url := req.URL.String()
	key := req.Method + " " + url

	filename, ok := f.urlToFileMap[key]
	if ok {
		f, err := os.Open("../../testdata/" + filename)

		if err != nil {
			return nil, err
		}
		resp := &http.Response{
			Body: f,
		}
		return resp, nil
	}

	content, ok := f.urlToContentMap[key]
	if ok {
		resp := &http.Response{
			Body: ioutil.NopCloser(bytes.NewReader([]byte(content))),
		}
		return resp, nil
	}

	for prefix, handler := range f.urlToHandlerMap {
		if !strings.HasPrefix(url, prefix) {
			continue
		}
		return createHandleFunc(handler)(req), nil
	}

	glog.Errorf("Content not defined for %v", url)
	return nil, errors.New("Content not defined for " + url)
}

// BindFile binds the content of a file to a URL.
func (f *FakeRoundTripper) BindFile(url string, testDataFileName string) {
	f.urlToFileMap["GET "+url] = testDataFileName
}

// BindContent binds the content of a string to a URL.
func (f *FakeRoundTripper) BindContent(url string, content string) {
	f.urlToContentMap["GET "+url] = content
}

// BindHandler binds a handler to a URL.
func (f *FakeRoundTripper) BindHandler(prefix string, handler http.Handler) {
	f.urlToHandlerMap[prefix] = handler
}

// NewFakeRoundTripper creates a new FakeRoundTripper.
func NewFakeRoundTripper() *FakeRoundTripper {
	return &FakeRoundTripper{
		make(map[string]string),
		make(map[string]string),
		make(map[string]http.Handler),
	}
}

func createHandleFunc(h http.Handler) func(request *http.Request) *http.Response {
	return func(request *http.Request) *http.Response {
		recorder := httptest.NewRecorder()
		recorder.Body = new(bytes.Buffer)
		h.ServeHTTP(recorder, request)
		resp := &http.Response{
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			StatusCode:    recorder.Code,
			Status:        http.StatusText(recorder.Code),
			Header:        recorder.HeaderMap,
			ContentLength: int64(recorder.Body.Len()),
			Body:          ioutil.NopCloser(recorder.Body),
			Request:       request,
		}
		return resp
	}
}
