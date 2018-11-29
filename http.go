package tata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func MakeGetRequest(client *Client, path string, query url.Values) (string, error) {
	url, err := composeURL(client.getBaseURL(), path, composeQueryString(query))
	if err != nil {
		return "", err
	}

	fmt.Sprintf("This is url: %s", url)

	httpReq, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", err
	}
	fillHeaders(&httpReq.Header, client.HeaderAuthentication)

	httpRes, err := client.getClient().Do(httpReq)
	if err != nil {
		return "", err
	}

	httpResBody, err := handleResponse(httpRes)
	if err != nil {
		return "", err
	}

	//tata is using Content-type: text/html
	if len(httpResBody) <= 0 {
		return "", errEmptyResponse
	}

	return string(httpResBody), nil
}

func composeURL(baseURL string, path string, query string) (*url.URL, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	full, err := base.Parse(path + query)
	return full, err
}

func composeQueryString(data url.Values) string {
	if data == nil {
		return ""
	}

	var query string
	for key, val := range data {
		joined := strings.Join(val, ",")
		escaped := url.QueryEscape(joined)
		query += fmt.Sprintf("&%s=%s", key, escaped)
	}
	if query != "" {
		query = "?" + query[1:]
	}
	return query
}

func fillHeaders(header *http.Header, creds Credentials) {
	header.Add("Authorization", creds.GetBasicAuthString())
	header.Add("Content-Type", "application/json")
	header.Add("Accept", "application/json")
	header.Add("User-Agent", "tata:APIv1_GO_CLIENT")
}

func handleResponse(res *http.Response) ([]byte, error) {
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if !isSuccess(res.StatusCode) {
		var rerr Error
		rerr.StatusCode = res.StatusCode
		return nil, rerr
	}

	return responseData, nil
}

func isSuccess(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusBadRequest
}
