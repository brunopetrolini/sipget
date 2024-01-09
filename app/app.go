package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func getIPs(c *cli.Context) {
	url := c.String("url")

	ips, error := net.LookupIP(url)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("IPs for %s\n", url)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Generate returns an instance of the CLI app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "sipget"
	app.Usage = "Retrieves IPs and Server Names from a web address"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Retrieves the public IP address of a web address",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url",
					Value: "google.com",
				},
			},
			Action: getIPs,
		},
	}

	return app
}
