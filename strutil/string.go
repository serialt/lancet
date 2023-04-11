// Use of this source code is governed by MIT license

// Package strutil implements some functions to manipulate string.
package strutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CamelCase coverts string to camelCase string. Non letters and numbers will be ignored.
// Play: https://go.dev/play/p/9eXP3tn2tUy
func CamelCase(s string) string {
	var builder strings.Builder

	strs := splitIntoStrings(s, false)
	for i, str := range strs {
		if i == 0 {
			builder.WriteString(strings.ToLower(str))
		} else {
			builder.WriteString(Capitalize(str))
		}
	}

	return builder.String()
}

// Capitalize converts the first character of a string to upper case and the remaining to lower case.
// Play: https://go.dev/play/p/2OAjgbmAqHZ
func Capitalize(s string) string {
	result := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			result[i] = unicode.ToUpper(v)
		} else {
			result[i] = unicode.ToLower(v)
		}
	}

	return string(result)
}

// UpperFirst converts the first character of string to upper case.
// Play: https://go.dev/play/p/sBbBxRbs8MM
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)

	return string(r) + s[size:]
}

// LowerFirst converts the first character of string to lower case.
// Play: https://go.dev/play/p/CbzAyZmtJwL
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)

	return string(r) + s[size:]
}

// PadStart pads string on the left and right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/NzImQq-VF8q
func Pad(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 0)
}

// PadStart pads string on the left side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/xpTfzArDfvT
func PadStart(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 1)
}

// PadEnd pads string on the right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/9xP8rN0vz--
func PadEnd(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 2)
}

// KebabCase coverts string to kebab-case, non letters and numbers will be ignored.
// Play: https://go.dev/play/p/dcZM9Oahw-Y
func KebabCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "-")
}

// UpperKebabCase coverts string to upper KEBAB-CASE, non letters and numbers will be ignored
// Play: https://go.dev/play/p/zDyKNneyQXk
func UpperKebabCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "-")
}

// SnakeCase coverts string to snake_case, non letters and numbers will be ignored
// Play: https://go.dev/play/p/tgzQG11qBuN
func SnakeCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "_")
}

// UpperSnakeCase coverts string to upper SNAKE_CASE, non letters and numbers will be ignored
// Play: https://go.dev/play/p/4COPHpnLx38
func UpperSnakeCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "_")
}

// Before returns the substring of the source string up to the first occurrence of the specified string.
// Play: https://go.dev/play/p/JAWTZDS4F5w
func Before(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[0:i]
}

// BeforeLast returns the substring of the source string up to the last occurrence of the specified string.
// Play: https://go.dev/play/p/pJfXXAoG_Te
func BeforeLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[0:i]
}

// After returns the substring after the first occurrence of a specified string in the source string.
// Play: https://go.dev/play/p/RbCOQqCDA7m
func After(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[i+len(char):]
}

// AfterLast returns the substring after the last occurrence of a specified string in the source string.
// Play: https://go.dev/play/p/1TegARrb8Yn
func AfterLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[i+len(char):]
}

// IsString check if the value data type is string or not.
// Play: https://go.dev/play/p/IOgq7oF9ERm
func IsString(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

// Reverse returns string whose char order is reversed to the given string.
// Play: https://go.dev/play/p/adfwalJiecD
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Wrap a string with given string.
// Play: https://go.dev/play/p/KoZOlZDDt9y
func Wrap(str string, wrapWith string) string {
	if str == "" || wrapWith == "" {
		return str
	}
	var sb strings.Builder
	sb.WriteString(wrapWith)
	sb.WriteString(str)
	sb.WriteString(wrapWith)

	return sb.String()
}

// Unwrap a given string from anther string. will change source string.
// Play: https://go.dev/play/p/Ec2q4BzCpG-
func Unwrap(str string, wrapToken string) string {
	if str == "" || wrapToken == "" {
		return str
	}

	firstIndex := strings.Index(str, wrapToken)
	lastIndex := strings.LastIndex(str, wrapToken)

	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapToken) <= lastIndex {
			str = str[len(wrapToken):lastIndex]
		}
	}

	return str
}

// SplitEx split a given string which can control the result slice contains empty string or not.
// Play: https://go.dev/play/p/Us-ySSbWh-3
func SplitEx(s, sep string, removeEmptyString bool) []string {
	if sep == "" {
		return []string{}
	}

	n := strings.Count(s, sep) + 1
	a := make([]string, n)
	n--
	i := 0
	sepSave := 0
	ignore := false

	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		ignore = false
		if removeEmptyString {
			if s[:m+sepSave] == "" {
				ignore = true
			}
		}
		if !ignore {
			a[i] = s[:m+sepSave]
			s = s[m+len(sep):]
			i++
		} else {
			s = s[m+len(sep):]
		}
	}

	var ret []string
	if removeEmptyString {
		if s != "" {
			a[i] = s
			ret = a[:i+1]
		} else {
			ret = a[:i]
		}
	} else {
		a[i] = s
		ret = a[:i+1]
	}

	return ret
}

