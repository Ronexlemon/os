package main

import (
	"fmt"
	"os"
)

func ChangeDir(){
	err:=os.Chdir("chmod")
	if err !=nil{
		println("Error for changing directory",err)
	}
	file,err:=os.Create("chmod.go")
	if err!=nil{
		println("Error for creating file",err)
		}
		defer file.Close()
		_,err= file.WriteString("package main")
		fmt.Println("Hello,World!")
}


func main(){
	ChangeDir()
}