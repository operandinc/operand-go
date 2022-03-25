package operand

import (
	"context"
	"net/url"
	"strconv"
)

// ListAtoms returns a list of all configured atoms.
func (c *Client) ListAtoms(ctx context.Context, collectionId *string, groupId *string) ([]Atom, error) {
	atoms := []Atom{}
	for {
		params := url.Values{
			"limit": []string{"100"},
			"offset": []string{
				strconv.Itoa(len(atoms)),
			},
		}
		if collectionId != nil {
			params.Add("collection", *collectionId)
		}
		if groupId != nil {
			params.Add("group", *groupId)
		}
		url := c.endpoint + "/v2/atom?" + params.Encode()
		resp, err := doRequest[any, ListResponse[Atom]](ctx, c, "GET", url, nil)
		if err != nil {
			return nil, err
		}
		atoms = append(atoms, resp.Items...)
		if !resp.More {
			break
		}
	}
	return atoms, nil
}

// GetAtom returns an atom by ID.
func (c *Client) GetAtom(ctx context.Context, id string) (*Atom, error) {
	url := c.endpoint + "/v2/atom/" + url.PathEscape(id)
	resp, err := doRequest[any, Atom](ctx, c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateAtomRequest is used to create an atom.
type CreateAtomRequest struct {
	GroupID    string     `json:"groupId"`
	Content    string     `json:"content"`
	Properties Properties `json:"properties"`
}

// CreateAtom creates a new atom.
func (c *Client) CreateAtom(ctx context.Context, req *CreateAtomRequest) (*Atom, error) {
	resp, err := doRequest[CreateAtomRequest, Atom](ctx, c, "POST", c.endpoint+"/v2/atom", req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateAtomRequest is used to update an atom.
type UpdateAtomRequest struct {
	Properties Properties `json:"properties"`
}

// UpdateAtom updates an atom.
func (c *Client) UpdateAtom(ctx context.Context, id string, req *UpdateAtomRequest) (*Atom, error) {
	url := c.endpoint + "/v2/atom/" + url.PathEscape(id)
	resp, err := doRequest[UpdateAtomRequest, Atom](ctx, c, "PUT", url, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteAtom deletes an atom.
func (c *Client) DeleteAtom(ctx context.Context, id string) (bool, error) {
	url := c.endpoint + "/v2/atom/" + url.PathEscape(id)
	resp, err := doRequest[any, DeleteResponse](ctx, c, "DELETE", url, nil)
	if err != nil {
		return false, err
	}
	return resp.Deleted, nil
}
