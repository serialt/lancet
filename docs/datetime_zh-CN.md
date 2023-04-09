# Datetime
datetime日期时间处理包，格式化日期，比较日期。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/datetime/datetime.go](https://github.com/duke-git/lancet/blob/main/datetime/datetime.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/serialt/lancet/datetime"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [AddDay](#AddDay)
- [AddHour](#AddHour)
- [AddMinute](#AddMinute)
- [BeginOfMinute](#BeginOfMinute)
- [BeginOfHour](#BeginOfHour)
- [BeginOfDay](#BeginOfDay)
- [BeginOfWeek](#BeginOfWeek)
- [BeginOfMonth](#BeginOfMonth)
- [BeginOfYear](#BeginOfYear)
- [EndOfMinute](#EndOfMinute)
- [EndOfHour](#EndOfHour)
- [EndOfDay](#EndOfDay)
- [EndOfWeek](#EndOfWeek)
- [EndOfMonth](#EndOfMonth)
- [EndOfYear](#EndOfYear)
- [GetNowDate](#GetNowDate)
- [GetNowTime](#GetNowTime)
- [GetNowDateTime](#GetNowDateTime)
- [GetZeroHourTimestamp](#GetZeroHourTimestamp)
- [GetNightTimestamp](#GetNightTimestamp)
- [FormatTimeToStr](#FormatTimeToStr)
- [FormatStrToTime](#FormatStrToTime)
- [NewUnixNow](#NewUnixNow)
- [NewUnix](#NewUnix)
- [NewFormat](#NewFormat)
- [NewISO8601](#NewISO8601)
- [ToUnix](#ToUnix)
- [ToFormat](#ToFormat)
- [ToFormatForTpl](#ToFormatForTpl)
- [ToIso8601](#ToIso8601)

<div STYLE="page-break-after: always;"></div>

## 文档

## 注:
1. 方法FormatTimeToStr和FormatStrToTime中的format参数值需要传以下类型之一：
- yyyy-mm-dd hh:mm:ss
- yyyy-mm-dd hh:mm
- yyyy-mm-dd hh
- yyyy-mm-dd
- yyyy-mm
- mm-dd
- dd-mm-yy hh:mm:ss
- yyyy/mm/dd hh:mm:ss
- yyyy/mm/dd hh:mm
- yyyy-mm-dd hh
- yyyy/mm/dd
- yyyy/mm
- mm/dd
- dd/mm/yy hh:mm:ss
- yyyy
- mm
- hh:mm:ss
- mm:ss


### <span id="AddDay">AddDay</span>

<p>将日期加/减天数。</p>

<b>Signature:</b>

```go
func AddDay(t time.Time, day int64) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()

    tomorrow := datetime.AddDay(now, 1)
    diff1 := tomorrow.Sub(now)

    yesterday := datetime.AddDay(now, -1)
    diff2 := yesterday.Sub(now)

    fmt.Println(diff1)
    fmt.Println(diff2)

    // Output:
    // 24h0m0s
    // -24h0m0s
}
```

### <span id="AddHour">AddHour</span>

<p>将日期加/减小时数。</p>

<b>Signature:</b>

```go
func AddHour(t time.Time, hour int64) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()

    after2Hours := datetime.AddHour(now, 2)
    diff1 := after2Hours.Sub(now)

    before2Hours := datetime.AddHour(now, -2)
    diff2 := before2Hours.Sub(now)

    fmt.Println(diff1)
    fmt.Println(diff2)

    // Output:
    // 2h0m0s
    // -2h0m0s
}
```

### <span id="AddMinute">AddMinute</span>

<p>将日期加/减分钟数。</p>

<b>Signature:</b>

```go
func AddMinute(t time.Time, minute int64) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()

    after2Minutes := datetime.AddMinute(now, 2)
    diff1 := after2Minutes.Sub(now)

    before2Minutes := datetime.AddMinute(now, -2)
    diff2 := before2Minutes.Sub(now)

    fmt.Println(diff1)
    fmt.Println(diff2)

    // Output:
    // 2m0s
    // -2m0s
}
```

### <span id="BeginOfMinute">BeginOfMinute</span>

<p>返回指定时间的分钟开始时间。</p>

<b>Signature:</b>

```go
func BeginOfMinute(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfMinute(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:50:00 +0000 UTC
}
```

### <span id="BeginOfHour">BeginOfHour</span>

<p>返回指定时间的小时开始时间。</p>

<b>Signature:</b>

```go
func BeginOfHour(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfHour(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:00:00 +0000 UTC
}
```

### <span id="BeginOfDay">BeginOfDay</span>

<p>返回指定时间的当天开始时间。</p>

<b>Signature:</b>

```go
func BeginOfDay(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfDay(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 00:00:00 +0000 UTC
}
```

### <span id="BeginOfWeek">BeginOfWeek</span>

<p>返回指定时间的每周开始时间,默认开始时间星期日。</p>

<b>Signature:</b>

```go
func BeginOfWeek(t time.Time, beginFrom ...time.Weekday) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfWeek(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 00:00:00 +0000 UTC
}
```

### <span id="BeginOfMonth">BeginOfMonth</span>

<p>返回指定时间的当月开始时间。</p>

<b>Signature:</b>

```go
func BeginOfMonth(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfMonth(input)

    fmt.Println(result)

    // Output:
    // 2023-01-01 00:00:00 +0000 UTC
}
```

### <span id="BeginOfYear">BeginOfYear</span>

<p>返回指定时间的当年开始时间</p>

<b>Signature:</b>

```go
func BeginOfYear(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfYear(input)

    fmt.Println(result)

    // Output:
    // 2023-01-01 00:00:00 +0000 UTC
}
```

### <span id="EndOfMinute">EndOfMinute</span>

<p>返回指定时间的分钟结束时间。</p>

<b>Signature:</b>

```go
func EndOfMinute(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfMinute(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:50:59.999999999 +0000 UTC
}
```

### <span id="EndOfHour">EndOfHour</span>

<p>返回指定时间的小时结束时间。</p>

<b>Signature:</b>

```go
func EndOfHour(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfHour(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfDay">EndOfDay</span>

<p>返回指定时间的当天结束时间。</p>

<b>Signature:</b>

```go
func EndOfDay(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfDay(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 23:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfWeek">EndOfWeek</span>

<p>返回指定时间的星期结束时间,默认结束时间星期六。</p>

<b>Signature:</b>

```go
func EndOfWeek(t time.Time, endWith ...time.Weekday) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfWeek(input)

    fmt.Println(result)

    // Output:
    // 2023-01-14 23:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfMonth">EndOfMonth</span>

<p>返回指定时间的当月结束时间。</p>

<b>Signature:</b>

```go
func EndOfMonth(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfMonth(input)

    fmt.Println(result)

    // Output:
    // 2023-01-31 23:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfYear">EndOfYear</span>

<p>返回指定时间的当年结束时间。</p>

<b>Signature:</b>

```go
func EndOfYear(t time.Time) time.Time
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfYear(input)

    fmt.Println(result)

    // Output:
    // 2023-12-31 23:59:59.999999999 +0000 UTC
}
```

### <span id="GetNowDate">GetNowDate</span>

<p>获取当天日期，返回格式：yyyy-mm-dd。</p>

<b>Signature:</b>

```go
func GetNowDate() string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()
    currentDate := datetime.GetNowDate()

    fmt.Println(currentDate) 
    
    // Output:
    // 2022-01-28
}
```

### <span id="GetNowTime">GetNowTime</span>

<p>获取当时时间，返回格式：hh:mm:ss</p>

<b>Signature:</b>

```go
func GetNowTime() string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()
    currentTime := datetime.GetNowTime()

    fmt.Println(currentTime) // 15:57:33

    // Output:
    // 15:57:33
}
```

### <span id="GetNowDateTime">GetNowDateTime</span>

<p>获取当时日期和时间，返回格式：yyyy-mm-dd hh:mm:ss。</p>

<b>Signature:</b>

```go
func GetNowDateTime() string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()
    current := datetime.GetNowDateTime()

    fmt.Println(current) 
    
    // Output:
    // 2022-01-28 15:59:33
}
```

### <span id="GetZeroHourTimestamp">GetZeroHourTimestamp</span>

<p>获取零点时间戳(timestamp of 00:00)</p>

<b>Signature:</b>

```go
func GetZeroHourTimestamp() int64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()
    zeroTime := datetime.GetZeroHourTimestamp()
    
    fmt.Println(zeroTime) 
    
    // Output:
    // 1643299200
}
```

### <span id="GetNightTimestamp">GetNightTimestamp</span>

<p>获取午夜时间戳(timestamp of 23:59)。</p>

<b>Signature:</b>

```go
func GetNightTimestamp() int64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    now := time.Now()
    nightTime := datetime.GetNightTimestamp()

    fmt.Println(nightTime) 
    
    // Output:
    // 1643385599
}
```

### <span id="FormatTimeToStr">FormatTimeToStr</span>

<p>将日期格式化成字符串，`format` 参数格式参考注1。</p>

<b>Signature:</b>

```go
func FormatTimeToStr(t time.Time, format string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/serialt/lancet/datetime"
)

func main() {
    t, _ := time.Parse("2006-01-02 15:04:05", "2021-01-02 16:04:08")

    result1 := datetime.FormatTimeToStr(t, "yyyy-mm-dd hh:mm:ss")
    result2 := datetime.FormatTimeToStr(t, "yyyy-mm-dd")
    result3 := datetime.FormatTimeToStr(t, "dd-mm-yy hh:mm:ss")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 2021-01-02 16:04:08
    // 2021-01-02
    // 02-01-21 16:04:08
}
```

### <span id="FormatStrToTime">FormatStrToTime</span>

<p>将字符串格式化成时间，`format` 参数格式参考注1。</p>

<b>Signature:</b>

```go
func FormatStrToTime(str, format string) (time.Time, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    result1, _ := datetime.FormatStrToTime("2021-01-02 16:04:08", "yyyy-mm-dd hh:mm:ss")
    result2, _ := datetime.FormatStrToTime("2021-01-02", "yyyy-mm-dd")
    result3, _ := datetime.FormatStrToTime("02-01-21 16:04:08", "dd-mm-yy hh:mm:ss")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 2021-01-02 16:04:08 +0000 UTC
    // 2021-01-02 00:00:00 +0000 UTC
    // 2021-01-02 16:04:08 +0000 UTC
}
```

### <span id="NewUnixNow">NewUnixNow</span>

<p>创建一个当前时间的unix时间戳。</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewUnixNow() *theTime
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm := datetime.NewUnixNow()
    fmt.Println(tm) 
    
    // Output:
    // &{1647597438}
}
```

### <span id="NewUnix">NewUnix</span>

<p>创建一个unix时间戳。</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewUnix(unix int64) *theTime
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm := datetime.NewUnix(1647597438)
    fmt.Println(tm) 
    
    // Output:
    // &{1647597438}
}
```

### <span id="NewFormat">NewFormat</span>

<p>创建一个yyyy-mm-dd hh:mm:ss格式时间字符串的unix时间戳。</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewFormat(t string) (*theTime, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm, err := datetime.NewFormat("2022-03-18 17:04:05")
    fmt.Println(tm) 
    
    // Output:
    // &{1647594245}
}
```

### <span id="NewISO8601">NewISO8601</span>

<p>创建一个iso8601格式时间字符串的unix时间戳。</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewISO8601(iso8601 string) (*theTime, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm, err := datetime.NewISO8601("2006-01-02T15:04:05.999Z")
    fmt.Println(tm) 
    
    // Output:
    // &{1136214245}
}
```

### <span id="ToUnix">ToUnix</span>

<p>返回unix时间戳。</p>

<b>Signature:</b>

```go
func (t *theTime) ToUnix() int64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm := datetime.NewUnixNow()
    fmt.Println(tm.ToUnix()) 
    
    // Output:
    // 1647597438
}
```

### <span id="ToFormat">ToFormat</span>

<p>返回格式'yyyy-mm-dd hh:mm:ss'的日期字符串。</p>

<b>Signature:</b>

```go
func (t *theTime) ToFormat() string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm, _ := datetime.NewFormat("2022-03-18 17:04:05")
    fmt.Println(tm.ToFormat()) 
    
    // Output:
    // 2022-03-18 17:04:05
}
```

### <span id="ToFormatForTpl">ToFormatForTpl</span>

<p>返回tpl格式指定的日期字符串。</p>

<b>Signature:</b>

```go
func (t *theTime) ToFormatForTpl(tpl string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm, _ := datetime.NewFormat("2022-03-18 17:04:05")
    ts := tm.ToFormatForTpl("2006/01/02 15:04:05")
    fmt.Println(ts) 
    
    // Output:
    // 2022/03/18 17:04:05
}
```

### <span id="ToIso8601">ToIso8601</span>

<p>返回iso8601日期字符串。</p>

<b>Signature:</b>

```go
func (t *theTime) ToIso8601() string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/serialt/lancet/datetime"
)

func main() {
    tm, _ := datetime.NewISO8601("2006-01-02T15:04:05.999Z")
    ts := tm.ToIso8601()
    fmt.Println(ts) 
    
    // Output:
    // 2006-01-02T23:04:05+08:00
}
```