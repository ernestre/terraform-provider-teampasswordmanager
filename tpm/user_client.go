package tpm

import (
	"fmt"
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

func (c UserClient) Create(request CreateUserRequest) (CreateUserResponse, error) {
	response := CreateUserResponse{}
	endpoint := c.getCreateUserEndpoint()

	err := c.client.CreateResource(endpoint, request, &response)
	if err != nil {
		return response, fmt.Errorf("failed to create user: %w", err)
	}
	return response, nil
}

func (c UserClient) Delete(ID int) error {
	endpoint := c.getDeleteUserEndpoint(ID)

	err := c.client.DeleteResource(endpoint)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
