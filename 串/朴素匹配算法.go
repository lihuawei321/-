package main

import (
	"errors"
	"fmt"
)

/*
串的最常见操作就是查找子串，也称为模式匹配。

比如要从S="goodgoogle"中找到T="google"这个子串的位置，通常需要如下步骤：

1 索引从主串S第一位开始匹配，前三个匹配成功，第四个字母失败，索引返回第二个位置
2 索引从主串S第二位开始重新匹配，匹配失败，依次类推
3 到达google匹配成功
在上述案例中，n为主串长度，m为要匹配的子串长度，根据等概率原则，平均是(n+m)/2次查找，时间复杂度为O(n+m)，
在最坏的情况下，比如S="00000001"，要匹配T="01"，时间复杂度是O(m*(n-m+1))，这种算法的低效程度令人发指。*/

// 给定字符串s，子串 t，查找t在str中的位置（索引）
func Index(s string, t string) (pos int, err error) {

	sLength := len(s)
	tLength := len(t)

	if sLength == 0 || tLength == 0 || sLength < tLength {
		err = errors.New("字符串长度不合法")
		return
	}

	for i := 0; i <= sLength-tLength; i++ {

		// s 截取和 t 一样长度的子串
		sonStr := s[i : i+tLength]
		if sonStr != t {
			err = errors.New("未匹配成功")
			continue
		} else {
			pos = i
			err = nil
			break
		}
	}

	return

}

//不利用go特性
// 给定字符串s，子串 t，查找t在str中的位置（索引）
func Index2(s string, t string) (pos int, err error) {

	sLength := len(s)
	tLength := len(t)
	if sLength == 0 || tLength == 0 || sLength < tLength {
		return 0, errors.New("字符串长度不合法")
	}

	for i := 0; i <= sLength-tLength; i++ {

		pos = i

		for j := 0; j < tLength; j++ {

			if s[i] != t[j] {
				i = i - j + 1 // 与到不相等 i 回退
				pos = 0
				err = errors.New("未匹配到数据")
				break
			} else {
				i++
				err = nil
				continue
			}
		}

	}

	return
}

func main() {
	s1 := "goodgoogle"
	s2 := "google"

	r, e := Index2(s1, s2)

	fmt.Println("最终e:", e)
	fmt.Println("最终r:", r)
}
