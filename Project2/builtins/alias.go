package builtins

import (
	"fmt"
	"io"
)

// var aliases = make(map[string]string)

func HandleAlias(w io.Writer, args ...string) error {
	// Implementation of the function goes here
	fmt.Fprintln(w, "HandleAlias function called with arguments:", args)
	return nil
}

// func HandleAlias(w io.Writer, args ...string) error {
// 	switch len(args) {
// 	case 0:
// 		return fmt.Errorf("%w: expected one or two arguments (alias, command)", ErrInvalidArgCount)
// 	case 1:
// 		// print alias
// 		if val, ok := aliases[args[0]]; ok {
// 			fmt.Println(val)
// 			return nil
// 		}
// 		return fmt.Errorf("alias %s not found", args[0])
// 	case 2:
// 		// set alias
// 		aliases[args[0]] = args[1]
// 		return nil
// 	default:
// 		return fmt.Errorf("%w: expected one or two arguments (alias, command)", ErrInvalidArgCount)
// 	}
// }
