package main

import (
	"context"
	"fmt"
	tccompose "github.com/testcontainers/testcontainers-go/modules/compose"
	"log"
	"runtime"
	"testing"
)

func TestSomething(t *testing.T) {
	fmt.Println("Go version:", runtime.Version())
	// Setup Phase: Initialize TestContainer
	compose, err := tccompose.NewDockerCompose("testdata/docker-compose.yml")
	if err != nil {
		log.Fatalf("Could not create DockerCompose: %s", err)
		//os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Will be called upon exiting TestMain

	println("Starting compose")

	if err := compose.Up(ctx, tccompose.Wait(true)); err != nil {
		log.Fatalf("Could not create DockerCompose: %s", err)
		//os.Exit(1)
	}

	// Run all the test suites
	t.Log("testing something")

	// Teardown Phase: Clean up TestContainer
	if err := compose.Down(context.Background(), tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal); err != nil {
		log.Fatalf("Could not create DockerCompose: %s", err)
		//os.Exit(1)
	}

	// Exit with code from test suites
	//os.Exit(0)
}
