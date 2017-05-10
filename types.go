package memex

// EntityState represents visibility of space/media
type EntityState int16

const (
	// Visible spaces/media are visible to user
	Visible EntityState = 0
	// Trashed spaces will be removed after one month
	Trashed EntityState = 1
)

// Environment represents service environment
type Environment int16

const (
	// Production environment server
	Production Environment = 0
	// Stage environment server
	Stage Environment = 1
	// Local environment server
	Local Environment = 2
	// Sandbox environment server
	Sandbox Environment = 3
)
