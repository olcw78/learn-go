package greeting

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprintf("Hello, %s", name)
	return msg
}
