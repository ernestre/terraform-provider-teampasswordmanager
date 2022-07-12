package tpm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectClient struct {
	client Client
}

func NewProjectClient(c Config) ProjectClient {
	return ProjectClient{
		client: NewClient(c),
	}
}

func (c ProjectClient) getProjectsEndpoint() string {
	return fmt.Sprintf("api/%s/projects.json", c.client.config.ApiVersion)
}

func (c ProjectClient) getSpecificProjectEndpoint(ID int) string {
	return fmt.Sprintf("api/%s/projects/%d.json", c.client.config.ApiVersion, ID)
}

func (c ProjectClient) generateURL(p string) string {
	return fmt.Sprintf("%s/index.php/%s", c.client.config.Host, p)
}

func (c ProjectClient) Create(r CreateProjectRequest) (CreateProjectResponse, error) {
	result := CreateProjectResponse{}
	if r.Name == "" {
		return result, ErrProjectNameIsRequired
	}

	body, err := json.Marshal(r)
	if err != nil {
		return result, fmt.Errorf("failed to marshal project body: %w", err)
	}

	endpoint := c.getProjectsEndpoint()

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

	return result, fmt.Errorf("failed to create '%s' project: %w", r.Name, apiError)
}

func (c ProjectClient) Update(ID int, r UpdateProjectRequest) error {
	body, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("failed to marshal project body: %w", err)
	}

	endpoint := c.getSpecificProjectEndpoint(ID)

	req, err := http.NewRequest(http.MethodPut, c.generateURL(endpoint), bytes.NewReader(body))
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

	return fmt.Errorf("failed to update project (ID: %d): %w", ID, apiError)
}

func (c ProjectClient) Delete(ID int) error {
	endpoint := c.getSpecificProjectEndpoint(ID)

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

	return fmt.Errorf("failed to delete '%d' project: %w", ID, apiError)
}

func (c ProjectClient) Get(ID int) (ProjectData, error) {
	var result ProjectData
	endpoint := c.getSpecificProjectEndpoint(ID)

	req, err := http.NewRequest(http.MethodGet, c.generateURL(endpoint), nil)
	if err != nil {
		return result, err
	}

	resp, err := c.client.sendRequest(req)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return result, ErrProjectNotFound
	}

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&result)
		return result, err
	}

	apiError, err := errorResponseToApiError(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to parse api error: %w", err)
	}

	return result, fmt.Errorf("failed to retrieve project (ID: %d): %w", ID, apiError)
}
