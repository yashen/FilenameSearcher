package commands

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)


var ignoreFolderNames = []string{"node_modules","vender","pkg","bin"}

var ignoreFolderPrefixs = []string{"_",".","$"}

func isIgnoredFolder(folderName string) bool {

	for _, name := range ignoreFolderNames {
		if name == folderName{
			return true
		}
	}

	for _, prefix := range ignoreFolderPrefixs {
		if strings.HasPrefix(folderName,prefix){
			return true
		}
	}

	return false
}

//如果目录中存在这些文件则忽略目录
var ignoreFolderFiles = []string{
	searchTargetConfigFile,
	".encfs6.xml",

}

func walk() []string{
	wd := wd()

	list := make([]string,0)

	top := true
	if walkErr := filepath.Walk(wd,func(filePath string, info fs.FileInfo, err error) error{
		relPath,_ := filepath.Rel(wd,filePath)
		isSymlink := info.Mode() & fs.ModeSymlink != 0
		if info.IsDir(){
			if top{
				top = false
				return nil
			}
			if isSymlink{
				return fs.SkipDir
			}
			if err != nil || isIgnoredFolder(info.Name()){
				return fs.SkipDir
			}

			for _, fileName := range ignoreFolderFiles {
				if gfile.Exists(path.Join(filePath, fileName)){
					return fs.SkipDir
				}
			}

			fmt.Print("\r" + filePath)

			return nil
		}

		if !info.Mode().IsRegular(){
			return nil
		}

		if isSymlink{
			return nil
		}

		if err != nil{
			return err
		}


		if strings.HasPrefix(info.Name(),"$") || strings.HasPrefix(info.Name(),"."){
			return nil
		}


		list = append(list,relPath)

		return nil
	});walkErr != nil{
		panic(walkErr)
	}

	fmt.Println("\n更新完成")
	return list
}
