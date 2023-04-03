package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "610", ToString(610))
	assert.Equal(t, "6", ToString(int8(6)))
	assert.Equal(t, "610", ToString(int16(610)))
	assert.Equal(t, "610", ToString(int32(610)))
	assert.Equal(t, "610", ToString(int64(610)))
	assert.Equal(t, "610", ToString("610"))
}
