package tool

import (
	"github.com/tedcy/fdfs_client"
	"fmt"
	"os"
	"bufio"
	"strings"
)

//文件上传到fdfs
func UploadFile(fileName string) string {

	client, err := fdfs_client.NewClientWithConfig("../config/fastdfs.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return fileId
}

//从配置文件读取文件服务器的ip和端口相关配置
func FileServerAddr() string {

	file, err := os.Open("../config/fastdfs.conf")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return ""
		}

		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)

		if str[0] == "http_server_port" {
			return str[1]
		}
	}
}