package memex

import (
	"encoding/json"
	"fmt"
	"time"
)

// Link represents link between spaces
type Link struct {

	// Unique identifier
	MUID *string `json:"muid,omitempty"`
	// Creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp of last update
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Visibility state
	State *EntityState `json:"state,omitempty"`
	// Owner user ID
	OwnerID *int64 `json:"owner_id,omitempty"`
	// Index that is used for sorting of links in space
	Order *int64 `json:"order,omitempty"`
	// Origin space MUID
	OriginSpaceMUID *string `json:"origin_space_muid,omitempty"`
	// Target space MUID
	TargetSpaceMUID *string `json:"target_space_muid,omitempty"`
}

type linksResponse struct {
	Links []Link `json:"links"`
}

type linksRequest struct {
	Links       []*Link `json:"links"`
	RemoveToken *string `json:"remove_token,omitempty"`
}

// GetSpaceLinks returns links from space
func (spaces *Spaces) GetSpaceLinks(muid string) (*[]Link, error) {
	path := fmt.Sprintf("/spaces/%v/links", muid)
	var responseObject linksResponse
	_, requestError := spaces.perform("GET", path, nil, &responseObject)
	if requestError != nil {
		return nil, requestError
	}
	return &responseObject.Links, nil
}

// UpdateLinks updates multiple links
func (spaces *Spaces) UpdateLinks(array []*Link, removeToken *string) error {
	message := &linksRequest{
		Links:       array,
		RemoveToken: removeToken,
	}
	body, serializationError := json.Marshal(message)
	if serializationError != nil {
		return serializationError
	}
	path := fmt.Sprintf("/links/multiple")
	var responseObject emptyResponse
	_, requestError := spaces.perform("POST", path, body, &responseObject)
	if requestError != nil {
		return requestError
	}
	return nil
}
