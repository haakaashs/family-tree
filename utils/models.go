package utils

type Person struct {
	Id     uint32
	Name   string
	Gender string
}

type Relation struct {
	Id   uint32
	Name string
}

type Relationship struct {
	Id         uint32
	PersonId1  uint32
	PersonId2  uint32
	RelationId uint32
}

func (*Person) TableName() string {
	return "persons"
}

func (*Relation) TableName() string {
	return "relations"
}

func (*Relationship) TableName() string {
	return "relationships"
}
