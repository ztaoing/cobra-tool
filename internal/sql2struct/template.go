/**
* @Author:zhoutao
* @Date:2020/7/28 上午10:29
 */

package sql2struct

import (
	"cobra-tool/internal/word"
	"fmt"
	"os"
	"text/template"
)

//模板对象声明，把得到的列信息按照特定的规则转为go结构体,这里采用模板渲染的方案
const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}} {{$length:=len.Comment}} {{if gt $length 0}}//
{{.Comment}} {{else}}// {{.Name}}{{end}}
{{$typeLen :=len .Type}} {{if gt $typeLen 0}} {{.Name | ToCamelCase}}
{{end}}}

func (model {{.TableName | ToCamelCase}})TableName()string{
return "{{.TableName}}"
}`

//StructTemplate和StructColumn 实际上对应的是不同阶段的模板对象信息
//对后续模板渲染对象进行声明
type StructTemplate struct {
	structTpl string
}

//用来存储转换后的go结构体中的所有字段信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Commnet string
}

//用来存储最终用于渲染的模板对象信息
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{
		structTpl: structTpl,
	}
}

//通过查询COLUMNS表所组装得到的tbColumn进行进一步的分解和转换
//例如：数据库类型到GO结构体的转换和堆json tag的处理，都在这一层完成
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.ColumnType],
			Tag:     fmt.Sprintf("`json:"+"`", column.ColumnName),
			Commnet: column.ColumnComment,
		})
	}
	return tplColumns
}

//对模板渲染的自定义函数和模板对象进行处理
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase}).Parse(t.structTpl))
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return nil
	}
	return nil
}
