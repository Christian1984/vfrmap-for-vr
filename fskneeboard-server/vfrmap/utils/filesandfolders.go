package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func OpenExplorer(folder string) error {
	exe, exeErr := os.Executable()

	if exeErr == nil {
		exePath := filepath.Dir(exe)
		autosavePath := exePath + "\\" + folder

		cmd := exec.Command("explorer", autosavePath)
		cmd.Run()
	}

	return exeErr
}