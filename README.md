# TIMP, Type IMProve

A simple terminal based typing game. Created just for fun, not extensively tested, USE AT OWN RISK.

The game creates a hidden folder named .timp under the home folder where all data is stored. Shouldnt really be an issue, but there is no limit to amount of texts or history so just keep that in mind. 

## How to play

The cli has a --help flag for all commands.
To get started with a bunch of random wikipedia texts, run `$ timp wikiRand -t 20` to get some texts, then run `$ timp play`.

## Install

* Install and configure golang
* Configure PATH variable, for example `export PATH=$(go env GOPATH)/bin:$PATH`
* Clone repo into golang workdir
* Change dir to repo, run `go get -d ./...`, this installs all dependencies
* run `go install`
