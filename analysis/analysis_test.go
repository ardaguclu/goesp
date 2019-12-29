package analysis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	a := New()
	assert.NotNil(t, a.PathRgx)
	assert.NotNil(t, a.StackRgx)
	assert.NotNil(t, a.HeapRgx)
}
