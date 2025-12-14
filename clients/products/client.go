package products

import (
	"context"
	"encoding/json"
	"fmt"

	"go-api-tests/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	http *http.Client
}

func NewClient(cfg http.ClientConfig, log *http.ClientLogger, cfgTest *axiom.Config) *Client {
	return &Client{http: http.NewClient(cfg, log, cfgTest)}
}

func (c *Client) GetProductAPI(ctx context.Context, id int) (*resty.Response, error) {
	return c.http.Get(ctx, http.GetHTTPRequest{URL: fmt.Sprintf("/products/%d", id)})
}

func (c *Client) GetProductsAPI(ctx context.Context) (*resty.Response, error) {
	return c.http.Get(ctx, http.GetHTTPRequest{URL: "/products"})
}

func (c *Client) CreateProductAPI(ctx context.Context, req CreateProductRequest) (*resty.Response, error) {
	return c.http.Post(ctx, http.PostHTTPRequest{URL: "/products/add", Body: req})
}

func (c *Client) GetProduct(ctx context.Context, id int) (*Product, *resty.Response, error) {
	resp, err := c.GetProductAPI(ctx, id)
	if err != nil {
		return nil, resp, err
	}

	if resp.IsError() {
		return nil, resp, fmt.Errorf("unexpected status: %d", resp.StatusCode())
	}

	var result Product
	err = json.Unmarshal(resp.Body(), &result)
	return &result, resp, err
}

func (c *Client) GetProducts(ctx context.Context) (*GetProductsResponse, *resty.Response, error) {
	resp, err := c.GetProductsAPI(ctx)
	if err != nil {
		return nil, resp, err
	}

	if resp.IsError() {
		return nil, resp, fmt.Errorf("unexpected status: %d", resp.StatusCode())
	}

	var result GetProductsResponse
	err = json.Unmarshal(resp.Body(), &result)
	return &result, resp, err
}

func (c *Client) CreateProduct(ctx context.Context, req CreateProductRequest) (*Product, *resty.Response, error) {
	resp, err := c.CreateProductAPI(ctx, req)
	if err != nil {
		return nil, resp, err
	}

	if resp.IsError() {
		return nil, resp, fmt.Errorf("unexpected status: %d", resp.StatusCode())
	}

	var result Product
	err = json.Unmarshal(resp.Body(), &result)
	return &result, resp, err
}
