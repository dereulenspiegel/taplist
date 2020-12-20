package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDExtraction(t *testing.T) {

	for _, td := range []struct {
		id        string
		prefix    string
		numString string
	}{
		{
			id:        "brewfather:7",
			prefix:    "brewfather",
			numString: "7",
		},
		{
			id:        "7",
			prefix:    "",
			numString: "",
		},
		{
			id:        "",
			prefix:    "",
			numString: "",
		},
	} {
		pref, id := extractID(td.id)

		assert.Equal(t, td.prefix, pref)
		assert.Equal(t, td.numString, id)
	}
}
