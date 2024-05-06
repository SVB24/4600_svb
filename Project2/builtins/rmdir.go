package builtins

import (
	"fmt"
	"os"
)

// RemoveDirectory removes the specified directory.
func RemoveDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("rmdir: missing operand")
	}

	for _, dir := range args {
		err := os.Remove(dir)
		if err != nil {
			return err
		}
	}

	return nil
}
