package memex

// EntityState represents visibility of space/media
type EntityState int16

const (
	// Visible spaces/media are visible to user
	Visible EntityState = 0
	// Trashed spaces will be removed after one month
	Trashed EntityState = 1
)
