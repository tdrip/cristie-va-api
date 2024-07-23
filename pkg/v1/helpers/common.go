package helpers

import (
	"fmt"
	"time"

	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
)

const layout = "01-02-2006-15-04-05"

type NamingConvention func(prefix string, fd bs.Client, t time.Time) string
type UpdateUI func(msg string)

func DefaultNamingConvention(prefix string, fd bs.Client, t time.Time) string {
	return fmt.Sprintf("%s%s-%s", prefix, fd.Name, t.Format(layout))
}

func DefaultUpdateUI(msg string) {
	fmt.Println(msg)
}
