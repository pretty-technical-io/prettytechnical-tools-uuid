package uuid

import (
	"errors"

	"github.com/google/uuid"
)

// Errors.
var (
	ErrUUIDNotGenerated = errors.New("the uuid could not be generated")
	ErrInvalidUUID      = errors.New("the id is not a valid uuid")
)

// UUID is a 128 bit (16 byte) Universal Unique Identifier as defined in RFC.
type UUID string

// Generate generates a new UUID.
func (u *UUID) Generate() error {
	id, err := uuid.NewV7()
	if err != nil {
		return ErrUUIDNotGenerated
	}

	*u = UUID(id.String())

	return nil
}

// Validate confirm that the UUID is valid.
func (u UUID) Validate() error {
	_, err := uuid.Parse(string(u))
	if err != nil {
		return ErrInvalidUUID
	}

	return nil
}

// String returns the uuid converted to string.
func (u UUID) String() string {
	return string(u)
}

// New returns a new UUID.
func New() (UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", ErrUUIDNotGenerated
	}

	return UUID(id.String()), nil
}

// Parse confirm that the id is valid and returns the UUID.
func Parse(id string) (UUID, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return "", ErrInvalidUUID
	}

	return UUID(u.String()), nil
}
