/**
* @Author:zhoutao
* @Date:2020/7/28 上午8:12
 */

package timer

import "time"

func GetNow() time.Time {
	location, _ := time.LoadLocation("Asia/shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(duration), nil
}
