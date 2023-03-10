package service

import (
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/static"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

type HttpService struct {
}

func NewHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) DoRequest(req *http.Request) ([]byte, *http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("network error: %v", err)
	}
	body := resp.Body
	defer body.Close()

	bs, err := ioutil.ReadAll(body)
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
	req.Header.Add("User-Agent", static.USER_AGENT)
	if fn != nil {
		fn(req)
	}

	bs, resp, err := h.DoRequest(req)
	return bs, resp, err
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
	req.Header.Add("User-Agent", static.USER_AGENT)
	if fn != nil {
		fn(req)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()
	return resp, err
}