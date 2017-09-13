package connection

import (
	"github.com/fatih/color"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os/exec"
)

func getSolrVersion(host, port string)(version string, err error) {
	resp, err := http.Get("http://"+host +":"+ port + "/solr/admin/info/system?wt=json")
	if err != nil{
		fmt.Println("Please press the right address！")
	}
	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil{
		return "", errBody
	}
	var msg map[string]interface{}
	err = json.Unmarshal(body, &msg)

	if lucene, ok := msg["lucene"].(map[string]interface{}) ; ok{
		if version, ok = lucene["solr-spec-version"].(string); ok{
			return
		}
	}

	defer resp.Body.Close()

	version = ""

	return  //解析正确，但是没有想要的json key
}


//连接solr
func Connection() bool{
	var host string
	var port string
	hint := color.New(color.FgBlue)
	color.Yellow("please press host and port! \n")
	hint.Print("host:")
	fmt.Scanf("%s", &host)
	hint.Print("port:")
	fmt.Scanf("%s", &port)

	version, err := getSolrVersion(host, port)
	if err != nil {
		color.Red("出现了错误！" + err.Error())
		return false
	}
	cmd := exec.Command("clear")
	cmd.Run()
	color.Blue("connect!")
	color.Blue("welcome! the version is:" + version)
	color.Blue("")
	color.Blue("")
	color.Blue("")
	return true
}
