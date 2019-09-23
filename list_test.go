package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	str := Hello()
	assert.Equal(t, str, "hello") // must be signed in
}
