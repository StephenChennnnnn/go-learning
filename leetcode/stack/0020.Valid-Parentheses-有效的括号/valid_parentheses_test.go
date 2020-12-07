package leetcode

import (
	"fmt"
	"testing"
)

type question20 struct {
	param20 string
	ans20   bool
}

func Test_P20_ValidParentheses(t *testing.T) {
	qs := []question20{
		{param20: "()[]{}", ans20: true},
		{param20: "(]", ans20: false},
		{param20: "({[]})", ans20: true},
		{param20: "(){[({[]})]}", ans20: true},
		{param20: "((([[[{{{", ans20: false},
		{param20: "(())]]", ans20: false},
		{param20: "", ans20: true},
		{param20: "[", ans20: false},
	}

	fmt.Println("--------20 Valid Parentheses--------")

	for _, q := range qs {
		fmt.Printf("[Input]: %v		[output]: %v\n", q.param20, isValid(q.param20))
	}
}
