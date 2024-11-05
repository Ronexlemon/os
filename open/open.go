package main

import (
	"log"
	"os"
)


func OpenFile(){
	//change directory
	err:=os.Chdir("newfile")
	if err !=nil{
		log.Fatal("failed to change directory",err)
	}
	//open file
	file,err:=os.Open("Global.go")
	if err !=nil{
		log.Fatal("failed to open file",err)
		}
	//read file content
	data:=make([]byte,100)
	d,err:=file.Read(data)
	if err !=nil{
		log.Fatal("failed to read file",err)
		}
		log.Println("read file content:",string(data[:d]))
}
func main(){
	OpenFile()
}