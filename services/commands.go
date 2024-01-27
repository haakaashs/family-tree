package services

import (
	"fmt"
	"os"

	"github.com/haakaashs/family-tree/utils"
)

// tool get only first name of the person
// family-tree add haak male
// family-tree add relationship brother
// family-tree connect jany as brother of haak
// family-tree count son of haak
// family-tree count daughter of haak
// family-tree count wifes of haak
// family-tree father of haak
func Commands() {
	switch os.Args[1] {
	case utils.ADD:
		if os.Args[2] == utils.RELATIONSHIP && len(os.Args) == 4 {
			addRelation(os.Args[3])
			return
		} else if len(os.Args) == 4 {
			addPerson(os.Args[2], os.Args[3])
			return
		}
		fmt.Println(os.Args, len(os.Args))
		help()
		os.Exit(0)
	case utils.CONNECT:
		if os.Args[3] != utils.AS || os.Args[5] != utils.OF && len(os.Args) != 7 {
			help()
			os.Exit(0)
		}
		connectPersons(getMap(getList()))
	case utils.COUNT:
		if os.Args[3] == utils.OF && len(os.Args) != 5 {
			help()
			os.Exit(0)
		}
		countPersons(getMap(getList()))
	case utils.FATHER:
		if os.Args[2] != utils.OF {
			help()
			os.Exit(0)
		}
		getPersonName(os.Args[3])

	default:
		help()
	}
}

func getList() []string {
	var cmds []string
	for i, c := range os.Args {
		if i == 0 {
			continue
		}
		cmds = append(cmds, c)
	}
	return cmds
}

func getMap(c []string) map[string]string {
	mp := make(map[string]string)
	i, j := 0, 1
	for j <= len(c) {
		mp[c[i]] = c[j]
		i, j = j+1, j+2
	}
	return mp
}

func help() {
	fmt.Println("helper")
}
