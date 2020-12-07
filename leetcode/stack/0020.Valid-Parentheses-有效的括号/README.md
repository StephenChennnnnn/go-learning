# [20.Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)

## 题目

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
  1. Open brackets must be closed by the same type of brackets.
  2. Open brackets must be closed in the correct order.
  
Example 1:
```
Input: s = "()"
Output: true
```

Example 2:
```
Input: s = "()[]{}"
Output: true
```

Example 3:
```
Input: s = "(]"
Output: false
```

Example 4:
```
Input: s = "([)]"
Output: false
```

Example 5:
```
Input: s = "{[]}"
Output: true
```

Constraints:
  - 1 <= s.length <= 104
  - s consists of parentheses only '()[]{}'.

## 中文

给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串，判断字符串是否有效。

有效字符串需满足：
  1. 左括号必须用相同类型的右括号闭合。
  2. 左括号必须以正确的顺序闭合。
  
注意空字符串可被认为是有效字符串。

