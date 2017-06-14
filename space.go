package memex

import (
	"encoding/json"
	"fmt"
	"time"
)

// SpaceType represents type of space
type SpaceType string

const (
	// Origin represents starting point space for user
	Origin SpaceType = "com.memex.origin"
	// Collection is set/list of links to other spaces
	Collection SpaceType = "com.memex.media.collection"
	// WebPage space represents web link URL
	WebPage SpaceType = "com.memex.media.webpage"
	// Image space represents image/diagram space
	Image SpaceType = "com.memex.media.image"
	// Text space represents textual data
	Text SpaceType = "com.memex.media.text"
)

// Space represents folder/text/everything
type Space struct {
	// Unique identifier
	MUID *string `json:"muid,omitempty"`
	// Creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp of last update
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Timestamp of last visit
	VisitedAt *time.Time `json:"visited_at,omitempty"`
	// Visibility state
	State *EntityState `json:"state,omitempty"`
	// Owner user ID
	OwnerID *int64 `json:"owner_id,omitempty"`
	// Type (eg. com.memex.media.collection, etc.)
	SpaceType *SpaceType `json:"type_identifier,omitempty"`
	// Caption
	Caption *string `json:"tag_label,omitempty"`
	// Tint color
	Color *string `json:"tag_color,omitempty"`
	// Set of media that represents space (eg webpage space is represented by url, thumbnail, summary)
	Representations *[]Media `json:"representations,omitempty"`
	// Unread flag (if user needs to be notified about changes)
	Unread *bool `json:"unread,omitempty"`
}

type spaceResponse struct {
	Space Space `json:"space"`
}

type spacesRequest struct {
	Spaces []*Space `json:"spaces"`
}

type markAsUnreadRequest struct {
	MUIDs []*string `json:"space_MUIDs"`
}

// RepresentationWithType returns representation with specified media type
func (space *Space) RepresentationWithType(mediaType MediaType) *Media {
	if space.Representations == nil {
		return nil
	}
	for _, media := range *space.Representations {
		if media.MediaType != nil && *media.MediaType == mediaType {
			return &media
		}
	}
	return nil
}

// GetSpace returns space with representations
func (spaces *Spaces) GetSpace(muid string) (*Space, error) {
	path := fmt.Sprintf("/spaces/%v", muid)
	var responseObject spaceResponse
	_, requestError := spaces.perform("GET", path, nil, &responseObject)
	if requestError != nil {
		return nil, requestError
	}
	return &responseObject.Space, nil
}

// UpdateSpaces updates spaces
func (spaces *Spaces) UpdateSpaces(array []*Space, ownerID int64) error {
	message := &spacesRequest{
		Spaces: array,
	}
	body, serializationError := json.Marshal(message)
	if serializationError != nil {
		return serializationError
	}
	path := fmt.Sprintf("/spaces/multiple")
	var responseObject spaceResponse
	_, requestError := spaces.perform("POST", path, body, &responseObject)
	if requestError != nil {
		return requestError
	}
	return nil
}

// UpdateSpace updates single space
func (spaces *Spaces) UpdateSpace(space *Space) error {
	array := []*Space{space}
	if space.OwnerID == nil {
		return fmt.Errorf("Missing ownerID")
	}
	return spaces.UpdateSpaces(array, *space.OwnerID)
}

// MarkSpacesAsUnread marks spaces as unread
func (spaces *Spaces) MarkSpacesAsUnread(muids []*string) error {
	message := &markAsUnreadRequest{
		MUIDs: muids,
	}
	body, serializationError := json.Marshal(message)
	if serializationError != nil {
		return serializationError
	}
	path := fmt.Sprintf("/spaces/mark-as-unread")
	var responseObject spaceResponse
	_, requestError := spaces.perform("POST", path, body, &responseObject)
	if requestError != nil {
		return requestError
	}
	return nil
}
