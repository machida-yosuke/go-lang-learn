package main

import (
	"fmt"
	"strings"
)

func main() {
	// s1 := strings.Join([]string{"a", "b", "c"}, ",")
	// fmt.Println(s1)
	// s2 := strings.Join([]string{"a", "b", "c"}, "")
	// fmt.Println(s2)

	// i1 := strings.Index("abc", "b")
	// fmt.Println(i1)

	// i2 := strings.Index("abcabc", "b") // 最初のbの位置
	// fmt.Println(i2)

	// i3 := strings.LastIndex("abcabc", "b") // 最後のbの位置
	// fmt.Println(i3)

	// i4 := strings.IndexAny("abc", "bc") //
	// fmt.Println(i4)

	// b1 := strings.HasPrefix("abc", "a") // aから始まっているか
	// fmt.Println(b1)

	// b2 := strings.HasSuffix("abc", "c") // cで終わっているか
	// fmt.Println(b2)

	// b3 := strings.Contains("abc", "b") // 含まれているか
	// fmt.Println(b3)
	// b4 := strings.ContainsAny("abc", "ce") // いずれか含まれているか
	// fmt.Println(b4)

	// i6 := strings.Count("abcabc", "a") // aの数
	// fmt.Println(i6)

	// s3 := strings.Repeat("abc", 3) // 3回繰り返す
	// fmt.Println(s3)

	// s5 := strings.Replace("abcabc", "a", "z", 1) // 1回だけ置換
	// fmt.Println(s5)
	// s6 := strings.Replace("abcabc", "a", "z", -1) // 全て置換
	// fmt.Println(s6)

	// s7 := strings.Split("a,b,c", ",") // ,で分割
	// fmt.Println(s7)

	// s8 := strings.SplitAfter("a,b,c", ",") // ,の後ろで分割
	// fmt.Println(s8)

	// s9 := strings.SplitN("a,b,c", ",", 2)
	// fmt.Println(s9)

	// s10 := strings.SplitAfterN("a,b,c", ",", 1)
	// fmt.Println(s10)

	// s11 := strings.ToLower("ABC")
	// fmt.Println(s11)

	// s12 := strings.ToUpper("abc")
	// fmt.Println(s12)

	s15 := strings.TrimSpace("  abc  def  ") // 前後のスペースを削除
	fmt.Println(s15)

	s16 := strings.TrimSpace("　　　あああ　　いいいい　　ああ") // 前後のスペースを削除
	fmt.Println(s16)

	s18 := strings.Fields("abc def ghi") // スペースで分割
	fmt.Println(s18)
	fmt.Println(s18[0])
}
