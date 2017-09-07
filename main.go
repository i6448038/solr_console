package main

import (
	//"os"
	"github.com/urfave/cli"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/fatih/color"
	"solr_console/utils"
	"os"
)

var app *cli.App


func init(){
	app = cli.NewApp()
	app.Name = "solr_console"
	app.Usage = "This is a solr console management tool"
	//initActions()
	initFlags()
	initCommands()
}



func main() {
	args := os.Args
	app.Run(args)
	for true {
		str := ""
		fmt.Printf("%s", flag)
		fmt.Scanf("%s", &str)
		if str == "exit" {
			break
		}else {
			args = append(args, str)
			app.Run(args)
		}
	}

}

const flag = "golang>"

func initActions(){
	app.Action = func(c * cli.Context) error {
		color.Yellow("welcome to the solr management platform!")

		version, err := getSolrVersion()

		if err != nil {
			color.Red("出现了错误！" + err.Error())
			return err
		}

		color.Yellow("the version is:" + version)

		color.Yellow("if you need any help，please type --help")


		return nil
	}
}

//初始化
func initCommands(){
	app.Commands = []cli.Command{
		{
			Name: "add",
			Usage: "add a new document or some documents",
			Aliases: []string{"c"},
			Description: "the c of the `CURD`,insert a new data into the solr core",
			ArgsUsage: `solr的格式为json字符串或者字符串数组（例如: {"id": 22, "title": "abc"}
			或者：[{"id": 22, "title": "abc"}, {"id": 23, "title": "def"}, {"id": 24, "title": "def"}]）` ,
			Action: func (c *cli.Context)error {
				resp := utils.Create(c.Args().First())
				if(resp != nil){
					color.Blue(resp.String())
				}
				return nil
			},
		},
		{
			Name: "update",
			Usage: "update a document or some documents",
			Aliases: []string{"u"},
			Description: "the u of the `CURD`,update documents",
			ArgsUsage: `solr的格式为json字符串或者字符串数组（例如: {"id": 22, "title": "abc"}
			或者：[{"id": 22, "title": "abc"}, {"id": 23, "title": "def"}, {"id": 24, "title": "def"}]）` ,
			Action: func (c *cli.Context)error {
				color.Yellow(c.Args().First())
				resp := utils.Update(c.Args().First())
				if(resp != nil){
					color.Blue(resp.String())
				}
				return nil
			},
		},
		{
			Name: "delete",
			Usage: "delete a document or some documents",
			Aliases: []string{"d"},
			Description: "the d of the `CURD`,delete documents",
			Action: func (c *cli.Context)error {
				color.Red(c.Args().First())
				resp := utils.Delete(c.Args().First())
				if(resp != nil){
					color.Blue(resp.String())
				}
				return nil
			},
		},
		{
			Name: "test",
			Usage: "test",
			Aliases: []string{"t"},
			Description: "the d of the `CURD`,delete documents",
			Action: func (c *cli.Context)error {
				color.Red("success")
				return nil
			},
		},
	}
}


func initFlags(){
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "language for the greeting",
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
	}
}


func getSolrVersion()(version string, err error) {
	host := "localhost"
	port := "8983"
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

	version = "未知"

	return  //解析正确，但是没有想要的json key
}
