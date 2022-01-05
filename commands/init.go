package commands

import (
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/guid"
	cli "github.com/jawher/mow.cli"
	"path"
)



func cmdInitStore(cmd *cli.Cmd){
	var nameOpt = cmd.StringOpt("n :nameOpt", "", "The store's nameOpt")

	cmd.Action = func() {

		name := *nameOpt

		if name == ""{
			name = wd()
		}
		wd := wd()

		storeID := guid.S()
		storeIDFile := path.Join(wd, searchTargetConfigFile)
		if gfile.Exists(storeIDFile){
			storeID = gfile.GetContents(storeIDFile)
		}

		for _, strings := range getAllStore() {
			if strings[0] == storeID{
				panic("Store has registered")
			} else if strings[1] == name {
				panic("Store's name has used")
			}
		}

		gfile.PutContents(storeIDFile,storeID)
		saveStoreName(storeID,name)
	}

}

