package main

import (
	"fmt"
	"os"
)

func getAllENVSFilesData(filepath string){
	path,err:=os.ReadDir(filepath)
	if err !=nil{
		fmt.Println(err)
	}
	for _,files:= range path{
		if files.IsDir(){
			getAllENVSFilesData(filepath + files.Name() + "/")
			continue
		}
		if files.Name() == ".env"{
			data,err:= os.ReadFile(filepath +files.Name())
			if err!=nil{
				fmt.Println(err)
				}
				fmt.Printf("Data from %s:\n%s\n", files.Name(), string(data))
		fmt.Println("-------------------------") 
		}
	}
}

func main(){
	source:= "/home/lemon/Desktop/"
	getAllENVSFilesData(source)
}