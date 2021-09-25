package main

import (
	"fmt"
	"os"

	"github.com/juju/errors"
	"github.com/martinisecurity/trusty/backend/trustymain"
)

const (
	rcError   = 1
	rcSuccess = 0
)

func main() {
	rc := rcSuccess

	app := trustymain.NewApp(os.Args[1:])

	err := app.Run(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", errors.ErrorStack(err))
		rc = rcError
	}
	app.Close()

	os.Exit(rc)
}
