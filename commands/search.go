package commands

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/os/gfile"
	cli "github.com/jawher/mow.cli"
	"path"
	"regexp"
	"strings"
)


func match(input string,name string,expression string) bool{
	if name != ""{
		if !strings.Contains(input,name){
			return false
		}
	}

	if expression != ""{
		match,_ := regexp.MatchString(expression,input)
		if !match{
			return false
		}
	}

	return true
}


func cmdSearch(cmd *cli.Cmd){
	var (
		limitOpt = cmd.IntOpt("c count",10,"result count")
		nameOpt = cmd.StringOpt("n name","","search by name")
		expressionOpt = cmd.StringOpt("e expression","","search by expression")
	)

	cmd.Action = func() {

		limit := *limitOpt
		name := *nameOpt
		expression := *expressionOpt

		matchCount := 0
		for _, storeInfo := range getAllStore() {
			storeID := storeInfo[0]
			storeName := storeInfo[1]
			if storeName == ""{
				storeName = storeID
			}

			fmt.Println(storeName)

			content := getStoreContent(storeID)

			for _, line := range content {
				if match(line,name,expression){
					fmt.Println("\t" + line)
					matchCount++
					if matchCount == limit{
						break
					}
				}
			}
		}
	}


}


func getStoreContent(storeID string) []string{
	root_dir := rootFolder()

	wd_id := storeID
	if wd_id == ""{
		wd_id = wdId()
	}

	wd_id_dir := path.Join(root_dir,wd_id)
	content_file := path.Join(wd_id_dir,"content.json")

	if !gfile.Exists(content_file){
		return []string{}
	}

	data := gfile.GetBytes(content_file)

	var list []string
	if err := json.Unmarshal(data,&list);err != nil{
		panic(err)
	}
	return list
}