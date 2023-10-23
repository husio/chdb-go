package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chdb-io/chdb-go"
)

func main() {
	query := flag.String("query", "SELECT version()", "Query to execute")
	format := flag.String("format", "CSV", "Query output format")
	path := flag.String("path", "", "Table persistence path")
	flag.Parse()

	var (
		result string
		err    error
	)
	if len(*path) > 0 {
		result, err = chdb.Session(*query, *format, *path)
	} else {
		result, err = chdb.Query(*query, *format)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
}
