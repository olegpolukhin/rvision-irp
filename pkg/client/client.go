package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olegpolukhin/rvision-irp/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type RevisionApp struct {
	Auth    *config.Auth
	BaseURL string
	Client  *http.Client
}

// SetHTTPClient with timeouts or insecure transport, etc.
func (r *RevisionApp) SetHTTPClient(client *http.Client) {
	r.Client = client
}

func (r *RevisionApp) buildURL(path string) (requestURL string) {
	u, _ := url.ParseRequestURI(r.BaseURL)
	u.Path = path
	requestURL = u.String()

	return
}

func (r *RevisionApp) sendRequest(req *http.Request) (*http.Response, error) {
	if r.Auth != nil {
		req.SetBasicAuth(r.Auth.Username, r.Auth.APIToken)
	}

	r.Client = &http.Client{Transport: &http.Transport{
		MaxIdleConns:       1,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}}

	response, err := r.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("(sendRequest)(client.Do) Error: %w", err)
	}

	return response, nil
}

func (r *RevisionApp) parseResponse(resp *http.Response, body interface{}) (err error) {
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("(parseResponse)(ReadAll) Error: %w", err)
	}

	if err := json.Unmarshal(data, body); err != nil {
		return fmt.Errorf("(parseResponse)(Unmarshal) Error: %w", err)
	}

	return nil
}

func (r *RevisionApp) Get(path string, params url.Values, body interface{}) error {
	requestURL := r.buildURL(path)

	req, err := http.NewRequest(http.MethodGet, requestURL, strings.NewReader(params.Encode()))
	if err != nil {
		return fmt.Errorf("(get)(NewRequest) Error: %w", err)
	}

	resp, err := r.sendRequest(req)
	if err != nil {
		return fmt.Errorf("(get)(sendRequest) Error: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound ||
		resp.StatusCode == http.StatusBadRequest ||
		resp.StatusCode == http.StatusInternalServerError {
		return fmt.Errorf("(get)(parseResponse) HTTP Error: status code %d", resp.StatusCode)
	}

	if err := r.parseResponse(resp, body); err != nil {
		return fmt.Errorf("(get)(parseResponse) Error: %w", err)
	}

	return nil
}

func (r *RevisionApp) Post(path string, params map[string]string, body interface{}) error {
	requestURL := r.buildURL(path)

	var buf *bytes.Buffer

	if len(params) != 0 {
		byteData, err := json.Marshal(params)
		if err != nil {
			return fmt.Errorf("(post)(Marshal) Error: %w", err)
		}

		buf = bytes.NewBufferString(string(byteData))
	} else {
		buf = &bytes.Buffer{}
	}

	req, err := http.NewRequest(http.MethodPost, requestURL, buf)
	if err != nil {
		return fmt.Errorf("(post)(NewRequest) Error: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("Content-Length", strconv.Itoa(len(params)))

	resp, err := r.sendRequest(req)
	if err != nil {
		return fmt.Errorf("(post)(sendRequest) Error: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound ||
		resp.StatusCode == http.StatusBadRequest ||
		resp.StatusCode == http.StatusInternalServerError {
		return fmt.Errorf("(post)(parseResponse) HTTP Error: status code %d", resp.StatusCode)
	}

	if err := r.parseResponse(resp, body); err != nil {
		return fmt.Errorf("(post)(parseResponse) Error: %w", err)
	}

	return nil
}
