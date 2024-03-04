package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindRootDir(t *testing.T) {
	t.Run("find root dir", func(t *testing.T) {
		cwd, _ := os.Getwd()
		rootDir, _ := FindRootDir()
		if rootDir != filepath.Dir(cwd) {
			t.Fatalf("expected %v but got %v", filepath.Dir(cwd), rootDir)
		}
	})

	t.Run("cannot find root dir", func(t *testing.T) {
		tempDir := t.TempDir()
		err := os.Chdir(tempDir)
		if err != nil {
			t.Fatalf("Error changing directory: %v", err)
		}
		_, err = FindRootDir()
		if err == nil || err.Error() != "root directory not found" {
			t.Fatalf("expected root directory not found but got %v", err.Error())
		}
	})
}


