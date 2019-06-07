#!/bin/sh

go build -ldflags "-X github.com/Karm/trg/cmd.version=`git describe --tags`" -o trg
