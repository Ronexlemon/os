package main

import (
	"log"
	"os"
)

func Getwd() {
	//get working dir
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("err in geting a wk", err)
	}
	log.Println(path)
	//changeDir
	err=os.Chdir("newfile")
	if err !=nil{
		log.Fatal("Failed to change dir",err)
	}
	file,err:=os.Create("Global.go")
	if err !=nil{
		log.Fatal("failed to create a file",err)
	}
	defer file.Close()
	_,err=file.WriteString("package main \n  import (`fmt`) \n func Global(){\n num:=300 \n fmt.Println(num) }")
	_,err=file.WriteString("package main \n  import (`fmt`) \n func Global(){\n num:=300 \n fmt.Println(num) }")
	
	if err!=nil{
		log.Fatal("failed to write to file",err)
		}
}

func main(){
	Getwd()
}
