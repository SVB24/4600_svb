package builtins

import (
	"fmt"
	"os"
)

// MakeDirectory creates a new directory with the specified name.
func MakeDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("mkdir: missing operand")
	}

	for _, dir := range args {
		err := os.Mkdir(dir, 0777) // Create directory with full permissions
		if err != nil {
			return err
		}
	}

	return nil
}
