package solr

import (
	"github.com/rtt/Go-Solr"
	"github.com/fatih/color"
	"fmt"
	"solr_console/utils"
)

const (
	CREATE = iota
	UPDATE
	DELETE
)

var client *solr.Connection
var err error

func init(){
	client, err = solr.Init("localhost", 8983, "jiyu_test")

	if err != nil{
		color.Red("solr 链接有问题！请查对IP地址、端口号和solr collection")
		return
	}
}

func writeSolrDate(json string, flag int) *solr.UpdateResponse{
	document, err := utils.Decode(json)

	if err != nil{
		color.Red("请输入合法的json字符串!", err.Error())
		return nil
	}

	var resp *solr.UpdateResponse

	switch flag {
	case CREATE:
		var list []interface{}
		list = append(list, document)
		fmt.Println(list)
		doc := map[string]interface{}{
			"add":list,
		}
		resp, err = client.Update(doc, true)

		if err != nil {
			color.Red("Solr数据添加有误！")
			color.Red(err.Error())
			return nil
		}
		break
	case UPDATE:
		var list []interface{}
		list = append(list, document)
		doc := map[string]interface{}{
			"add":list,
		}
		resp, err = client.Update(doc, true)

		if err != nil {
			color.Red("Solr数据更新有误！\n")
			color.Red(err.Error())
			return nil
		}
		break
	case DELETE:
		var list []interface{}
		list = append(list, document)
		doc := map[string]interface{}{
			"delete":list,
		}

		resp, err = client.Update(doc, true)

		if err != nil {
			color.Red("删除有误！")
			color.Red(err.Error())
			return nil
		}
	}

	return resp
}




func Create(json string) *solr.UpdateResponse{
	fmt.Println(json)
	return writeSolrDate(json, CREATE)
}

func Delete(json string) *solr.UpdateResponse {
	return writeSolrDate(json, DELETE)
}

func Update(json string) *solr.UpdateResponse{
	return writeSolrDate(json, UPDATE)
}
