package imagick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestBooleanTypeTrue(t *testing.T) {
	tt := BooleanType(1)
	assert.True(t, tt.True())
	ff := BooleanType(0)
	assert.False(t, ff.True())
}


func TestBooleanTypeGoBool(t *testing.T) {
	tt := BooleanType(1)
	assert.True(t, tt.GoBool())
	ff := BooleanType(0)
	assert.False(t, ff.GoBool())
}