package tata

//
// All rights reserved.
//
// Author: Junjie Zhu <zjunjie@uber.com>
//
// Package tata_go is an library for interacting with Tata-Communications sms API.

import (
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultBaseUrl = "https://sms3.tatacommunications.com:49801"

	DefaultPath = "/sendsms"

	DlrMask = "15"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 15 * time.Second
)

var (
	defaultHTTPClient = http.DefaultClient
)

// Client is used to access API with a given key.
// Uses standard lib HTTP client internally, so should be reused instead of created as needed and it is safe for concurrent use.
type Client struct {
	BaseURL              string       //Base URL
	BodyAuthentication   Credentials // Credentials for body
	HeaderAuthentication Credentials // Crendentials for header
	HTTPClient           *http.Client // The HTTP client to send requests on
}

func (c Client) getClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return defaultHTTPClient
}

func (c Client) getBaseURL() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}
	return DefaultBaseUrl
}

func (c *Client) SendSMSMessage(from string, to string, body string, callbackUrl string) (string, error) {
	queryParams := prepareQueryParams(from, to, body, callbackUrl, c.BodyAuthentication)

	res, err := MakeGetRequest(c, DefaultPath, queryParams)
	if err != nil {
		return "", err
	}

	return res, nil
}

func prepareQueryParams(from string, to string, body string, callbackUrl string, bodyAuth Credentials) url.Values {
	v := url.Values{}
	v.Set("text", body)
	v.Set("from", from)
	v.Set("to", to)
	v.Set("username", bodyAuth.Username)
	v.Set("password", bodyAuth.Password)

	if callbackUrl != "" {
		v.Set("dlr-mask", DlrMask)
		v.Set("dlr-url", callbackUrl)
	}

	return v
}
