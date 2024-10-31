package main

import (
	"fmt"
	"os"
	"time"
)

func ChangeTime(){
	currentTime:= time.Now()
	endTime := currentTime.Add(time.Hour * 1)
	err:=os.Chtimes("chtimes.go",currentTime,endTime)
	if err !=nil{
		fmt.Println(err)
	}
}

func main(){
	ChangeTime()
}