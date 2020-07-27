/**
* @Author:zhoutao
* @Date:2020/7/27 下午11:22
 */

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModelUpper                      = iota + 1 //转大写
	ModelLower                                 //转小写
	ModelUnderscoreToUpperCamelCase            //下划线转大写驼峰
	ModelUnderscoreToLowerCamelCase            //下划线转小写驼峰
	ModelCameCaseToUnderscore                  //驼峰转下划线
)

var str string
var mode int8
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModelUpper:
		case ModelLower:
		case ModelUnderscoreToUpperCamelCase:
		case ModelUnderscoreToLowerCamelCase:
		case ModelCameCaseToUnderscore:
		default:
			log.Fatalf("暂时不支持该转换模式，请执行help word查看帮助文档")
		}
		log.Printf("输出结果:%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
