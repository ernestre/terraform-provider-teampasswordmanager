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

	c.Host = strings.ReplaceAll(c.Host, " ", "")
	c.Host = strings.TrimRight(c.Host, "/")

	return Client{
		httpClient: http.Client{
			Timeout: time.Second * 15,
		},
		config: c,
	}
}

func (c Client) generateURL(p string) string {
	return fmt.Sprintf("%s/index.php/%s", c.config.Host, p)
}

func (c Client) sendRequest(r *http.Request) (*http.Response, error) {
	err := addRequiredHeaders(c.config, r)
	if err != nil {
		return nil, fmt.Errorf("failed to add required headers to the request: %w", err)
	}

	return c.httpClient.Do(r)
}

func (c Client) UpdateResource(endpoint string, requestBody any) error {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal project body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, c.generateURL(endpoint), bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to marshal create user request body: %w", err)
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse api error: %w", err)
	}

	return fmt.Errorf("failed to update resource: %w", apiError)
}

func (c Client) CreateResource(
	endpoint string,
	requestBody any,
	response any,
) error {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal create user request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.generateURL(endpoint), bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	decoder := json.NewDecoder(resp.Body)

	if resp.StatusCode == http.StatusCreated {
		err = decoder.Decode(&response)
		if err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return nil
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse api error: %w", err)
	}

	return fmt.Errorf("failed to create group: %w", apiError)
}

func (c Client) GetResource(endpoint string, response any) error {
	req, err := http.NewRequest(http.MethodGet, c.generateURL(endpoint), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return ErrProjectNotFound
	}

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&response)

		if err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return nil
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse api error: %w", err)
	}

	return fmt.Errorf("failed to retrieve resource: %w", apiError)
}

func (c Client) DeleteResource(endpoint string) error {
	req, err := http.NewRequest(http.MethodDelete, c.generateURL(endpoint), nil)
	if err != nil {
		return err
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse api error: %w", err)
	}

	return fmt.Errorf("failed to delete resource: %w", apiError)
}

func trimEndpoint(requestURI string) string {
	parts := strings.Split(requestURI, "index.php")

	if len(parts) > 1 {
		lastPart := parts[len(parts)-1]
		return strings.TrimLeft(lastPart, "/")
	}

	return strings.TrimLeft(parts[0], "/")
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

	endpoint := trimEndpoint(request.URL.RequestURI())
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
