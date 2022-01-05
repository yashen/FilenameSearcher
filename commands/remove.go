package commands

import (
	"github.com/gogf/gf/os/gfile"
	cli "github.com/jawher/mow.cli"
)

func cmdRemove(cmd *cli.Cmd){

	var (
		storeIDOpt = cmd.StringArg("STORE","","store id or name")
	)

	cmd.Action = func() {

		storeID := *storeIDOpt

		retry:
		dataFolder := dataFolder(storeID,false)

		if !gfile.Exists(dataFolder){
			storeName := storeID
			for _, strings := range getAllStore() {
				if strings[1] == storeName{
					storeID = strings[0]
					goto retry
				}
			}
			
			panic("目标仓库不存在的")
		}

		if err := gfile.Remove(dataFolder);err != nil{
			panic(err)
		}
	}


}
