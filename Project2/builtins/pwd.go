package builtins

import (
	"fmt"
	"io"
	"os"
)

// pwd prints the current working directory
func Pwd(w io.Writer, args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s", wd)
	return nil
}
