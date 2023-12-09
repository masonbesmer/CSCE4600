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
	//check if alias exists
	if _, ok := aliases[args[0]]; !ok {
		return fmt.Errorf("alias %s does not exist", args[0])
	}
	delete(aliases, args[0])
	return nil
}
