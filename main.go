package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"sqlbeat/beater"
)

func main() {
	err := beat.Run("sqlbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
