package users

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

func (c *Client) GetUserAPI(ctx context.Context, id int) (*resty.Response, error) {
	return c.http.Get(ctx, http.GetHTTPRequest{URL: fmt.Sprintf("/users/%d", id)})
}

func (c *Client) GetUsersAPI(ctx context.Context) (*resty.Response, error) {
	return c.http.Get(ctx, http.GetHTTPRequest{URL: "/users"})
}

func (c *Client) CreateUserAPI(ctx context.Context, req CreateUserRequest) (*resty.Response, error) {
	return c.http.Post(ctx, http.PostHTTPRequest{URL: "/users/add", Body: req})
}

func (c *Client) GetUser(ctx context.Context, id int) (*User, *resty.Response, error) {
	resp, err := c.GetUserAPI(ctx, id)
	if err != nil {
		return nil, resp, err
	}

	if resp.IsError() {
		return nil, resp, fmt.Errorf("unexpected status: %d", resp.StatusCode())
	}

	var result User
	err = json.Unmarshal(resp.Body(), &result)
	return &result, resp, err
}

func (c *Client) GetUsers(ctx context.Context) (*GetUsersResponse, *resty.Response, error) {
	resp, err := c.GetUsersAPI(ctx)
	if err != nil {
		return nil, resp, err
	}

	if resp.IsError() {
		return nil, resp, fmt.Errorf("unexpected status: %d", resp.StatusCode())
	}

	var result GetUsersResponse
	err = json.Unmarshal(resp.Body(), &result)
	return &result, resp, err
}

func (c *Client) CreateUser(ctx context.Context, req CreateUserRequest) (*User, *resty.Response, error) {
	resp, err := c.CreateUserAPI(ctx, req)
	if err != nil {
		return nil, resp, err
	}

	if resp.IsError() {
		return nil, resp, fmt.Errorf("unexpected status: %d", resp.StatusCode())
	}

	var result User
	err = json.Unmarshal(resp.Body(), &result)
	return &result, resp, err
}
