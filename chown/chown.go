package main

import "os"

func ChangeOwner(){
	err:=os.Chown("chown.go",1000,1000)
	if err!=nil{
		panic(err)
		}

}

func main(){
	ChangeOwner()
}