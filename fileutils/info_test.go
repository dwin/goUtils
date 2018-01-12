package fileutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirectorySize(t *testing.T) {
	size, err := DirectorySize("../.")
	assert.Nil(t, err)
	assert.NotZero(t, size)
}
