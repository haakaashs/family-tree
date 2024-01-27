package main

import (
	"github.com/haakaashs/family-tree/services"
	"github.com/haakaashs/family-tree/utils"
)

func main() {
	utils.InitializeDB("config.env")
	services.Commands()
}
