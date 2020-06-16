package gorvison

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Auth struct {
	Username string
	APIToken string
}

type Revision struct {
	auth    *Auth
	baseURL string
	client  *http.Client
}

func NewRevision(auth *Auth, baseURL string) *Revision {
	return &Revision{
		auth:    auth,
		baseURL: baseURL,
		client:  http.DefaultClient,
	}
}

// SetHTTPClient with timeouts or insecure transport, etc.
func (r *Revision) SetHTTPClient(client *http.Client) {
	r.client = client
}

func (r *Revision) buildURL(path string, params url.Values) (requestURL string) {
	requestURL = r.baseURL + path + "/api/json"

	if params != nil {
		queryString := params.Encode()
		if queryString != "" {
			requestURL = requestURL + "?" + queryString
		}
	}

	return
}

func (r *Revision) sendRequest(req *http.Request) (*http.Response, error) {
	if r.auth != nil {
		req.SetBasicAuth(r.auth.Username, r.auth.APIToken)
	}

	r.client = &http.Client{Transport: &http.Transport{
		MaxIdleConns:       1,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}}

	response, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("(sendRequest)(client.Do) Error: %w", err)
	}

	return response, nil
}

func (r *Revision) parseResponse(resp *http.Response, body interface{}) (err error) {
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

func (r *Revision) get(path string, params url.Values, body interface{}) error {
	requestURL := r.buildURL(path, params)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return fmt.Errorf("(get)(NewRequest) Error: %w", err)
	}

	resp, err := r.sendRequest(req)
	if err != nil {
		return fmt.Errorf("(get)(sendRequest) Error: %w", err)
	}

	if err := r.parseResponse(resp, body); err != nil {
		return fmt.Errorf("(get)(parseResponse) Error: %w", err)
	}

	return nil
}
