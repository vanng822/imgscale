package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatorTrue(t *testing.T) {
	v := defaultValidator{}
	assert.True(t, v.Validate("kth.jpg"))
	assert.True(t, v.Validate("http://127.0.0.1:8080/img/original/kth.jpg"))
	assert.True(t, v.Validate("350px-cf910efc-d69a-47ac-a5d2-7a66b72b890f.jpg"))
}


func TestValidatorFalse(t *testing.T) {
	v := defaultValidator{}
	assert.False(t, v.Validate(".."))
	assert.False(t, v.Validate("http://127.0.0.1:8080/img/../../../ect/pwd"))
	assert.False(t, v.Validate("350px-cf910efc-d69a-47ac-a5d2-7a..66b72b890f.jpg"))
}

