package builtins

import (
	"fmt"
	"os"
)

// MoveFiles moves files or directories from source to destination.
func MoveFiles(args ...string) error {

	if len(args) < 2 {
		return fmt.Errorf("mv: missing operand")
	}

	source := args[0]
	destination := args[1]

	err := os.Rename(source, destination)
	if err != nil {
		return err
	}

	return nil
}
