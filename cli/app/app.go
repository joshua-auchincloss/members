package app

import (
	"github.com/rs/zerolog/log"

	"os"

	"github.com/urfave/cli/v2"
)

var cmds = []*cli.Command{
	&cluster,
	&database,
}

func closer() {
	os.Exit(0)
}

func BuildApp() *cli.App {
	log.Print("ok")
	app := cli.NewApp()
	app.Name = "mm"
	app.Usage = ""
	app.Version = "0.1"
	app.Suggest = true
	app.EnableBashCompletion = true
	app.Commands = cmds
	return app
}
