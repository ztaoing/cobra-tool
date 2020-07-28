/**
* @Author:zhoutao
* @Date:2020/7/27 下午11:18
 */

package main

import (
	"cobra-tool/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
}
