package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/takecy/slack-cli/slacker"
)

const usage = `
Usage:
  $ slack-cli config <incoming web hook URL>
  $ slack-cli post [-c <other channel>] <message>

Commands:
  config  Set configuration
  post    Post massage to slack
`

func main() {
	flag.Usage = func() { usageAndExit() }
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	if flag.Args()[0] == "config" {
		if len(flag.Args()) < 2 {
			fmt.Fprint(os.Stderr, "Error: URL required. \n")
			flag.Usage()
			return
		}

		slacker.CreateConfig()
		c := &slacker.Config{
			URL: flag.Args()[1],
		}
		slacker.WriteConfig(c)
		fmt.Fprint(os.Stdout, "Success: register config \n")
		return
	}

	if flag.Args()[0] == "post" {
		if flag.NArg() < 2 {
			flag.Usage()
			return
		}

		c, err := slacker.ReadConfig()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: Failed read config. \n Require [config] command first. \n")
			flag.Usage()
			return
		}

		fs := flag.NewFlagSet(flag.Args()[0], flag.ExitOnError)
		ch := fs.String("c", "", "Other Channel")
		fs.Parse(flag.Args()[1:])

		slack := &slacker.Slack{C: c}
		err = slack.Post(slacker.Message{
			Text:    flag.Args()[flag.NArg()-1],
			Channel: "#" + *ch,
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed Post. \n %+v \n", err)
			return
		}
		return
	}

	fmt.Fprint(os.Stderr, "Error: Bad command \n")
	flag.Usage()
}

func usageAndExit() {
	fmt.Fprintf(os.Stdout, "%s \n", usage)
	os.Exit(1)
}
