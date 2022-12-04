package tpm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type PasswordClient struct {
	client Client
}

func NewPasswordClient(c Config) PasswordClient {
	return PasswordClient{
		client: NewClient(c),
	}
}

func (c PasswordClient) Create(createPasswordRequest CreatePasswordRequest) (CreatePasswordResponse, error) {
	result := CreatePasswordResponse{}
	if createPasswordRequest.Name == "" {
		return result, ErrPasswordNameIsRequired
	}

	if createPasswordRequest.ProjectID == 0 {
		return result, ErrProjectIDIsRequired
	}

	reqBody, err := json.Marshal(createPasswordRequest)
	if err != nil {
		return result, fmt.Errorf("failed to marshal password body: %w", err)
	}

	endpoint := c.getPasswordsEndpoint()

	req, err := http.NewRequest(http.MethodPost, c.generateURL(endpoint), bytes.NewReader(reqBody))
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

	return result, fmt.Errorf(
		"failed to create '%s' password in project '%d': %w",
		createPasswordRequest.Password,
		createPasswordRequest.ProjectID,
		apiError,
	)
}

func (c PasswordClient) Get(passwordID int) (PasswordData, error) {
	result := PasswordData{}
	endpoint := c.getPasswordByIDEndpoint(passwordID)

	req, err := http.NewRequest(http.MethodGet, c.generateURL(endpoint), nil)
	if err != nil {
		return result, err
	}

	resp, err := c.client.sendRequest(req)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&result)

		return result, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return result, ErrPasswordNotFound
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to parse api error: %w", err)
	}

	return result, fmt.Errorf(
		"failed to fetch password by id %d': %w",
		passwordID,
		apiError,
	)
}

func (c PasswordClient) Delete(passwordID int) error {
	endpoint := c.getPasswordByIDEndpoint(passwordID)

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

	return fmt.Errorf(
		"failed to delete password by id %d': %w",
		passwordID,
		apiError,
	)
}

func (c PasswordClient) UpdatePasswordSecurity(
	passwordID int,
	updatePasswordSecurity UpdatePasswordSecurityRequest,
) error {
	reqBody, err := json.Marshal(updatePasswordSecurity)
	if err != nil {
		return fmt.Errorf("failed to marshal update password security request body: %w", err)
	}

	endpoint := c.getUpdatePasswordSecurityEndpoint(passwordID)

	req, err := http.NewRequest(http.MethodPut, c.generateURL(endpoint), bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create 'update password security' request: %w", err)
	}

	resp, err := c.client.sendRequest(req)
	if err != nil {
		return fmt.Errorf("update password security request failed: %w", err)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse api error: %w", err)
	}

	return fmt.Errorf(
		"failed to update password by id %d: %w",
		passwordID,
		apiError,
	)
}

func (c PasswordClient) Update(passwordID int, createPasswordRequest UpdatePasswordRequest) error {
	reqBody, err := json.Marshal(createPasswordRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal password update body: %w", err)
	}

	endpoint := c.getUpdatePasswordEndpoint(passwordID)

	req, err := http.NewRequest(http.MethodPut, c.generateURL(endpoint), bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create update request: %w", err)
	}

	resp, err := c.client.sendRequest(req)
	if err != nil {
		return fmt.Errorf("password update request failed: %w", err)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse api error: %w", err)
	}

	return fmt.Errorf(
		"failed to update password by id %d: %w",
		passwordID,
		apiError,
	)
}

func (c PasswordClient) Find(field string) ([]Password, error) {
	var passwords []Password

	endpoint := c.getPasswordSearchEndpoint(field)

	req, err := http.NewRequest(http.MethodGet, c.generateURL(endpoint), nil)
	if err != nil {
		return passwords, err
	}

	resp, err := c.client.sendRequest(req)
	if err != nil {
		return passwords, err
	}

	decoder := json.NewDecoder(resp.Body)

	if resp.StatusCode == http.StatusOK {
		err = decoder.Decode(&passwords)

		return passwords, err
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return passwords, fmt.Errorf("failed to parse api error: %w", err)
	}

	return passwords, fmt.Errorf("failed to find password: %w", apiError)
}

// TODO: Make this function general, or move this logic to the base client struct.
func (c PasswordClient) generateURL(p string) string {
	return fmt.Sprintf("%s/index.php/%s", c.client.config.Host, p)
}

func (c PasswordClient) getProjectPasswordsEndpoint(projectID int) string {
	return fmt.Sprintf("api/%s/projects/%d/passwords.json", c.client.config.ApiVersion, projectID)
}

func (c PasswordClient) getPasswordsEndpoint() string {
	return fmt.Sprintf("api/%s/passwords.json", c.client.config.ApiVersion)
}

func (c PasswordClient) getPasswordByIDEndpoint(passwordID int) string {
	return fmt.Sprintf("api/%s/passwords/%d.json", c.client.config.ApiVersion, passwordID)
}

func (c PasswordClient) getUpdatePasswordEndpoint(passwordID int) string {
	return fmt.Sprintf("api/%s/passwords/%d.json", c.client.config.ApiVersion, passwordID)
}

func (c PasswordClient) getUpdatePasswordSecurityEndpoint(passwordID int) string {
	return fmt.Sprintf("api/%s/passwords/%d/security.json", c.client.config.ApiVersion, passwordID)
}

func (c PasswordClient) getPasswordSearchEndpoint(searchString string) string {
	return fmt.Sprintf(
		"api/%s/passwords/search/%s.json",
		c.client.config.ApiVersion,
		url.QueryEscape(searchString),
	)
}
