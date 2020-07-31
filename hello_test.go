package golang_mod_demo

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestHello(t *testing.T) {
    assert.Equal(t, "Hello World v0.0.2", Hello())
}