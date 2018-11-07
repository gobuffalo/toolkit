package models

import (
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"regexp"
)

type Tool struct {
	ID               uuid.UUID     `json:"id" db:"id"`
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at" db:"updated_at"`
	Name             string        `json:"name" db:"name"`
	NameWithOwner    string        `json:"name_with_owner" db:"name_with_owner"`
	URL              string        `json:"url" db:"url"`
	DiscoveryService string        `json:"discovery" db:"discovery"`
	Stars            int           `json:"stars" db:"stars"`
	Watchers         int           `json:"watchers" db:"watchers"`
	Forks            int           `json:"forks" db:"forks"`
	Description      nulls.String  `json:"description" db:"description"`
	Readme           nulls.String  `json:"readme" db:"readme"`
	Topics           slices.String `json:"topics" db:"topics"`
	License          License       `json:"license" has_one:"license"`
}

// String is not required by pop and may be deleted
func (t Tool) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tools is not required by pop and may be deleted
type Tools []Tool

// String is not required by pop and may be deleted
func (t Tools) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Tool) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
		&validators.StringIsPresent{Field: t.NameWithOwner, Name: "NameWithOwner"},
		&validators.StringIsPresent{Field: t.URL, Name: "URL"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Tool) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Tool) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
