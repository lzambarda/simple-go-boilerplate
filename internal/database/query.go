package database

import (
	"embed"
	"fmt"
)

// Make sure that queries are embedded in the built binary and are easily
// accessible at development time.
//
// This setup can be tedious to maintain, if you like it and would like an
// automated version, check out this:
// https://gist.github.com/lzambarda/2767e6480ca751ca34cf085b5a44c440

//go:embed query
var queryFS embed.FS

//nolint:gochecknoglobals // Fine in this example.
var queries map[string]string

const (
	queryFooName = "query/foo.sql"
)

//nolint:gochecknoinits // Sometimes they can be helpful.
func init() {
	queries = map[string]string{}

	b, err := queryFS.ReadFile(queryFooName)
	if err != nil {
		panic(fmt.Errorf("reading %s: %w", queryFooName, err))
	}

	queries[queryFooName] = string(b)
}
