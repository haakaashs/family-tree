package services

import (
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
// family-tree get person
var helper = utils.NewHelper()

func Commands() {

	if os.Args[1] == utils.HP {
		helper.Defaul.Help()
		os.Exit(0)
	}
	switch os.Args[1] {
	case utils.ADD:
		if len(os.Args) != 4 {
			helper.Add.Help()
			os.Exit(0)
		}
		if os.Args[2] == utils.RELATIONSHIP {
			addRelation(os.Args[3])
			return
		}
		addPerson(os.Args[2], os.Args[3])
	case utils.CONNECT:
		if os.Args[3] != utils.AS || os.Args[5] != utils.OF && len(os.Args) != 7 {
			helper.Connect.Help()
			os.Exit(0)
		}
		connectPersons(getMap(getList()))
	case utils.COUNT:
		if os.Args[3] == utils.OF && len(os.Args) != 5 {
			helper.Count.Help()
			os.Exit(0)
		}
		countPersons(getMap(getList()))
	case utils.FATHER:
		if os.Args[2] != utils.OF {
			helper.Father.Help()
			os.Exit(0)
		}
		getPersonName(os.Args[3])
	case utils.GET:
		if os.Args[2] != utils.PERSON || len(os.Args) != 3 {
			helper.Get.Help()
			os.Exit(0)
			return
		}
		getRelation()
	default:
		helper.Defaul.Help()
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
