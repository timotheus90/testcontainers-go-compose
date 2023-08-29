package main

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	tccompose "github.com/testcontainers/testcontainers-go/modules/compose"
)

func TestCompose(t *testing.T) {
	compose, err := tccompose.NewDockerCompose(filepath.Join("testdata", "docker-compose.yml"))
	assert.NoError(t, err, "NewDockerCompose()")

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	assert.NoError(t, compose.Up(ctx, tccompose.Wait(true)), "compose.Up()")

	// do some testing here
}
