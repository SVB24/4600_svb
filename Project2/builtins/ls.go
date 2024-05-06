package builtins

import (
	"fmt"
	"io"
	"io/ioutil"
)

// ListDirectoryContents lists the contents of the current directory.
func ListDirectoryContents(w io.Writer) error {
	// Read the contents of the current directory.
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	// Print the names of the files and directories to the standard output.
	for _, file := range files {
		_, err := fmt.Fprintf(w, "%s\t", file.Name())
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintln(w) // Add a newline after listing all files.
	return err
}
