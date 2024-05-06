package builtins

import (
	"fmt"
	"os"
)

// Touch creates an empty file with the specified name.
func Touch(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("touch: missing operand")
	}

	for _, file := range args {
		// Create an empty file.
		_, err := os.Create(file)
		if err != nil {
			return err
		}
	}

	return nil
}
