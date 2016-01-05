package main

import (
	"fmt"
	"github.com/drone/drone-plugin-go/plugin"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	build     string
	buildDate string
)

type AzureBlobxfer struct {
	StorageAccountKey  string `json:"account_key"`
	StorageAccountName string `json:"storage_account"`
	Container          string `json:"container"`
	Source             string `json:"source"`
}

func main() {
	fmt.Printf("Drone Azure Storage Plugin built at %s\n", buildDate)
	workspace := plugin.Workspace{}
	vargs := AzureBlobxfer{}

	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.StorageAccountKey) == 0 {
		return
	}

	if len(vargs.Container) == 0 {
		return
	}

	cmd := command(vargs, workspace)
	cmd.Env = os.Environ()
	cmd.Dir = workspace.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	trace(cmd)
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

func command(s AzureBlobxfer, w plugin.Workspace) *exec.Cmd {

	args := []string{
		"--storageaccountkey",
		s.StorageAccountKey,
		s.StorageAccountName,
		s.Container,
		filepath.Join(w.Path, s.Source),
	}
	return exec.Command("blobxfer", args...)
}

// trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
