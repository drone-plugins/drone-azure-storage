package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	Config struct {
		AccountKey string
		Account    string
		Container  string
		Source     string
		Destination string
		Operation  string
	}

	Plugin struct {
		Config Config
	}
)

func (p Plugin) Exec() error {
	return p.execute(p.command())
}

func (p *Plugin) command() *exec.Cmd {
	args := []string{}

	switch p.Config.Operation {
	case "upload":
		args = append(args, "upload")
		args = append(args, fmt.Sprintf("--local-path=%s", p.Config.Source))
		args = append(
			args,
			fmt.Sprintf("--remote-path=%s/%s", p.Config.Container, p.Config.Destination)
		)
	case "download":
		args = append(args, "download")
		args = append(
			args,
			fmt.Sprintf("--remote-path=%s/%s", p.Config.Container, p.Config.Source)
		)
		args = append(args, fmt.Sprintf("--local-path=%s", p.Config.Destination))
	default:
		// TODO return error
		return nil
	}

	args = append(args, fmt.Sprintf("--storage-account=%s", p.Config.Account))
	args = append(
		args,
		fmt.Sprintf("--storage-account-key=%s", p.Config.AccountKey),
	)

	return exec.Command(
		"blobxfer",
		args...,
	)
}

func (p Plugin) execute(cmd *exec.Cmd) error {
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	p.trace(cmd)
	return cmd.Run()
}

func (p *Plugin) trace(cmd *exec.Cmd) {
	fmt.Println("+", strings.Replace(strings.Join(cmd.Args, " "), p.Config.AccountKey, "***", -1))
}
