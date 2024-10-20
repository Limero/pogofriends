package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatFriendCodes(t *testing.T) {
	codes := []string{
		"012345678900",
		"0123 4567 8900",
	}
	expected := []string{
		"0123 4567 8900",
		"0123 4567 8900",
	}
	assert.Equal(t, expected, formatFriendCodes(codes))
	assert.NotEqual(t, expected, codes) // confirm it made a copy
}

func Test_findFriendCodes(t *testing.T) {
	content := `
	012345678900
	0123 4567 8900
	`
	expected := []string{
		"012345678900",
		"0123 4567 8900",
	}
	assert.Equal(t, expected, findFriendCodes(content))
}
