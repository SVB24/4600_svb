package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Echo prints its arguments to the standard output.
func Echo(w io.Writer, args ...string) error {
	// Join the arguments into a single string with spaces.
	echoString := strings.Join(args, " ")

	// Print the echo string to the standard output.
	_, err := fmt.Fprintln(w, echoString)
	return err
}
