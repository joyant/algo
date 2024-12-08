package formula

import "time"

/**
1185. 一周中的第几天
level: easy

给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。

输入为三个整数：day、month 和 year，分别表示日、月、年。

您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}。

示例 1：
输入：day = 31, month = 8, year = 2019
输出："Saturday"

示例 2：
输入：day = 18, month = 7, year = 1999
输出："Sunday"

示例 3：
输入：day = 15, month = 8, year = 1993
输出："Sunday"
*/

/**
如果不使用 time 包，我们可以使用 Zeller's Congruence（Zeller公式）来解决这个问题。Zeller公式是一个计算任意日期是星期几的算法。

Zeller公式如下：

$$ h = (q + \frac{{13(m+1)}}{5} + K + \frac{K}{4} + \frac{J}{4} + 5J) \mod 7 $$

其中：
- \( h \) 是一周中的天数（0 = Saturday, 1 = Sunday, 2 = Monday, ..., 6 = Friday）
- \( q \) 是日
- \( m \) 是月（3 = March, 4 = April, ..., 14 = February）
- \( K \) 是年份的后两位数
- \( J \) 是年份的前两位数

注意：
- 1月和2月被看作是前一年的13月和14月
- 年份的前两位数和后两位数是以公元后为基准的，所以公元前的年份需要特殊处理

以下是使用 Zeller公式的 Go 语言实现：

```go
package main

import (
	"fmt"
)

func dayOfTheWeek(day int, month int, year int) string {
	if month < 3 {
		month += 12
		year--
	}

	K := year % 100
	J := year / 100
	h := (day + 13*(month+1)/5 + K + K/4 + J/4 + 5*J) % 7

	switch h {
	case 0:
		return "Saturday"
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	}

	return ""
}

func main() {
	fmt.Println(dayOfTheWeek(8, 12, 2024)) // 输出 "Sunday"
}
```

在这个代码中，我们首先检查月份，如果是1月或2月，我们将其视为前一年的13月或14月。然后，我们使用 Zeller公式来计算一周中的天数，最后我们根据这个天数返回对应的星期几。
*/

func dayOfTheWeek(day int, month int, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Weekday().String()
}
