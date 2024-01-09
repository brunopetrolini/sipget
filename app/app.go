package app

import "github.com/urfave/cli"

// Generate returns an instance of the CLI app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "sipget"
	app.Usage = "Retrieves IPs and Server Names from a web address"

	return app
}
