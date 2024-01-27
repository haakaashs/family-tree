package services

import (
	"fmt"
	"os"

	"github.com/haakaashs/family-tree/utils"
)

func addPerson(n, g string) {
	person := utils.Person{Name: n, Gender: g}
	utils.Conn.Model(utils.Person{}).Create(&person)
	fmt.Println("added", person.Name)
}

func addRelation(s string) {
	relation := utils.Relation{Name: s}
	utils.Conn.Model(utils.Relation{}).Create(&relation)
	fmt.Println("added relation", relation.Name)
}

func connectPersons(c map[string]string) {

	var count int64
	utils.Conn.Model(&utils.Relationship{}).
		Joins("JOIN persons p1 ON p1.id = relationships.person_id1").
		Joins("JOIN persons p2 ON p2.id = relationships.person_id2").
		Where("(p1.name=? AND p2.name=?)", c[utils.CONNECT], c[utils.OF]).Count(&count)
	if count > 0 {
		fmt.Printf("relationship already exist between %s %s\n", c[utils.CONNECT], c[utils.OF])
		os.Exit(0)
	}

	relation := utils.Relation{}
	utils.Conn.Model(&utils.Relation{}).Where("name=?", c[utils.AS]).Find(&relation)
	if relation.Id == 0 {
		fmt.Printf("add relation %s to use\n", c[utils.AS])
		os.Exit(0)
	}

	persons := []utils.Person{}
	utils.Conn.Model(&utils.Person{}).Where("name in ?", []string{c[utils.CONNECT], c[utils.OF]}).Find(&persons)
	if len(persons) < 2 {
		fmt.Println("given person does not exist")
		os.Exit(0)
	}

	relationship := utils.Relationship{RelationId: relation.Id}
	for _, person := range persons {
		if c[utils.CONNECT] == person.Name {
			relationship.PersonId1 = person.Id
		} else if c[utils.OF] == person.Name {
			relationship.PersonId2 = person.Id
		}
	}

	utils.Conn.Model(utils.Relationship{}).Create(&relationship)
	fmt.Printf("added relationship %s %s of %s\n", c[utils.CONNECT], c[utils.AS], c[utils.OF])
}

func countPersons(c map[string]string) {

	val := c[utils.COUNT]
	var count int64
	relation := utils.RelationMap[c[utils.COUNT]]
	utils.Conn.Model(&utils.Relationship{}).
		Joins("JOIN relations ON relations.id = relationships.relation_id").
		Where("relationships.person_id2=(select id from persons where name=?) and relations.name=?", c[utils.OF], relation).Count(&count)

	if val == utils.SONS {
		fmt.Printf("%s has %d %s\n", c[utils.OF], count, utils.SONS)
	} else if val == utils.DAUGHTERS {
		fmt.Printf("%s has %d %s\n", c[utils.OF], count, utils.DAUGHTERS)
	} else if val == utils.WIVES {
		fmt.Printf("%s has %d %s\n", c[utils.OF], count, utils.WIVES)
	} else {
		help()
		os.Exit(0)
	}
}

func getPersonName(s string) {
	var father string
	utils.Conn.Model(&utils.Relationship{}).
		Select("persons.name").
		Joins("JOIN persons ON persons.id = relationships.person_id2 ").
		Where("relationships.person_id1 = (SELECT id FROM persons WHERE name = ?)", s).
		First(&father)
	fmt.Printf("%s is the father of %s\n", father, s)
}

func getRelation() {
	var person []utils.Person
	utils.Conn.Model(&utils.Person{}).Find(&person)
	fmt.Println("Name        Gender")
	for _, per := range person {
		fmt.Printf("%s   ------->   %s\n", per.Name, per.Gender)
	}
}
