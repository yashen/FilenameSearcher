package commands

import (
	"encoding/json"
	"github.com/gogf/gf/os/gfile"
)


type DataMeta struct {
	Name string `json:"name"`
	UpdatePath string `json:"update_path"`
}

func saveStoreName(storeID string,name string){
	meta, dataMetaFileName := getMeta(storeID)

	if name == "" {
		panic("名称不能为空")
	}
	meta.Name = name

	metaData, _ := json.Marshal(meta)

	gfile.PutBytes(dataMetaFileName, metaData)
}

func saveStoreUpdatePath(storeID string,path string){
	meta, dataMetaFileName := getMeta(storeID)

	meta.UpdatePath = path

	metaData, _ := json.Marshal(meta)

	gfile.PutBytes(dataMetaFileName, metaData)
}

func getMeta(storeID string) (DataMeta, string) {
	var meta DataMeta
	dataFolder(storeID, true)

	dataMetaFileName := dataMetaFile(storeID)
	if gfile.Exists(dataMetaFileName) {
		if err := json.Unmarshal(gfile.GetBytes(dataMetaFileName), &meta); err != nil {
			panic(err)
		}
	} else {
		meta = DataMeta{}
	}
	return meta, dataMetaFileName
}
