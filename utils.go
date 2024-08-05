/*
@Time : 2024/8/4 下午4:35
@Author : ljn
@File : utils
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
)

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// CloseFile attempts to close the passed file
// or panics with the actual er	ror
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// WriteToFile creates a file and writes content to it
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}

func ExecuteCommand(name string, subname string, args ...string) error {
	args = append([]string{subname}, args...)
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}
