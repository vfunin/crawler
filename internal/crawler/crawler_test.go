package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New(0, 0)
	assert.NotNil(t, c, "new request is nil")
	assert.Equal(t, c.GetCnt(), int64(1))
}

func Test_crawler_Crawl(t *testing.T) {

}

func Test_crawler_DecCnt(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.GetCnt(), int64(1))
	c.DecCnt()
	assert.Equal(t, c.GetCnt(), int64(0))
}

func Test_crawler_GetCnt(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.GetCnt(), int64(1))
	c.DecCnt()
	assert.Equal(t, c.GetCnt(), int64(0))
	c.IncCnt()
	assert.Equal(t, c.GetCnt(), int64(1))
}

func Test_crawler_IncCnt(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.GetCnt(), int64(1))
	c.IncCnt()
	assert.Equal(t, c.GetCnt(), int64(2))
}

func Test_crawler_IncMaxDepth(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.MaxDepth(), uint64(0))
	c.IncMaxDepth(uint64(1))
	assert.Equal(t, c.MaxDepth(), uint64(1))
}

func Test_crawler_ResultCh(t *testing.T) {
	c := New(0, 0)
	assert.NotEqual(t, c.ResultCh(), make(<-chan Result))
}
