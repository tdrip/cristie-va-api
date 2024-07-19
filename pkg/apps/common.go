package apps

import "fmt"

type UpdateUI func(msg string)

func DefaultUpdateUI(msg string) {
	fmt.Println(msg)
}
