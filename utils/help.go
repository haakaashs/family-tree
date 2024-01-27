package utils

import "fmt"

type add struct{}
type connect struct{}
type count struct{}
type father struct{}

type Helper struct {
	add     add
	connect connect
	count   count
	father  father
}

func NewHelper() *Helper {
	var (
		a  add
		c  connect
		co count
		f  father
	)
	return &Helper{
		add:     a,
		connect: c,
		count:   co,
		father:  f,
	}
}

func (*add) Help() {
	fmt.Println("add is used with subcommands")
	fmt.Println("add <name>")
	fmt.Println("add relationship <name>")
}

func (*connect) Help() {
	fmt.Println("connect is used with subcommands")
	fmt.Println("connect <name1> as <relation> of <name2>")
}

func (*count) Help() {
	fmt.Println("count is used with subcommands")
	fmt.Println("count sons of <name>")
	fmt.Println("count daughters of <name>")
	fmt.Println("count wifes of <name>")
}

func (*father) Help() {
	fmt.Println(" is used with subcommands")
	fmt.Println("")
}
