package main_test

import (
	"testing"

	"github.com/kommunkod/restclone/pkg/server"
)

func StartServer(t *testing.T) {
	t.Log("Starting RestClone Server....")
	server.Run()
}
