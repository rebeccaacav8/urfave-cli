package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "fix-duplication-issue",
		Usage: "Demonstrate fix for StringSliceFlag default value duplication",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "flag",
				Aliases: []string{"f"},
				Usage:   "A string slice flag with default values",
				Value:   cli.NewStringSlice("default1", "default2"),
			},
		},
		Action: func(c *cli.Context) error {
			values := c.StringSlice("flag")
			fmt.Printf("Flag values: %v\n", values)
			
			// Expected behavior: 
			// If run with --flag value1 --flag value2, output should be ["value1", "value2"]
			// NOT ["default1", "default2", "value1", "value2"]
			
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}