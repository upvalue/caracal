package cmd

import (
	"fmt"
	"os"

	"github.com/segmentio/cli"
)

type commonFlags struct {
	ConfigPath string `flag:"-c,--config" help:"Path to configuration"`
}

type linksFlags struct {
	commonFlags
	_ struct{} `help:"Pretty print configuration"`
}

func configCommand(flags linksFlags) {
	cfg, err := loadConfig(flags.ConfigPath)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("# Links")
	for _, link := range cfg.Link {
		fmt.Printf("%s => %s\n", link.Name, link.Url)
	}
}

func CommandLine() {
	cli.Exec(cli.CommandSet{
		"config": cli.Command(configCommand),
		"serve":  cli.Command(serveCommand),
	})
}
