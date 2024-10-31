package main

import (
	"fmt"
	"io/fs"
	"os"
)

func changeMode(){
	err:=os.Chmod("chmod.go",fs.ModeAppend)
	if err !=nil{
		fmt.Println("change mod error",err)
	}
}

func main(){
	changeMode()
}