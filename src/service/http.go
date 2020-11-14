package service

import (
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

type HttpService struct {
}

func NewHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) HttpGet(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", static.USER_AGENT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	body := resp.Body
	defer body.Close()

	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("response error: %v", err)
	}
	return bs, err
}

func (h *HttpService) HttpGetDocument(url string) (*goquery.Document, error) {
	bs, err := h.HttpGet(url)
	if err != nil {
		return nil, err
	}
	if bytes.Contains(bs, []byte(static.NOT_FOUND_TOKEN)) || bytes.Contains(bs, []byte(static.NOT_FOUND2_TOKEN)) || bytes.Contains(bs, []byte(static.NOT_FOUND3_TOKEN)) {
		return nil, nil
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
	if err != nil {
		return nil, fmt.Errorf("document error: %v", err)
	}
	return doc, nil
}
