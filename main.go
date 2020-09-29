package main

/**
* @Author: super
* @Date: 2020-09-29 14:50
* @Description:
**/

const s = "Go101.org"

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

func main() {
	println(a, b) //4 0
	println(len(s)) //9
	println(len(s[:])) //9
}