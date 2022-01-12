package main

import (
	"docker/container"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var runCommand = cli.Command{
	Name:      "run",
	ShortName: "",
	Aliases:   nil,
	Usage: "Create a container with namespace and cgroups limit " +
		"docker run -ti [command]",
	UsageText:    "",
	Description:  "",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Before:       nil,
	After:        nil,
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
	OnUsageError: nil,
	Subcommands:  nil,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "ti",
			Usage:       "enable tty",
			EnvVar:      "",
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			Destination: nil,
		},
	},
	SkipFlagParsing:        false,
	SkipArgReorder:         false,
	HideHelp:               false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}
var initCommand = cli.Command{
	Name:         "init",
	ShortName:    "",
	Aliases:      nil,
	Usage:        "Init container process run user's process in container. Do not call it outside",
	UsageText:    "",
	Description:  "",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Before:       nil,
	After:        nil,
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)

		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
	OnUsageError:           nil,
	Subcommands:            nil,
	Flags:                  nil,
	SkipFlagParsing:        false,
	SkipArgReorder:         false,
	HideHelp:               false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
