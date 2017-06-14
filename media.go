package memex

import (
	"encoding/json"
	"fmt"
	"time"
)

// MediaType tells what kind of media it is (semantically), mime is in metadata
type MediaType string

const (
	// Source of data (every other representation can be derived from it).
	Source MediaType = "source"
	// Reference is link to source
	Reference MediaType = "reference"
	// Preview is visual/graphical abstraction of source/reference
	Preview MediaType = "preview"
	// Summary is textual abstraction of source/reference
	Summary MediaType = "summary"
)

// MediaDataState tells if media data is valid (downloadable) or not
type MediaDataState int16

const (
	// WaitingForNewUploadURL represents state where media is waiting for renewed upload url
	WaitingForNewUploadURL MediaDataState = 0
	// ReadyForDataUpload media has valid upload url and waits for uplod
	ReadyForDataUpload MediaDataState = 1
	// DataValid represents state where embedData or data at downloadURL is valid
	DataValid MediaDataState = 2
)

// Media represents folder/text/everything
type Media struct {
	// Unique identifier
	MUID *string `json:"muid,omitempty"`
	// Creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Last update timestamp
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Visibility state
	State EntityState `json:"state"`
	// JSON encodec dictionary of media metadata eg. size, encoding, etc.
	Metadata *string `json:"metadata,omitempty"`
	// Type of media
	MediaType MediaType `json:"type"`
	// Owner user ID
	OwnerID *int64 `json:"owner_id,omitempty"`
	// If media represents any space then its MUID is present
	RepresentedSpaceMUID *string `json:"represented_space_muid,omitempty"`
	// Validity of media data
	DataState MediaDataState `json:"data_state"`
	// Embed media binary data (only if small enough, otherwise use dataDownloadURL and dataUploadURL)
	EmbededData []byte `json:"embeded_data,omitempty"`
	// Download url for data (exclusive with embedData)
	DataDownloadURL *string `json:"data_download_url,omitempty"`
	// Upload link for new data. After data is uploaded it is needed to call mark media as uploaded function.
	DataUploadURL *string `json:"data_upload_url,omitempty"`
}

type mediaRequest struct {
	Media []*Media `json:"media"`
}

type emptyResponse struct {
}

// UpdateMedia updates multiple media
func (spaces *Spaces) UpdateMedia(array []*Media) error {
	message := &mediaRequest{
		Media: array,
	}
	body, serializationError := json.Marshal(message)
	if serializationError != nil {
		return serializationError
	}
	path := fmt.Sprintf("/media/multiple")
	var responseObject emptyResponse
	_, requestError := spaces.perform("POST", path, body, &responseObject)
	if requestError != nil {
		return requestError
	}
	return nil
}
