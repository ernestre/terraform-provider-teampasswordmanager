package tpm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const DefaultApiVersion = "v5"

type (
	Config struct {
		Host       string
		PublicKey  string
		PrivateKey string
		ApiVersion string
	}

	Client struct {
		httpClient http.Client
		config     Config
	}
)

func NewClient(c Config) Client {
	if c.ApiVersion == "" {
		c.ApiVersion = DefaultApiVersion
	}
	return Client{
		httpClient: http.Client{
			Timeout: time.Second * 15,
		},
		config: c,
	}
}

func (c Client) sendRequest(r *http.Request) (*http.Response, error) {
	err := addRequiredHeaders(c.config, r)
	if err != nil {
		return nil, fmt.Errorf("failed to add required headers to the request: %w", err)
	}

	return c.httpClient.Do(r)
}

func addRequiredHeaders(
	config Config,
	request *http.Request,
) error {
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	var body []byte

	if request.Body != nil {
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return err
		}

		body = b
	}

	endpoint := strings.TrimPrefix(request.URL.RequestURI(), "/index.php/")
	time := time.Now().Unix()
	hash := generateAuthHash(endpoint, time, body, config.PrivateKey)
	headers := generateAuthHeaders(config.PublicKey, hash, time)

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return nil
}

func errorResponseToApiError(responseBody io.ReadCloser) (ApiError, error) {
	apiError := ApiError{}
	errorResponse := ErrorResponse{}
	err := json.NewDecoder(responseBody).Decode(&errorResponse)
	if err != nil {
		return apiError, err
	}

	apiError.Message = errorResponse.Message
	apiError.Type = errorResponse.Type

	return apiError, err
}
