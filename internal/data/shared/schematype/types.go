package schematype

import (
	"fmt"

	"entgo.io/ent/dialect"
)

func Decimal(precision, scale int) map[string]string {
	return map[string]string{
		dialect.MySQL:    fmt.Sprintf("decimal(%d, %d)", precision, scale),
		dialect.Postgres: fmt.Sprintf("numeric(%d, %d)", precision, scale),
	}
}

func Date() map[string]string {
	return map[string]string{
		dialect.Postgres: "date",
		dialect.MySQL:    "date",
	}
}
