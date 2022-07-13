package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unsafe"
)

func main() {
	var a, b []byte
	a = []byte("aaabbb")
	b = []byte("bbb")
	// a<b return -1 , a>b return 1, a==b return 0
	// -1
	fmt.Println(bytes.Compare(a, b))
	// true
	fmt.Println(bytes.Contains(a, b))

	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!"))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的."))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), ""))
	fmt.Println(bytes.ContainsAny([]byte(""), ""))

	// 计算非重复的切片数量
	a = []byte("aaabbbb")
	b = []byte("bbb")
	// 1
	fmt.Println(bytes.Count(a, b))
	a = []byte("aaabbbbbb")
	b = []byte("bbb")
	// 2
	fmt.Println(bytes.Count(a, b))
	a = []byte("aaabbbbbb")
	b = []byte("")
	// 10 len(a)+1
	fmt.Println(bytes.Count(a, b))

	a = []byte("aaabbbbbccc")
	b = []byte("bbb")
	// aaa bbccc true
	c, d, result := bytes.Cut(a, b)
	fmt.Println(string(c), string(d), result)

	// true false
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))

	a = []byte("aaabbb")
	b = []byte("bbb")
	// false
	fmt.Println(bytes.EqualFold(a, b))

	// 返回utf-8码点 使用'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0 分割
	fmt.Println(bytes.Fields([]byte("  foo bar  baz  ")))

	a = []byte("aaabbbccc")
	b = []byte("aaa")
	c = []byte("ccc")
	// true true
	fmt.Println(bytes.HasPrefix(a, b))
	fmt.Println(bytes.HasSuffix(a, c))

	// 0 6 -1
	fmt.Println(bytes.Index(a, b))
	fmt.Println(bytes.Index(a, c))
	fmt.Println(bytes.Index(a, []byte("ddd")))

	// 3 -1
	fmt.Println(bytes.IndexByte([]byte("aaabbb"), 'b'))
	fmt.Println(bytes.IndexByte([]byte("aaabbb"), 'c'))

	a = []byte("aaabbbcccaaa")
	b = []byte("aaa")
	// 9
	fmt.Println(bytes.LastIndex(a, b))

	// bytes.Index()
	// 将二维字节切片使用指定byte连接
	var s [][]byte = [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Println(string(bytes.Join(s, []byte("-"))))

	// 3
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))

	// Map
	a = []byte("aaabbbccc")
	b = bytes.Map(func(r rune) rune { return r + 2 }, a)
	// cccdddeee
	fmt.Println(string(b), unsafe.Pointer(&b), unsafe.Pointer(&a))

	//repeat -----
	fmt.Println(string(bytes.Repeat([]byte("-"), 5)))

	// Replace
	a = []byte("aaabbbccc")
	// 第三个参数 返回切片的副本
	b = bytes.Replace(a, []byte("bbb"), []byte("ddd"), -1)
	// aaadddccc aaabbbccc
	fmt.Println(string(b), string(a))
	b = bytes.Replace(a, nil, []byte("xx"), -1)
	// xxaxxaxxaxxbxxbxxbxxcxxcxxcxx aaabbbccc
	fmt.Println(string(b), string(a))

	// Split
	// aaa bccc
	a = []byte("aaa,bbb,ccc")
	s = bytes.Split(a, []byte(","))
	fmt.Printf("%q\n", s)

	// 逐行打印
	// s = bytes.Split(a, nil)
	// for _, v := range s {
	// 	fmt.Println(string(v))
	// }

	// SplitAfter 原字符串：aaabbbccc 使用bb分割
	// aaabb bccc
	a = []byte("aaa,bbb,ccc")
	s = bytes.SplitAfter(a, []byte(","))
	fmt.Println(string(a))
	fmt.Printf("%q\n", s)

	// SplitAfterN

	// 分割成指定数量的切片 -1代表没有数量限制
	s = bytes.SplitAfterN(a, []byte(","), 100)
	fmt.Println(string(a))
	fmt.Printf("%q\n", s)

	// 转成小写 gopher
	fmt.Println(string(bytes.ToLower([]byte("Gopher"))))

	// 没搞懂
	fmt.Println(string(bytes.ToLowerSpecial(unicode.CaseRanges, []byte("Gopher"))))

	fmt.Println(string(bytes.ToTitle([]byte("gopher is me"))))
	fmt.Println(string(bytes.ToTitleSpecial(unicode.CaseRanges, []byte("gopher is me"))))

	// 转成大写
	fmt.Println(string(bytes.ToUpper([]byte("gopher is me"))))

	// Trim
	a = []byte("a aaa  foo bar  aaa a")
	// 去除首尾指定字符串
	fmt.Println(string(bytes.Trim(a, "a")))

	// TrimFunc
	fmt.Println(string(bytes.TrimFunc(a, func(r rune) bool {
		return r == 'b'
	})))

	fmt.Println(string(bytes.TrimLeft(a, "a")))
	fmt.Println(string(bytes.TrimRight(a, "a")))

	fmt.Println(string(bytes.TrimSpace(a)))
	fmt.Println(string(bytes.TrimSuffix(a, []byte("aaa a"))))
} // end main
