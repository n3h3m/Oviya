#! /opt/local/bin/fish
go run main.go examples/hello.ov
gcc examples/hello.c -o bin/hello
./bin/hello