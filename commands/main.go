package commands

import (
	cli "github.com/jawher/mow.cli"
	"os"
)

func Main(){
	app := cli.App("fns", "Filename search in store's index")
	app.Command("list","list stores", cmdListStore)
	app.Command("init","init store", cmdInitStore)
	app.Command("update","update current store's index", cmdUpdate)
	app.Command("search","search files in all stores", cmdSearch)
	app.Command("remove","remove store", cmdRemove)
	app.Command("install","install to path", cmdInstall)
	app.Command("rename","rename store",cmdRename)
	if err := app.Run(os.Args);err != nil{
		panic(err)
	}
}
