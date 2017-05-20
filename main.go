package main

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func main() {
	app := cli.NewApp()
	app.Name = "solr_console"
	app.Usage = "This is a solr console management tool"
	app.Action = func(c * cli.Context) error {
		fmt.Println("welcome to the solr management platform")

		version, err := getSolrVersion("lucene")

		if err != nil {
			fmt.Println("出现了错误！", err.Error())
			return err
		}
		fmt.Println("Solr的版本为:", version)

		return nil
	}
	app.Run(os.Args)
}

func getSolrVersion(key string)(version interface{}, err error) {
	host := "localhost"
	port := "8983"
	resp, err := http.Get("http://"+host +":"+ port + "/solr/admin/info/system?wt=json")
	if err != nil{
		fmt.Println("Please press the right address！")
	}
	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil{
		return nil, errBody
	}
	var msg map[string]interface{}
	err = json.Unmarshal(body, &msg)

	if lucene, ok := msg[key].(map[string]interface{}) ; ok{
		version = lucene["solr-spec-version"]
		return
	}

	defer resp.Body.Close()


	return  //解析正确，但是没有想要的json key
}
