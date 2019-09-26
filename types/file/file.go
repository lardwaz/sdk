package file

import (
	"encoding/json"
	"time"
)

// File represents a file which may be transferred or saved.
type File struct {
	ID          string    `json:"id"`
	ContentID   string    `json:"content_id"`
	ContentType string    `json:"content_type"`
	Filename    string    `json:"filename"`
	MIMEType    string    `json:"mimetype"`
	Filepath    string    `json:"filepath"`
	Size        int       `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
}

// New creates a new instance of File from a json string
func New(v string) (*File, error) {
	f := File{}

	if err := json.Unmarshal([]byte(v), &f); err != nil {
		return nil, err
	}

	return &f, nil
}
