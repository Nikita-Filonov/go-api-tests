package http

import (
	"context"

	"github.com/Nikita-Filonov/axiom"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	axiom  axiom.Config
	client *resty.Client
}

type GetHTTPRequest struct {
	URL     string
	Headers map[string]string
}

type PostHTTPRequest struct {
	URL     string
	Body    any
	Headers map[string]string
}

func NewClient(cfg ClientConfig, log *ClientLogger, cfgTest *axiom.Config) *Client {
	client := resty.New().
		SetLogger(log).
		SetBaseURL(cfg.URL).
		SetTimeout(cfg.Timeout).
		OnError(onErrorHook(cfgTest)).
		OnBeforeRequest(onBeforeRequestHook(cfgTest)).
		OnAfterResponse(onAfterResponseHook(cfgTest))

	return &Client{client: client}
}

func (c *Client) Get(ctx context.Context, req GetHTTPRequest) (*resty.Response, error) {
	return c.client.R().
		SetContext(ctx).
		SetHeaders(req.Headers).
		Get(req.URL)
}

func (c *Client) Post(ctx context.Context, req PostHTTPRequest) (*resty.Response, error) {
	return c.client.R().
		SetContext(ctx).
		SetHeaders(req.Headers).
		SetBody(req.Body).
		Post(req.URL)
}
