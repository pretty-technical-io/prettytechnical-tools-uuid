package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID_CreateUUID(t *testing.T) {
	var id UUID

	assert.Equal(t, 0, len(id))

	err := id.Generate()
	assert.Equal(t, nil, err)

	assert.Equal(t, 36, len(id))
}

func TestUUID_CreateUUIDEmbedded(t *testing.T) {
	type wrapper struct {
		id UUID
	}

	var w wrapper
	assert.Equal(t, 0, len(w.id))

	err := w.id.Generate()
	assert.Equal(t, nil, err)
	assert.Equal(t, 36, len(w.id))
}

func TestUUID_Validate(t *testing.T) {
	tests := []struct {
		name string
		id   UUID
		err  error
	}{
		{
			name: "Success",
			id:   "7d91a38f-6b07-4b47-ba9e-a53670fdcc27",
			err:  nil,
		}, {
			name: "EmptyFailure",
			id:   "",
			err:  ErrInvalidUUID,
		}, {
			name: "InvalidFailure",
			id:   "1",
			err:  ErrInvalidUUID,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.id.Validate()
			assert.Equal(t, test.err, err)
		})
	}
}

func TestNew(t *testing.T) {
	id, err := New()
	assert.Equal(t, nil, err)
	assert.Equal(t, 36, len(id))
}

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		id   string
		err  error
	}{
		{
			name: "Success",
			id:   "7d91a38f-6b07-4b47-ba9e-a53670fdcc27",
			err:  nil,
		}, {
			name: "EmptyFailure",
			id:   "",
			err:  ErrInvalidUUID,
		}, {
			name: "InvalidFailure",
			id:   "1",
			err:  ErrInvalidUUID,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Parse(test.id)
			assert.Equal(t, test.err, err)
		})
	}
}
