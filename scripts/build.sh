#!/usr/bin/env sh

GOPATH=/vagrant/go go build -o /vagrant/bin/controller /vagrant/go/src/main/main.go

mcs /vagrant/cs/ -out:/vagrant/bin/storage

# TODO: Compile Jade/Pug templates