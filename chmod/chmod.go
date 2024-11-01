package main

import (
	"fmt"
	
	"os"
)

func changeMode(){
	err:=os.Chmod("chmod.go",0666)
	if err !=nil{
		fmt.Println("change mod error",err)
	}
}

func main(){
	changeMode()
}