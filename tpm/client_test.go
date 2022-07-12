package tpm

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var testConfig = Config{
	Host:       "host",
	PublicKey:  "publicKey",
	PrivateKey: "privateKey",
	ApiVersion: "v5",
}

func TestNewClient(t *testing.T) {
	c := Config{
		Host:       "host",
		PublicKey:  "publicKey",
		PrivateKey: "privateKey",
		ApiVersion: "v4",
	}

	expectedVersion := "v4"

	client := NewClient(c)

	if client.config.ApiVersion != expectedVersion {
		t.Errorf(
			"client has invalid api version. Expected %s, got %s",
			expectedVersion,
			client.config.ApiVersion,
		)
	}

	if client.config.Host != c.Host {
		t.Errorf(
			"client has invalid host. Expected %s, got %s",
			c.Host,
			client.config.Host,
		)
	}

	if client.config.PublicKey != c.PublicKey {
		t.Errorf(
			"client has invalid public key. Expected %s, got %s",
			c.PublicKey,
			client.config.PublicKey,
		)
	}

	if client.config.PrivateKey != c.PrivateKey {
		t.Errorf(
			"client has invalid private key. Expected %s, got %s",
			c.PrivateKey,
			client.config.PrivateKey,
		)
	}
}

func TestAddRequiredHeaders(t *testing.T) {
	endpoint := "/foo"
	requestBody := []byte(`{"foo":123}`)
	data := bytes.NewReader(requestBody)

	request, err := http.NewRequest(http.MethodPost, endpoint, data)
	if err != nil {
		t.Error("failed to create request")
	}

	err = addRequiredHeaders(testConfig, request)
	if err != nil {
		t.Fatal("failed to add auth headers to request")
	}

	if request.Body == nil {
		t.Fatal("request body should not be empty after adding headers")
	}

	rb, err := io.ReadAll(request.Body)
	if err != nil {
		t.Fatal("failed to read request body")
	}

	if string(rb) != string(requestBody) {
		t.Fatal("request body was altered after header generation")
	}

	if ct, ok := request.Header["Content-Type"]; ok {
		if ct[0] != "application/json; charset=utf-8" {
			t.Errorf("invalid content type header. Expected %s, got %s", "application/json; charset=utf-8", ct[0])
		}
	} else {
		t.Fatal("request is missing content-type header")
	}

	assureHeaderExists := func(h http.Header, k string) {
		if h, ok := request.Header[k]; ok {
			if h[0] == "" {
				t.Fatalf("auth header %s is empty", k)
			}
		} else {
			t.Fatalf("auth header %s is missing", k)
		}
	}

	ts, err := strconv.ParseInt(request.Header.Get(AuthHeaderRequestTimestamp), 10, 64)
	if err != nil {
		t.Fatal("failed to convert timestamp to int")
	}

	authHash := generateAuthHash(
		endpoint,
		ts,
		requestBody,
		testConfig.PrivateKey,
	)

	if authHash != request.Header.Get(AuthHeaderRequestHash) {
		t.Fatal("auth hash does not match the auth hash header")
	}

	assureHeaderExists(request.Header, AuthHeaderPublicKey)
	assureHeaderExists(request.Header, AuthHeaderRequestHash)
	assureHeaderExists(request.Header, AuthHeaderRequestTimestamp)
}

func TestErrorResponseToApiError(t *testing.T) {
	errorType := "some-type"
	errorMessage := "something went wrong"
	json := fmt.Sprintf(`{"error":true, "type":"%s", "message":"%s"}`, errorType, errorMessage)

	r := io.NopCloser(strings.NewReader(json))

	apiErr, err := errorResponseToApiError(r)
	if err != nil {
		t.Error("failed to decode api error response")
	}

	if apiErr.Type != errorType {
		t.Errorf("invalid error type. Expected %s, got %s", errorType, apiErr.Type)
	}

	if apiErr.Message != errorMessage {
		t.Errorf("invalid error message. Expected %s, got %s", errorMessage, apiErr.Message)
	}
}

func TestSendRequestWithEmptyBody(t *testing.T) {
	var receivedHash string
	var receivedTimestamp int64
	var receivedPublicKey string
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHash = r.Header.Get(AuthHeaderRequestHash)
		ts, err := strconv.ParseInt(r.Header.Get(AuthHeaderRequestTimestamp), 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		receivedTimestamp = ts
		receivedPublicKey = r.Header.Get(AuthHeaderPublicKey)
	}))
	defer svr.Close()

	endpoint := "/index.php/foo/bar"

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s%s", svr.URL, endpoint),
		nil,
	)

	c := NewClient(testConfig)

	_, err := c.sendRequest(req)
	if err != nil {
		t.Errorf("failed to send request")
	}

	ts, err := strconv.ParseInt(req.Header.Get(AuthHeaderRequestTimestamp), 10, 64)
	if err != nil {
		t.Fatal("failed to convert timestamp to int")
	}

	authHash := generateAuthHash(
		"foo/bar",
		ts,
		nil,
		testConfig.PrivateKey,
	)

	if authHash != receivedHash {
		t.Error("generated hash does not match the hash received on the server side")
	}

	if ts != receivedTimestamp {
		t.Error("hash timestamp does not match the timestamp header value received on the server side")
	}

	if testConfig.PublicKey != receivedPublicKey {
		t.Error("generated hash does not match the hash received on the server side")
	}
}
