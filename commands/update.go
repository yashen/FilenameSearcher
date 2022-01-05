package commands

import (
	"encoding/json"
	"github.com/gogf/gf/os/gfile"
	cli "github.com/jawher/mow.cli"
	"path"
)


func getContent() []string{
	root_dir := rootFolder()

	wd_id := wdId()

	wd_id_dir := path.Join(root_dir,wd_id)
	content_file := path.Join(wd_id_dir,"content.json")

	data := gfile.GetBytes(content_file)

	var list []string
	if err := json.Unmarshal(data,&list);err != nil{
		panic(err)
	}
	return list
}


func cmdUpdate(cmd *cli.Cmd){
	cmd.Action = func(){
		root_dir := rootFolder()
		if !gfile.Exists(root_dir){
			gfile.Mkdir(root_dir)
		}

		wd_id := wdId()

		wd_id_dir := path.Join(root_dir,wd_id)
		if !gfile.Exists(wd_id_dir){
			gfile.Mkdir(wd_id_dir)
		}


		content_file := path.Join(wd_id_dir,"content.json")


		files := walk()
		data,_ := json.Marshal(files)

		gfile.PutBytes(content_file,data)
		saveStoreUpdatePath(wdId(),wd())
	}
}