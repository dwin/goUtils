package fileutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirSize(t *testing.T) {
	size, err := DirSize("../.")
	assert.Nil(t, err)
	assert.NotZero(t, size)
}

func TestGetDirFilenames(t *testing.T) {
	files, err := DirFilenames("../.")
	assert.Nil(t, err)
	fmt.Println(err)
	assert.NotEmpty(t, files)
}
