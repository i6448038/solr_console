package connection

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"solr_console/service/printer"
)

func getSolrVersion(host, port string)(version string, err error) {
	resp, err := http.Get("http://"+host +":"+ port + "/solr/admin/info/system?wt=json")
	if err != nil{
		return "", err
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
	var version string
	var err error

	host, port := checkToken()

	if host != "" && port != ""{
		version, err = getSolrVersion(host, port)
		if err != nil {
			return false
		}
	}else {
		printer.InfoPrinter.Println("please press host and port!")
		printer.InfoPrinter.Print("host:")
		fmt.Scanf("%s", &host)
		printer.InfoPrinter.Print("port:")
		fmt.Scanf("%s", &port)
		version, err = getSolrVersion(host, port)
		if err != nil {
			return false
		}

		err = createToken(host, port)
		if err != nil{
			return false
		}
	}


	printer.InfoPrinter.Println("connect!")
	printer.InfoPrinter.Println("welcome! the version is:" + version)
	printer.InfoPrinter.Println("\n\n\n")
	return true
}




