package main

import (
	"log"
	"os"
	
)


func Getwd(){
	//get working dir
	path,err:=os.Getwd()
	if err !=nil{
		log.Fatal("err in geting a wk",err)
	}
	//changeDir
	os.Chdir(path.Split("/"))
}