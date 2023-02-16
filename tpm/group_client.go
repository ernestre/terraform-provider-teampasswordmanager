package tpm

import (
	"fmt"
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

func (c GroupClient) getGroupByIdEndpoint(ID int) string {
	return fmt.Sprintf("api/%s/groups/%d.json", c.client.config.ApiVersion, ID)
}

func (c GroupClient) Create(request CreateGroupRequest) (CreateGroupResponse, error) {
	response := CreateGroupResponse{}
	endpoint := c.getCreateGroupEndpoint()

	err := c.client.CreateResource(endpoint, request, &response)
	if err != nil {
		return response, fmt.Errorf("failed to create group: %w", err)
	}
	return response, nil
}

func (c GroupClient) Update(ID int, request UpdateGroupRequest) error {
	endpoint := c.getGroupByIdEndpoint(ID)

	err := c.client.UpdateResource(endpoint, request)
	if err != nil {
		return fmt.Errorf("failed to update group: %w", err)
	}
	return nil
}

func (c GroupClient) Get(ID int) (GetGroupResponse, error) {
	var result GetGroupResponse
	endpoint := c.getGroupByIdEndpoint(ID)

	err := c.client.GetResource(endpoint, &result)
	if err != nil {
		return result, fmt.Errorf("failed to get group resource: %w", err)
	}

	return result, nil
}

func (c GroupClient) Delete(ID int) error {
	endpoint := c.getGroupByIdEndpoint(ID)

	err := c.client.DeleteResource(endpoint)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}
