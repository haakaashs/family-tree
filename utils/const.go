package utils

const (
	ADD          = "add"
	RELATIONSHIP = "relationship"
	CONNECT      = "connect"
	COUNT        = "count"
	GET          = "get"
	PERSON       = "person"
	FATHER       = "father"
	AS           = "as"
	OF           = "of"
	SONS         = "sons"
	DAUGHTERS    = "daughters"
	WIVES        = "wives"
)

var RelationMap = map[string]string{"sons": "son", "daughters": "daughter", "wives": "spouse"}
