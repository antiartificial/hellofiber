package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestBaseRoute(t *testing.T) {
	tests := []struct {
		description  string // description of test case
		route        string // route path to test
		expectedCode int    // expected HTTP status
	}{
		{
			description:  "get HTTP status 200",
			route:        "/",
			expectedCode: 200,
		},
		{
			description:  "get HTTP 404 when route does not exist",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	app := New()

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, -1)

		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestAvailabilityPostJSONEndpoint(t *testing.T) {
	p := new(Person)
	p.Name = "Test"
	p.Pass = "12345"
	personData, err := json.Marshal(p)
	if err != nil {
		// todo
	}

	tests := []struct {
		description  string // description of test case
		route        string // route path to test
		queryData    []byte // what's to be passed
		expectedCode int    // expected HTTP status
	}{
		{
			description:  "get HTTP status 200",
			route:        "/availability",
			queryData:    personData,
			expectedCode: 200,
		},
	}

	app := New()

	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, strings.NewReader(string(personData)))
		req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
		req.Header.Add("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestAvailabilityPostFormEncodedEndpoint(t *testing.T) {
	personData := []byte("Name=Test&Pass=12345")

	tests := []struct {
		description  string // description of test case
		route        string // route path to test
		queryData    []byte // what's to be passed
		expectedCode int    // expected HTTP status
	}{
		{
			description:  "get HTTP status 200",
			route:        "/availability",
			queryData:    personData,
			expectedCode: 200,
		},
	}

	app := New()

	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, strings.NewReader(string(personData)))
		req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, _ := app.Test(req, -1)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
