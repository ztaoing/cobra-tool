/**
* @Author:zhoutao
* @Date:2020/7/28 上午8:18
 */

package cmd

import (
	"cobra-tool/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNow()
		log.Fatalf("输出结果：%s,%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var caculateTime string
var duration string

var caculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if caculateTime == "" {
			currentTimer = timer.GetNow()
		} else {
			var err error
			space := strings.Count(caculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}

			currentTimer, err = time.Parse(layout, caculateTime)
			if err != nil {
				t, _ := strconv.Atoi(caculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalln("timer.GetCalculateTime err:%v", err)
		}
		log.Printf("输出结果:%s,%d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(caculateTimeCmd)

	caculateTimeCmd.Flags().StringVarP(&caculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已格式化的时间`)
	caculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns","us","ms","s","m","h"`)
}
