# GOESP
Goesp is a command line tool to convert gcflags output to human readable format. It renders as table view which variables are stored into heap or stack according to the result of escape analysis(-gcflags=-m).

In order to make enhancements in terms of performance, Goesp sheds light on GC pressure of packages.

## Usage
After installing Goesp running command below;

`go get github.com/ardaguclu/goesp`

Executing;

`goesp {path_of_package.go}`


