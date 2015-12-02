/*
slack-cli is simple command-line client for slack.

    $ go get github.com/takecy/slack-cli

First is
    $ slack-cli config <incoming web hook url>
Post message
  $ slack-cli post <message>
Post to specific channel
  $ slack-cli post -c general <message>
*/
package main
