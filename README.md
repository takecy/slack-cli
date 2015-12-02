# slack-cli
Simple command-line client for slack by golang

![](https://img.shields.io/badge/golang-1.5.1-blue.svg?style=flat)
[![GoDoc](https://godoc.org/github.com/takecy/slack-cli?status.svg)](https://godoc.org/github.com/takecy/slack-cli)

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
