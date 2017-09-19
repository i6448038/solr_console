package connection

import (
	"time"
	"os"
	"io/ioutil"
	"strings"
)

const (
	TOKEN_PATH = "/var/tmp/solortoken"
	EXPIRE = 10 * time.Minute
)

//连接创建后，生成token，可以保持多长时间的连接
func createToken(host,port string)error{
	//查看文件是否存在，不存在就创建
	if checkFileExist(TOKEN_PATH) == false {
		error := createFile(host, port)
		if error != nil{
			return error
		}
	}

	//文件存在，且没过期，就删掉重新
	error := os.Remove(TOKEN_PATH)

	if error != nil {
		return error
	}

	error = createFile(host, port)
	if error != nil{
		return error
	}

	return nil
}

//检查token内容，并连接
func checkToken()(host, port string){
	//查看文件是否存在，不存在就返回false
	if checkFileExist(TOKEN_PATH) == false {
		return "", ""
	}

	//文件存在的情况下，看是否文件过期,过期就返回false
	if checkFileExpired(TOKEN_PATH, EXPIRE) == false{
		return "",""
	}

	//文件即存在，且没过期，就取出内容并连接。
	fileContent, error := ioutil.ReadFile(TOKEN_PATH)
	if error != nil{
		return "",""
	}
	content := strings.Split(string(fileContent), ":")
	return content[0], content[1]

}

//查看文件是否存在
func checkFileExist(path string)bool{
	_, error := os.Stat(path)
	if error != nil{
		return false
	}
	return true
}

//查看文件是否过期
func checkFileExpired(path string, duration time.Duration)bool{
	fileInfo, error := os.Stat(path)
	if error != nil{
		return true
	}
	if fileInfo.ModTime().Add(duration).Unix() > time.Now().Unix(){
		return true
	}
	return false
}

//创建文件并写入内容
func createFile(host, port string) error{
	f, error := os.Create(TOKEN_PATH)
	if error != nil{
		return error
	}
	f.WriteString(host + ":" +port)
	f.Close()
	return nil
}
