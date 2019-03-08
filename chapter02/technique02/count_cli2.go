package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)


func upAction (c *cli.Context) error {
	stop := c.Int("stop")
	for i := 0; i <= stop; i++ {
		fmt.Println(i)
	}
	return nil
}


func main() {
	app:= cli.NewApp()
	app.Usage = "Count up or down"

	/*stopFlag := cli.IntFlag{Name: "stop, s", Usage:"Value to count up to", Value: 20}
	upFlags := [] cli.Flag{
		stopFlag,
	}

	upCommand := cli.Command{
		Name: "up", ShortName:"u", Usage:"Count Up", Flags:upFlags, Action:,
	}
	*/

	app.Commands = []cli.Command {
		{Name: "up", ShortName: "u", Usage: "Count Up",
			Flags: []cli.Flag{
				cli.IntFlag{Name: "stop, s", Usage: "Value to count up to", Value: 10},
			},
			Action: func(c *cli.Context) error {
				stop := c.Int("stop")
				for i := 0; i <= stop; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{ Name: "down", ShortName: "d", Usage: "Count Down",
			Flags: [] cli.Flag {
				cli.IntFlag{Name: "start, s", Usage: "Start counting down from", Value:10 },
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
