package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	code, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	is, ds := len(code), 30000
	if len(os.Args) > 2 {
		ds, _ = strconv.Atoi(os.Args[2])
	}
	var data = make([]byte, ds)
	var jump = make(map[int]int)
	var jumpStack []int
	for ip, jp := 0, 0; ip < is; ip++ {
		switch code[ip] {
		case '[':
			jumpStack = append(jumpStack, ip)
		case ']':
			if len(jumpStack) == 0 {
				os.Exit(1)
			}
			jp, jumpStack = jumpStack[len(jumpStack)-1], jumpStack[:len(jumpStack)-1]
			jump[ip], jump[jp] = jp, ip
		}
	}
	for ip, dp := 0, 0; ip < is && dp > -1 && dp < ds; ip++ {
		switch code[ip] {
		case '>':
			dp++
		case '<':
			dp--
		case '+':
			data[dp]++
		case '-':
			data[dp]--
		case '.':
			fmt.Printf("%c", data[dp])
		case ',':
			_, _ = os.Stdin.Read(data[dp : dp+1])
		case '[':
			if data[dp] == 0 {
				ip = jump[ip]
			}
		case ']':
			if data[dp] != 0 {
				ip = jump[ip]
			}
		}
	}
}
