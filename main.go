package main

import (
	"fmt"
	"os"

	"github.com/tonicbupt/caprice"
	"github.com/urfave/cli/v2"
)

func chooseCandidate(apikey string, candidates []string) string {
	if len(candidates) == 1 {
		return candidates[0]
	}

	client := caprice.TrueRNG(apikey)
	numbers, err := client.GenerateIntegers(1, 0, len(candidates)-1, false)
	// Never seen such stupid code
	// Error() method is not assigned to *Error but just Error!!!!
	if err.Code != 0 {
		return ""
	}
	if len(numbers) != 1 {
		return ""
	}
	return candidates[numbers[0]]
}

func main() {
	app := cli.NewApp()
	app.Name = "Aaron"
	app.Usage = "Aaron helps you to make decision!"
	app.Version = "0x336ae74f"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "apikey",
			Value:   "",
			Usage:   "API Key from random.org",
			EnvVars: []string{"AARON_APIKEY"},
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Println("Aaron needs options")
			return nil
		}
		if c.String("apikey") == "" {
			fmt.Println("Aaron needs apikey set")
			return nil
		}
		ret := chooseCandidate(c.String("apikey"), c.Args().Slice())
		fmt.Println(fmt.Sprintf("Aaron says %s", ret))
		return nil
	}

	app.Run(os.Args)
}
