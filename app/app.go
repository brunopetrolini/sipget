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

func getServers(c *cli.Context) {
	url := c.String("url")

	servers, error := net.LookupNS(url)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Servers for %s\n", url)
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

// Generate returns an instance of the CLI app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "sipget"
	app.Usage = "Retrieves IPs and Server Names from a web address"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Retrieves the public IP address of a web address",
			Flags:  flags,
			Action: getIPs,
		},
		{
			Name:   "server",
			Usage:  "Retrieves the public server names of a web address",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}
