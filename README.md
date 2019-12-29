# GOESP
[![GoDoc](https://godoc.org/github.com/ardaguclu/goesp?status.png)](http://godoc.org/github.com/ardaguclu/goesp)
[![Go Report Card](https://goreportcard.com/badge/github.com/ardaguclu/goesp)](https://goreportcard.com/report/github.com/ardaguclu/goesp)

Goesp is a command line tool to translate escape analysis executed in build phase to human readable format. 
It renders the result of escape analysis(-gcflags=-m) as table view showing which variables are stored into heap or stack.

In order to make enhancements in terms of performance, Goesp sheds light on GC pressure on packages.

### Install

`go get github.com/ardaguclu/goesp`

### Usage

`goesp {path_of_package.go}`