// Substring returns a substring of the specified length starting at the specified offset position.
// Play: https://go.dev/play/p/q3sM6ehnPDp
func Substring(s string, offset int, length uint) string {
	rs := []rune(s)
	size := len(rs)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}
	if offset > size {
		return ""
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	str := string(rs[offset : offset+int(length)])

	return strings.Replace(str, "\x00", "", -1)
}

// SplitWords splits a string into words, word only contains alphabetic characters.
// Play: https://go.dev/play/p/KLiX4WiysMM
func SplitWords(s string) []string {
	var word string
	var words []string
	var r rune
	var size, pos int

	isWord := false

	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s)

		switch {
		case isLetter(r):
			if !isWord {
				isWord = true
				word = s
				pos = 0
			}

		case isWord && (r == '\'' || r == '-'):
			// is word

		default:
			if isWord {
				isWord = false
				words = append(words, word[:pos])
			}
		}

		pos += size
		s = s[size:]
	}

	if isWord {
		words = append(words, word[:pos])
	}

	return words
}

// WordCount return the number of meaningful word, word only contains alphabetic characters.
// Play: https://go.dev/play/p/bj7_odx3vRf
func WordCount(s string) int {
	var r rune
	var size, count int

	isWord := false

	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s)

		switch {
		case isLetter(r):
			if !isWord {
				isWord = true
				count++
			}

		case isWord && (r == '\'' || r == '-'):
			// is word

		default:
			isWord = false
		}

		s = s[size:]
	}

	return count
}

// RemoveNonPrintable remove non-printable characters from a string.
// Play: https://go.dev/play/p/og47F5x_jTZ
func RemoveNonPrintable(str string) string {
	result := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)

	return result
}

// StringIsEmpty 判断字符串是否为空，是则返回true，否则返回false
func StringIsEmpty(str string) bool {
	return len(str) == 0
}

// StringIsNotEmpty 和 IsEmpty 的语义相反
func StringIsNotEmpty(str string) bool {
	return !StringIsEmpty(str)
}

// StringConvert 下划线转换，首字母小写变大写，
// 下划线去掉并将下划线后的首字母大写
func StringConvert(oriString string) string {
	cb := []byte(oriString)
	em := make([]byte, 0, 10)
	b := false
	for i, by := range cb {
		// 首字母如果是小写，则转换成大写
		if i == 0 && (97 <= by && by <= 122) {
			by = by - 32
		} else if by == 95 {
			// 下一个单词要变成大写
			b = true
			continue
		}
		if b {
			if 97 <= by && by <= 122 {
				by = by - 32
			}
			b = false
		}
		em = append(em, by)
	}
	return string(em)
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// StringRandSeq 创建指定长度的随机字符串
func StringRandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// StringRandSeq16 创建长度为16的随机字符串
func StringRandSeq16() string {
	return StringRandSeq(16)
}

// StringAllLetter 判断字符串是否只由字母组成
func StringIsLetter(str string) (bool, error) {
	return regexp.MatchString(`^[A-Za-z]+$`, str)
}

// StringTrim 去除字符串中的空格和换行符
func StringTrim(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	return StringTrimN(str)
}

// StringTrimN 去除字符串中的换行符
func StringTrimN(str string) string {
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

// ToString 将对象格式化成字符串
func ToString(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("%+v", i)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", i)
	}
	return out.String()
}

// StringSingleValue 将字符串内所有连续value替换为单个value
func StringSingleValue(res string, value string) string {
	doubleValue := StringBuild(value, value)
	for skip := false; !skip; {
		resNew := strings.Replace(res, doubleValue, value, -1)
		if res == resNew {
			skip = true
		}
		res = resNew
	}
	return res
}

// StringSingleSpace 将字符串内所有连续空格替换为单个空格
func StringSingleSpace(res string) string {
	return StringSingleValue(res, " ")
}

// StringPrefixSupplementZero 当字符串长度不满足时，将字符串前几位补充0
func StringPrefixSupplementZero(str string, offset int) string {
	backZero := offset - len(str)
	if backZero <= 0 {
		return str
	}
	for i := 0; i < backZero; i++ {
		str = strings.Join([]string{"0", str}, "")
	}
	return str
}

// SubString 截取字符串
func SubString(res string, start, end int) string {
	if start > end || start < 0 || end > len(res) {
		return ""
	}
	return res[start:end]
}

// StringBuild 拼接字符串
func StringBuild(arrString ...string) string {
	return strings.Join(arrString, "")
}

// StringBuildSep 拼接字符串
func StringBuildSep(sep string, arrString ...string) string {
	return strings.Join(arrString, sep)
}

// FilterPrefix 根据前缀过滤slice
func FilterPrefix(strs []string, s string) (r []string) {
	for _, v := range strs {
		if len(v) >= len(s) {
			if v[:len(s)] == s {
				r = append(r, v)
			}
		}
	}

	return r
}

// FindLongestStr 查询最长字符串
func FindLongestStr(strs []string) string {
	longestStr := ""
	for _, str := range strs {
		if len(str) >= len(longestStr) {
			longestStr = str
		}
	}

	return longestStr
}

// ArrayToString 数字切片变字符串
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
