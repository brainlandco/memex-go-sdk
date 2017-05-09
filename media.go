package memex

import (
	"encoding/json"
	"fmt"
	"time"
)

// MediaType tells what kind of media it is (semantically), mime is in metadata
type MediaType string

const (
	// Source media type is basic data source of media
	Source MediaType = "source"
	// Reference media contains indirect pointer to media data (eg URL)
	Reference MediaType = "reference"
	// Preview represents more abstract visual representation of space
	Preview MediaType = "preview"
	// Summary represents more abstract textual representation of space
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
	MUID                 *string        `json:"muid,omitempty"`
	CreatedAt            *time.Time     `json:"created_at,omitempty"`
	UpdatedAt            *time.Time     `json:"updated_at,omitempty"`
	State                EntityState    `json:"state"`
	OwnerID              *int64         `json:"owner_id,omitempty"`
	Metadata             *string        `json:"metadata,omitempty"`
	MediaType            MediaType      `json:"type"`
	DataState            MediaDataState `json:"data_state"`
	EmbededData          []byte         `json:"embeded_data,omitempty"`
	DataDownloadURL      *string        `json:"data_download_url,omitempty"`
	DataUploadURL        *string        `json:"data_upload_url,omitempty"`
	RepresentedSpaceMUID *string        `json:"represented_space_muid,omitempty"`
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
