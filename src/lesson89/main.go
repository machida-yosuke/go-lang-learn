package main

import (
	"fmt"
	"regexp"
)

// regex
func main() {
	// match, _ := regexp.MatchString("A", "ABC") // 第一引数に正規表現、第二引数に対象の文字列
	// fmt.Println(match)

	// re1, _ := regexp.Compile("A") // 正規表現をコンパイル
	// match1 := re1.MatchString("ABC")
	// fmt.Println(match1)

	// re2 := regexp.MustCompile("A") // エラーの場合、Panicになる
	// match2 := re2.MatchString("ABC")
	// fmt.Println(match2)

	// // regexp.MustCompile("\\d")
	// // regexp.MustCompile(`\d`)

	// re3 := regexp.MustCompile(`(?i)abc`) // (?i)は大文字小文字を区別しない
	// match3 := re3.MatchString("ABC")
	// fmt.Println(match3)

	// re4 := regexp.MustCompile(`^ABC$`) // ^は行の先頭、$は行の末尾
	// match4 := re4.MatchString("ABC")
	// fmt.Println(match4)

	// re6 := regexp.MustCompile(`a+b*`) // +は1回以上、*は0回以上

	// fmt.Println(re6.MatchString("aaabbb"))
	// fmt.Println(re6.MatchString("a"))
	// fmt.Println(re6.MatchString("b"))

	// re8 := regexp.MustCompile(`[XYZ]`) // []内のいずれかの文字
	// fmt.Println(re8.MatchString("X"))

	// re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`) // []内のいずれかの文字以外
	// fmt.Println(re9.MatchString("123"))
	// fmt.Println(re9.MatchString("12"))
	// fmt.Println(re9.MatchString("12z"))
	// fmt.Println(re9.MatchString("ああああ"))

	// re10 := regexp.MustCompile(`[^0-9A-Za-z_]`) // []内のいずれかの文字以外
	// fmt.Println(re10.MatchString("123"))
	// fmt.Println(re10.MatchString("ああ"))

	// re11 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`) // グループ化
	// fmt.Println(re11.MatchString("abc"))
	// fmt.Println(re11.MatchString("abXZ"))
	// fmt.Println(re11.MatchString("abcxyz"))

	// re := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))

	// fs1 := re.FindString("AAAAABCXYZZZZ")
	// fmt.Println(fs1)

	// fs2 := re.FindAllString("ABCXYZabcXYZ", -1)
	// fmt.Println(fs2)

	// re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	// s := `
	// 	2021-01-01
	// 	2022-02-02
	// 	2023-03-03
	// 	1-1+111
	// `

	// ms := re1.FindAllStringSubmatch(s, -1)
	// fmt.Println(ms)

	// for _, v := range ms {
	// 	fmt.Println(v[0])
	// 	fmt.Println(v[1])
	// 	fmt.Println(v[2])
	// 	fmt.Println(v[3])
	// }

	re1 := regexp.MustCompile(`\s+`)
	fmt.Println(re1.ReplaceAllString("AA BB CCC", ",")) // AA,BB,CCC

	re2 := regexp.MustCompile(`、|。`)
	fmt.Println(re2.ReplaceAllString("わたしは、ごー。", ",")) // AA,BB,CCC

	re3 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(re3.Split("dafasABCXYZafasdabcXYZfaf", -1)) // [dafas afasd faf]

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("AA BB CCC", -1)) // [AA BB CCC]
}
