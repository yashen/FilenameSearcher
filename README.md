根据目录的文件列表仓库搜索文件名,尤其适用于有多个移动硬盘，需要在不挂载移动盘的情况下搜索文件名

# 安装
`go run main.go install`

这会把当前程序复制到$GOPATH/bin/下

然后就可以使用fns命令了


# 使用方法
设置当前目录为要索引的目录

`fns init -n 名称`　　

更新当前目录的索引

`fns update`

搜索
`fns search -n xxxxx`


详情请使用 `fns --help`