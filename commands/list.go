package commands

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"strings"
)

func cmdListStore(cmd *cli.Cmd){
	var (
		nameOpt = cmd.StringOpt("n name","","Search by Name")
		longOpt = cmd.BoolOpt("l long",false,"Show more info")
	)

	cmd.Action = func() {
		name := *nameOpt
		for _, parts := range getAllStore() {
			storeID := parts[0]
			storeName := parts[1]

			if name == "" || strings.Contains(storeName,name){
				if *longOpt{
					fmt.Println(storeID,storeName,parts[2])
				}else{
					fmt.Println(storeID,storeName)
				}
			}

		}
	}

}
