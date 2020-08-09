# defanging-an-ip-address

## 題目解讀：

### 題目來源:
[defanging-an-ip-address](https://leetcode.com/problems/defanging-an-ip-address/)

### 原文:
Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".

### 解讀：

給定一個合法的ip字串 address 舉例來說: "127.0.0.1"

返回一個字串把"."代換為"[.]" 舉例來說: "127.0.0.1" => "127[.]0[.]0[.]1"

## 初步解法:
### 初步觀察:

首先只需要先把 輸入的address 對 "." 做split

然後在用"[.]" join為回去

### 初步設計:

Given a valid ipv4 string address

set a empty string formatedStr = ""

set splitedAdrs = address.split(".")

formatedStr = splitedAdrs.join("[.]")

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package defang_ip

import "strings"

func defangIPaddr(address string) string {
	ret := ""
	splitedAddr := strings.Split(address, ".")
	ret = strings.Join(splitedAddr, "[.]")
	return ret
}

```
## 測資的撰寫

```golang
package defang_ip

import "testing"

func Test_defangIPaddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				address: "1.1.1.1",
			},
			want: "1[.]1[.]1[.]1",
		},
		{
			name: "Example2",
			args: args{
				address: "255.100.50.0",
			},
			want: "255[.]100[.]50[.]0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defangIPaddr(tt.args.address); got != tt.want {
				t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

```

## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)