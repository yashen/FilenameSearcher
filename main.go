package main

import (
	"FilenameSearcher/commands"
	"fmt"
)

func main(){

	defer func() {
		if err := recover();err != nil{
			fmt.Println(fmt.Sprintf("操作中出现错误,%v",err))
		}
	}()
	commands.Main()
}