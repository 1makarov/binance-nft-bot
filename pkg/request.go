package utils

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func PostRequest(url string, b []byte) (int, *[]byte, error) {
	r := fasthttp.AcquireRequest()
	r.Header.SetMethod("POST")
	r.Header.SetRequestURI(url)
	r.Header.SetContentType("application/json")
	r.SetBody(b)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(r, resp); err != nil {
		return 0, nil, err
	}
	ba := resp.Body()
	return resp.StatusCode(), &ba, nil
}

func PostRequestProxy(url string, b []byte, proxy string) (int, *[]byte, error) {
	r := fasthttp.AcquireRequest()
	r.Header.SetMethod("POST")
	r.Header.SetRequestURI(url)
	r.Header.SetContentType("application/json")
	r.SetBody(b)
	resp := fasthttp.AcquireResponse()
	client := fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}
	if err := client.Do(r, resp); err != nil {
		return 0, nil, err
	}
	ba := resp.Body()
	return resp.StatusCode(), &ba, nil
}