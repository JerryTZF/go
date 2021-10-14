/**
 * Created by GoLand
 * Time: 2021/10/14 11:24 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: time_.go
 * Desc:
 */
package library

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"time"
)

var TimeFormat = "2006-01-02 15:04:05"

// 常规时间操作
func TimeDemo() {
	// 当前时间
	now := time.Now()
	// 年
	year := now.Year()
	// 月
	month := now.Month()
	// 日
	day := now.Day()
	// 时
	hour := now.Hour()
	// 分
	minute := now.Minute()
	// 秒
	second := now.Second()
	// 周
	week := now.Weekday()

	fmt.Println(now)
	fmt.Println(int(week))
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 时间戳(秒级)
	timestamp1 := now.Unix()
	// 时间戳(纳秒级)
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp1, timestamp2)

	// 时间戳转换时间对象
	timeObj := time.Unix(timestamp1, 0)
	sj := timeObj.Format(TimeFormat)
	fmt.Println(sj)

	// 时间加上指定时间长度
	later := now.Add(24 * 4 * time.Hour)
	fmt.Println(later.Format(TimeFormat))
	// 两个时间差值
	before := now.Sub(later)
	fmt.Println(before.Hours())

	// 判断一个时间点是否在另一个时间点之前
	isBefore := now.Before(later)
	fmt.Println(isBefore)
	// 判断一个时间点是否在另一个时间点之后
	isAfter := later.After(now)
	fmt.Println(isAfter)
}

// 时间字符串相关解析
func ParseStringTime() {
	timeStr := "2021-10-01 12:23:01"
	timeObj, _ := time.ParseInLocation(TimeFormat, timeStr, time.Local)
	fmt.Println(timeObj)
}

func GoCarbon() {
	// 今天日期
	fmt.Println(carbon.Now().ToDateString()) // 2021-10-14
	// 今天时间
	fmt.Println(carbon.Now().ToTimeString()) // 14:46:01
	// 今天日期+时间
	fmt.Println(carbon.Now().ToDateTimeString()) // 2021-10-14 14:46:01
	// 指定时区的今天此刻(伦敦)
	fmt.Println(carbon.Now(carbon.London).ToDateTimeString())               // 2021-10-14 07:47:28
	fmt.Println(carbon.SetTimezone(carbon.London).Now().ToDateTimeString()) // 2021-10-14 07:47:28
	// 时间戳
	fmt.Println(carbon.Now().Timestamp())                // 秒级时间戳
	fmt.Println(carbon.Now().TimestampWithSecond())      // 秒级时间戳
	fmt.Println(carbon.Now().TimestampWithMillisecond()) // 毫秒时间戳
	fmt.Println(carbon.Now().TimestampWithMicrosecond()) // 微秒时间戳
	fmt.Println(carbon.Now().TimestampWithNanosecond())  // 纳秒时间戳

	// 具体详情参见: https://github.com/golang-module/carbon/blob/master/README.cn.md
}
