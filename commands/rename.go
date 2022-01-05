package commands

import cli "github.com/jawher/mow.cli"

func cmdRename(cmd *cli.Cmd){
	newNameArg := cmd.StringArg("NAME","","new store name")

	cmd.Action = func() {
		storeID := wdId()
		newName := *newNameArg
		saveStoreName(storeID,newName)
	}
}
