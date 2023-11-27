package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEditor(t *testing.T) {
	t.Run("creates a new editor with args", func(t *testing.T) {
		r := require.New(t)
		edit, err := NewEditor([]string{defaultEditor})
		r.NoError(err)
		r.Equal(defaultEditor, edit.Args[0])
	})

	t.Run("creates a new editor with default args", func(t *testing.T) {
		testDefaultShell := os.Getenv("SHELL")
		r := require.New(t)
		edit, err := NewEditor([]string{})
		r.NoError(err)
		r.Equal(testDefaultShell, edit.Args[0])
		r.Equal("-c", edit.Args[1])
		r.Equal("EDITOR", edit.Args[2])
	})
}
