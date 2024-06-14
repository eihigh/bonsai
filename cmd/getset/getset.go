package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type hoge struct {
	a int    `get:"" set:""`
	b int    `getset:",setID"`
	c string `getset:"C,"`
}
