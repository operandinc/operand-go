package operand

import (
	"context"
	"net/url"
	"strconv"
)

// ListCollections returns a list of all configured collections.
func (c *Client) ListCollections(ctx context.Context) ([]Collection, error) {
	collections := []Collection{}
	for {
		params := url.Values{
			"limit": []string{"100"},
			"offset": []string{
				strconv.Itoa(len(collections)),
			},
		}
		url := c.endpoint + "/v2/collection?" + params.Encode()
		resp, err := doRequest[any, ListResponse[Collection]](ctx, c, "GET", url, nil)
		if err != nil {
			return nil, err
		}
		collections = append(collections, resp.Items...)
		if !resp.More {
			break
		}
	}
	return collections, nil
}

// GetCollection returns a collection by name.
func (c *Client) GetCollection(ctx context.Context, name string) (*Collection, error) {
	url := c.endpoint + "/v2/collection/name/" + url.PathEscape(name)
	resp, err := doRequest[any, Collection](ctx, c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetCollectionByID returns a collection by ID.
func (c *Client) GetCollectionByID(ctx context.Context, id string) (*Collection, error) {
	url := c.endpoint + "/v2/collection/" + url.PathEscape(id)
	resp, err := doRequest[any, Collection](ctx, c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateCollectionRequest is used to create a collection.
type CreateCollectionRequest struct {
	Name string `json:"name"`
}

// CreateCollection creates a new collection.
func (c *Client) CreateCollection(ctx context.Context, req *CreateCollectionRequest) (*Collection, error) {
	resp, err := doRequest[CreateCollectionRequest, Collection](ctx, c, "POST", c.endpoint+"/v2/collection", req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateCollectionRequest is used to update a collection.
type UpdateCollectionRequest struct {
	Name string `json:"name"`
}

// UpdateCollection updates an existing collection.
func (c *Client) UpdateCollection(ctx context.Context, id string, req *UpdateCollectionRequest) (*Collection, error) {
	url := c.endpoint + "/v2/collection/" + url.PathEscape(id)
	resp, err := doRequest[UpdateCollectionRequest, Collection](ctx, c, "PUT", url, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteCollection deletes an existing collection.
func (c *Client) DeleteCollection(ctx context.Context, id string) (bool, error) {
	url := c.endpoint + "/v2/collection/" + url.PathEscape(id)
	resp, err := doRequest[any, DeleteResponse](ctx, c, "DELETE", url, nil)
	if err != nil {
		return false, err
	}
	return resp.Deleted, nil
}
