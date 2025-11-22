package main

import (
	"context"
	"log"
	"os"

	"github.com/isokolovskii/commitizen/cmd"
)

func main() {
	err := cmd.Lefthook().Run(context.Background(), os.Args)
	if err != nil {
		if err.Error() != "" {
			log.Fatalf("Error: %s", err)
		}
		os.Exit(1)
	}
}
