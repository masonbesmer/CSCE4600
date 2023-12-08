package builtins

import (
	"fmt"
	"io"
)

// Unalias removes an alias from the alias map.
func Unalias(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no alias provided")
	}
	delete(aliases, args[0])
	return nil
}
