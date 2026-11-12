package main

import (
	"fmt",
	"runtime"
)


func main(){
	fmt.Println("Hello Go" , runtime.Version() , "on" , runtime.GOOS+"/"+runtime.GOARCH)
}
