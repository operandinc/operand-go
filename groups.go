package operand

import (
	"context"
	"net/url"
	"strconv"
)

// ListGroups returns a list of all configured groups.
func (c *Client) ListGroups(ctx context.Context, collectionId *string) ([]Group, error) {
	groups := []Group{}
	for {
		params := url.Values{
			"limit": []string{"100"},
			"offset": []string{
				strconv.Itoa(len(groups)),
			},
		}
		if collectionId != nil {
			params.Add("collection", *collectionId)
		}
		url := c.endpoint + "/v2/group?" + params.Encode()
		resp, err := doRequest[any, ListResponse[Group]](ctx, c, "GET", url, nil)
		if err != nil {
			return nil, err
		}
		groups = append(groups, resp.Items...)
		if !resp.More {
			break
		}
	}
	return groups, nil
}

// GetGroup returns a group by ID.
func (c *Client) GetGroup(ctx context.Context, id string) (*Group, error) {
	url := c.endpoint + "/v2/group/" + url.PathEscape(id)
	resp, err := doRequest[any, Group](ctx, c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateGroupRequest is used to create a group.
type CreateGroupRequest struct {
	CollectionID string `json:"collectionId"`
	Name         string `json:"name"`
}

// CreateGroup creates a new group.
func (c *Client) CreateGroup(ctx context.Context, req *CreateGroupRequest) (*Group, error) {
	resp, err := doRequest[CreateGroupRequest, Group](ctx, c, "POST", c.endpoint+"/v2/group", req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateGroupRequest is used to update a group.
type UpdateGroupRequest struct {
	Name string `json:"name"`
}

// UpdateGroup updates a group.
func (c *Client) UpdateGroup(ctx context.Context, id string, req *UpdateGroupRequest) (*Group, error) {
	url := c.endpoint + "/v2/group/" + url.PathEscape(id)
	resp, err := doRequest[UpdateGroupRequest, Group](ctx, c, "PUT", url, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteGroup deletes a group.
func (c *Client) DeleteGroup(ctx context.Context, id string) error {
	url := c.endpoint + "/v2/group/" + url.PathEscape(id)
	_, err := doRequest[any, any](ctx, c, "DELETE", url, nil)
	if err != nil {
		return err
	}
	return nil
}
