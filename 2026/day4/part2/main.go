package main

import (
	"log/slog"
	"strconv"

	"github.com/tboerger/advent-of-code/pkg"
)

func main() {
	pkg.Start(run)
}

func run(testing bool) string {
	var (
		result = 0
	)

	pkg.Parse(testing, func(rows []string) {
		for _, row := range rows {
			slog.Info(row)
		}
	})

	return strconv.Itoa(result)
}
