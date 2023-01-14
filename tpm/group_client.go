package tpm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GroupClient struct {
	client Client
}

func NewGroupClient(c Config) GroupClient {
	return GroupClient{
		client: NewClient(c),
	}
}

func (c GroupClient) getCreateGroupEndpoint() string {
	return fmt.Sprintf("api/%s/groups.json", c.client.config.ApiVersion)
}

func (c GroupClient) getDeleteGroupEndpoint(ID int) string {
	return fmt.Sprintf("api/%s/groups/%d.json", c.client.config.ApiVersion, ID)
}

func (c GroupClient) generateURL(p string) string {
	return fmt.Sprintf("%s/index.php/%s", c.client.config.Host, p)
}

func (c GroupClient) Create(r CreateGroupRequest) (CreateGroupResponse, error) {
	result := CreateGroupResponse{}

	body, err := json.Marshal(r)
	if err != nil {
		return result, fmt.Errorf("failed to marshal create user request body: %w", err)
	}

	endpoint := c.getCreateGroupEndpoint()

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

	return result, fmt.Errorf("failed to create group: %w", apiError)
}

func (c GroupClient) Delete(ID int) error {
	endpoint := c.getDeleteGroupEndpoint(ID)

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

	return fmt.Errorf("failed to delete group: %w", apiError)
}
