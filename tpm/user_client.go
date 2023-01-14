package tpm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserClient struct {
	client Client
}

func NewUserClient(c Config) UserClient {
	return UserClient{
		client: NewClient(c),
	}
}

func (c UserClient) getCreateUserEndpoint() string {
	return fmt.Sprintf("api/%s/users.json", c.client.config.ApiVersion)
}

func (c UserClient) getDeleteUserEndpoint(ID int) string {
	return fmt.Sprintf("api/%s/users/%d.json", c.client.config.ApiVersion, ID)
}

func (c UserClient) generateURL(p string) string {
	return fmt.Sprintf("%s/index.php/%s", c.client.config.Host, p)
}

func (c UserClient) Create(r CreateUserRequest) (CreateUserResponse, error) {
	result := CreateUserResponse{}

	body, err := json.Marshal(r)
	if err != nil {
		return result, fmt.Errorf("failed to marshal create user request body: %w", err)
	}

	endpoint := c.getCreateUserEndpoint()

	req, err := http.NewRequest(http.MethodPost, c.generateURL(endpoint), bytes.NewReader(body))
	if err != nil {
		return result, err
	}

	resp, err := c.client.sendRequest(req)
	if err != nil {
		return result, err
	}

	decoder := json.NewDecoder(resp.Body)

	if resp.StatusCode == http.StatusCreated {
		err = decoder.Decode(&result)

		return result, err
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to parse api error: %w", err)
	}

	return result, fmt.Errorf("failed to create user (username: %s): %w", r.Username, apiError)
}

func (c UserClient) Delete(ID int) error {
	endpoint := c.getDeleteUserEndpoint(ID)

	req, err := http.NewRequest(http.MethodDelete, c.generateURL(endpoint), nil)
	if err != nil {
		return err
	}

	resp, err := c.client.sendRequest(req)
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

	return fmt.Errorf("failed to delete user (id %d): %w", ID, apiError)
}
