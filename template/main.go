/**
* @Author:zhoutao
* @Date:2020/7/28 上午9:28
 */

package main

import (
	"os"
	"strings"
	"text/template"
)

const templateText = `
OutPut 0:{{title .Name1}}
OutPut 1:{{tile .Name2}}
OutPut 2:{{.Name3 | title}}
`

func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl, _ := template.New("go-programming").Funcs(funcMap).Parse(templateText)

	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)
}
