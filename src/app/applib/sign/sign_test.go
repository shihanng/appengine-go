package sign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGreeting(t *testing.T) {
	g := NewGreeting("test")
	assert.Equal(t, "test", g.Content)
}
