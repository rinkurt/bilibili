package bilibili

import (
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
)

type ContentType string

const (
	ContentTypeUrl  ContentType = "application/x-www-form-urlencoded"
	ContentTypeJson ContentType = "application/json"
	ContentTypeForm ContentType = "multipart/form-data"
)

func (c *Client) SendRequest(method, url string, contentType ContentType, urlParam map[string]string, bodyParam map[string]any) (out *simplejson.Json, err error) {
	r := c.resty.R()
	r.SetHeader("Content-Type", string(contentType))

	for k, v := range urlParam {
		r.SetQueryParam(k, v)
	}
	if len(bodyParam) > 0 {
		r.SetBody(bodyParam)
	}

	resp, err := r.Execute(method, url)
	if err != nil {
		return out, errors.WithStack(err)
	}
	if resp.StatusCode() != 200 {
		return out, errors.Errorf("status code: %d", resp.StatusCode())
	}
	c.SetCookies(resp.Cookies())
	body, err := simplejson.NewJson(resp.Body())
	if err != nil {
		return out, errors.WithStack(err)
	}
	code := body.Get("code").MustInt()
	message := body.Get("message").MustString()
	if code != 0 {
		return out, errors.WithStack(Error{Code: code, Message: message})
	}
	return body.Get("data"), nil
}

func (c *Client) SendUrlRequest(method, url string, urlParam map[string]string) (out *simplejson.Json, err error) {
	return c.SendRequest(method, url, ContentTypeUrl, urlParam, nil)
}

func (c *Client) SendJsonRequest(method, url string, bodyParam map[string]any) (out *simplejson.Json, err error) {
	return c.SendRequest(method, url, ContentTypeJson, nil, bodyParam)
}

func (c *Client) SendFormRequest(method, url string, bodyParam map[string]any) (out *simplejson.Json, err error) {
	return c.SendRequest(method, url, ContentTypeForm, nil, bodyParam)
}
