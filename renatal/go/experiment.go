package gobind

import "fmt"

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
