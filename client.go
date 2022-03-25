package operand

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is a client for the Operand API.
type Client struct {
	apiKey   string
	endpoint string
	client   *http.Client
}

// An Option can be used to configure a Client.
type Option interface {
	Apply(*Client)
}

// NewClient creates a new Client with a set of options.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:   apiKey,
		endpoint: "https://api.operand.ai",
		client:   http.DefaultClient,
	}

	for _, opt := range opts {
		opt.Apply(c)
	}

	return c
}

type optionFunc func(*Client)

func (of optionFunc) Apply(c *Client) {
	of(c)
}

// WithEndpoint sets the endpoint of the Client.
func WithEndpoint(endpoint string) Option {
	return optionFunc(func(c *Client) {
		c.endpoint = endpoint
	})
}

// WithHTTPClient sets the HTTP client of the Client.
func WithHTTPClient(client *http.Client) Option {
	return optionFunc(func(c *Client) {
		c.client = client
	})
}

// Common API Errors.
var (
	ErrNotFound = fmt.Errorf("not found")
)

func doRequest[Req, Resp any](
	ctx context.Context,
	client *Client,
	method, url string,
	body *Req,
) (*Resp, error) {
	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", client.apiKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var ret Resp
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return nil, err
		}
		return &ret, nil
	case http.StatusNotFound:
		return nil, ErrNotFound
	default:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(
			"unexpected status code %d (%s): %s",
			resp.StatusCode,
			resp.Status,
			string(body),
		)
	}
}
