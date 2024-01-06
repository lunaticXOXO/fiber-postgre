package main

import (

	"github.com/fiber-postgre/model"
	"github.com/fiber-postgre/route"
)

func main(){

	model.Connect()
	route.Start()

}