package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBag_Contains(t *testing.T) {
	t.Run("direct", func(t *testing.T) {
		bags := make(Container)
		bags["light red"] = []string{"bright white", "muted yellow"}

		assert.True(t, bags.Has("light red", "bright white"))
		assert.False(t, bags.Has("light red", "dark brown"))
	})

	t.Run("nested one level", func(t *testing.T) {
		bags := make(Container)
		bags["light red"] = []string{"bright white", "muted yellow"}
		bags["muted yellow"] = []string{"shiny gold"}

		assert.True(t, bags.Has("light red", "shiny gold"))
		assert.False(t, bags.Has("light red", "dark olive"))
	})

	t.Run("nested two levels", func(t *testing.T) {
		bags := make(Container)
		bags["light red"] = []string{"bright white", "muted yellow"}
		bags["muted yellow"] = []string{"yuted bellow"}
		bags["yuted bellow"] = []string{"shiny gold"}

		assert.True(t, bags.Has("light red", "shiny gold"))
		assert.False(t, bags.Has("light red", "dark olive"))
	})
}
