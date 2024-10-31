package main

import "os"

func ClearENV(){
	os.Clearenv()
}

func main(){
	ClearENV()
}