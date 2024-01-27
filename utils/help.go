package utils

import "fmt"

type add struct{}
type connect struct{}
type count struct{}
type father struct{}
type get struct{}
type defaul struct{}

type Helper struct {
	Add     *add
	Connect *connect
	Count   *count
	Father  *father
	Get     *get
	Defaul  *defaul
}

func NewHelper() *Helper {
	return &Helper{
		Add:     &add{},
		Connect: &connect{},
		Count:   &count{},
		Father:  &father{},
		Get:     &get{},
		Defaul:  &defaul{},
	}
}

func (*add) Help() {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        add is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree add <name>")
	fmt.Println("           usage: family-tree add relationship <name>")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
}

func (*connect) Help() {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        connect is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree connect <name1> as <relation> of <name2>")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
}

func (*count) Help() {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        count is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree count sons of <name>")
	fmt.Println("           usage: family-tree count daughters of <name>")
	fmt.Println("           usage: family-tree count wifes of <name>")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
}

func (*father) Help() {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        father is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree father of <name>")
	fmt.Println()
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
}

func (*get) Help() {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        get is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree get person")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
}

func (*defaul) Help() {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        add is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree add <name>")
	fmt.Println("           usage: family-tree add relationship <name>")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        connect is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree connect <name1> as <relation> of <name2>")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        count is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree count sons of <name>")
	fmt.Println("           usage: family-tree count daughters of <name>")
	fmt.Println("           usage: family-tree count wifes of <name>")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        father is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree father of <name>")
	fmt.Println()
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println()
	fmt.Println("        get is used with subcommands")
	fmt.Println()
	fmt.Println("           usage: family-tree get person")
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------")
}
