// +build windows

package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

// Go builds files on windows with an '.exe' suffix, so we need a few pieces of
// special logic to make sure things work there.

// buildRetool builds retool in a temporary directory and returns the path to
// the built binary
func buildRetool() (string, error) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", errors.Wrap(err, "unable to create temporary build directory")
	}
	output := filepath.Join(dir, "retool.exe")
	cmd := exec.Command("go", "build", "-o", output, ".")
	_, err = cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "unable to build retool binary")
	}
	return output, nil
}

func assertBinInstalled(t *testing.T, wd, bin string) {
	_, err := os.Stat(filepath.Join(wd, "_tools", "bin", bin+".exe"))
	if err != nil {
		t.Errorf("unable to find %s: %s", bin+".exe", err)
	}
}
