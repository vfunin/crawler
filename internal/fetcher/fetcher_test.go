package fetcher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	f := New(time.Duration(10))
	assert.NotNil(t, f, "new fetcher is nil")
}
