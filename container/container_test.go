package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer(t *testing.T) {
	c := NewContainer()

	c.Singleton("a", func(c *Container) interface{} {
		c.Singleton("b", func(c *Container) interface{} {
			return "b"
		})

		return "a"
	})
	c.Instance("c", "c")

	assert.Equal(t, "a", c.MustMake("a").(string))
	assert.Equal(t, "b", c.MustMake("b").(string))
	assert.Equal(t, "c", c.MustMake("c").(string))
	assert.Panics(t, func() {
		c.MustMake("d")
	})
}
