package firefly

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	contentType = "application/json"
	acceptType  = "application/vnd.api+json"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    *url.URL
	common     service

	Transactions TransactionsServiceInterface
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client, baseURL string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	parsedBaseURL, _ := url.Parse(baseURL)

	c := &Client{HTTPClient: httpClient, BaseURL: parsedBaseURL}
	c.common.client = c
	c.Transactions = (*TransactionsService)(&c.common)

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	reqURL, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, reqURL.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", contentType)
	}

	req.Header.Set("Accept", acceptType)

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, result interface{}) (*Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if code := resp.StatusCode; code < 200 || 299 < code {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))

		return nil, errors.New("invalid status code")
	}

	decErr := json.NewDecoder(resp.Body).Decode(result)
	if decErr == io.EOF {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		err = decErr
	}

	return &Response{Response: resp}, err
}

type Response struct {
	*http.Response
}

type Timestamp struct {
	time.Time
}

type Amount struct {
	*int
}
