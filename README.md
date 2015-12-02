# slack-cli
Simple command-line client for slack by golang

## Features
* Post message only
* Use Incoming Web hook service on slack

## Install
```shell
$ go get github.com/takecy/slack-cli
```

## Usage
```shell
$ slack-cli config <incoming web hook URL>
$ slack-cli post <message>
```
post to specific channel (not prefix [#])
```shell
$ slack-cli post -c general <message>
```

## Remove
```shell
$ rm $GOPATH/bin/slack-cli
$ rm -r $HOME/.slack_cli
```
