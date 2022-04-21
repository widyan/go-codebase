package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailOnError(t *testing.T) {
	var err error
	FailOnError(err, "")
	assert.NoError(t, err)
}
