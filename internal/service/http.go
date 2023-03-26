package service

import (
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xconstant/headers"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/static"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

type HttpService struct{}

func NewHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) DoRequest(client *http.Client, req *http.Request) ([]byte, *http.Response, error) {
	if req.Header.Get(headers.UserAgent) == "" {
		req.Header.Add(headers.UserAgent, static.USER_AGENT)
	}
	if req.Header.Get(headers.Referer) == "" {
		req.Header.Add(headers.Referer, static.REFERER)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("network error: %v", err)
	}
	body := resp.Body
	defer body.Close()

	bs, err := io.ReadAll(body)
	if err != nil {
		return nil, nil, fmt.Errorf("response error: %v", err)
	}
	return bs, resp, err
}

func (h *HttpService) HttpGet(url string, fn func(r *http.Request)) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if fn != nil {
		fn(req)
	}

	return h.DoRequest(&http.Client{}, req)
}

func (h *HttpService) HttpGetDocument(url string, fn func(*http.Request)) ([]byte, *goquery.Document, error) {
	bs, _, err := h.HttpGet(url, fn)
	if err != nil {
		return nil, nil, err
	}
	if bytes.Contains(bs, []byte(static.NOT_FOUND_TOKEN)) {
		return nil, nil, nil
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
	if err != nil {
		return nil, nil, fmt.Errorf("document error: %v", err)
	}
	return bs, doc, nil
}

func (h *HttpService) HttpHeadNoRedirect(url string, fn func(r *http.Request)) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	if fn != nil {
		fn(req)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	_, resp, err := h.DoRequest(client, req)
	return resp, err
}

func (h *HttpService) HttpPost(url string, body io.Reader, fn func(r *http.Request)) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}
	if fn != nil {
		fn(req)
	}

	return h.DoRequest(&http.Client{}, req)
}
