package main

import (
	"github.com/urfave/cli"
	"github.com/fatih/color"
	"os"
	"solr_console/service/solr"
	"bufio"
	"strings"
	"solr_console/service/connection"
	"solr_console/service/printer"
)

var app *cli.App


func init(){
	app = cli.NewApp()
	app.Name = "solr_console"
	app.Usage = "This is a solr console management tool"
	initFlags()
	initCommands()
}


func main() {
	isConnected := connection.Connection()
	if isConnected{
		args := os.Args
		for true {
			inputArgs := make([]string, 3)
			printer.Console.Printf("%s", FLAG)
			inputReader := bufio.NewReader(os.Stdin)
			str, _ := inputReader.ReadString('\n')
			inputArgs = strings.Split(str, " ")
			if strings.TrimSpace(inputArgs[0]) == "exit" {
				break
			}else {
				for _, v := range inputArgs{
					v = strings.TrimSpace(v)
					args = append(args, v)
				}
				app.Run(args)
				args = []string{}
			}
		}
	}else{
		printer.ErrorPrinter.Println("链接错误！请重新试！")
	}

}

const FLAG = "golang>"

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
				resp := solr.Create(c.Args().First())
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
				resp := solr.Update(c.Args().First())
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
				resp := solr.Delete(c.Args().First())
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
