package commands

import (
	"encoding/json"
	"github.com/gogf/gf/os/gfile"
	"os"
	"path"
)

var searchTargetConfigFile string = ".fns"

func wd() string{
	dir,err := os.Getwd()
	if err != nil{
		panic(err)
	}
	return dir
}

func wdId() string{
	configFile := path.Join(wd(), searchTargetConfigFile)
	return gfile.GetContents(configFile)
}

func rootFolder() string{
	userDir,_ := os.UserHomeDir()
	rootDir := path.Join(userDir,".fnsd")
	return rootDir
}

func dataFolder(storeID string,createFolder bool) string{
	rootDir := rootFolder()
	if createFolder && !gfile.Exists(rootDir){
		gfile.Mkdir(rootDir)
	}

	wdId := storeID
	result := path.Join(rootDir, wdId)
	if createFolder && !gfile.Exists(result){
		gfile.Mkdir(result)
	}

	return result
}

var dataMetaFileName = "meta.json"
var dataContentFileName = "content.json"

func dataMetaFile(storeID string) string{
	return path.Join(dataFolder(storeID,false), dataMetaFileName)
}

func dataContentFile(storeID string) string{
	return path.Join(dataFolder(storeID,false), dataContentFileName)
}

func getAllStore() [][3]string{
	rootFolder := rootFolder()

	result := make([][3]string,0)

	childNames,_ := readDirNames(rootFolder)

	for _, storeID := range childNames {
		childPath := path.Join(rootFolder, storeID)

		name, updatePath := getStoreName(childPath)
		result = append(result,[3]string{storeID,name, updatePath})
	}

	return result
}

func getStoreName(storePath string) (string,string){
	metaFile := path.Join(storePath,dataMetaFileName)
	if !gfile.Exists(metaFile){
		return "",""
	}

	jsonData := gfile.GetBytes(metaFile)
	var dataMeta DataMeta
	json.Unmarshal(jsonData,&dataMeta)
	return dataMeta.Name,dataMeta.UpdatePath
}


func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	return names, nil
}