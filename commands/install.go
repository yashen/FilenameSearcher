package commands

import (
	"github.com/gogf/gf/os/gfile"
	cli "github.com/jawher/mow.cli"
	"os"
	"path"
	"runtime"
)

func cmdInstall(cmd *cli.Cmd){
	cmd.Hidden = true
	cmd.Action = func() {
		selfPath := gfile.SelfPath()
		targetPath := path.Join(os.Getenv("GOPATH"),"bin","fns")
		if runtime.GOOS == "windows"{
			targetPath = targetPath + ".exe"
		}
		gfile.Copy(selfPath,targetPath)
	}


}
