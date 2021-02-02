package stack

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Stack interface {
	Push(value interface{}) int
	Pop() interface{}
	Top() interface{}
	Flush()
	IsEmpty() bool
}

// 1、检查给定的字符串是否是成对出现的
func CheckMultiBracket(in string) bool {
	if len(in) <= 0 {
		return false
	}

	brackets := map[string]string{")": "(", "}": "{", "]": "["}
	s := NewArrayStack()

	for _, bracket := range []rune(in) {
		mul, ok := brackets[string(bracket)]
		if !ok {
			s.Push(string(bracket))
			continue
		}

		value := s.Pop()
		if value == nil {
			return false
		}

		if value.(string) != mul {
			return false
		}
	}
	return s.IsEmpty()
}

/**
 ** 2、实现浏览器前进后退功能
 *
 * 需要两个stack，一个用来存储前进的网页Url，另外一个存储后退时的URL
 * 打开新页面时需要将地址push进back的stack中
 * 如果此时需要后退，则蒋back中的页面pop出来放进forward的stack中
 * 如果需要前进，则将forward中的页面pop出来放进back的stack中
 * 如果此时重新打开了一个新的页面则需要将新页面的URL push近back的stack中
 * 同时清除forward中的数据，此时不能在前进之前的页面了
 */
type Browser struct {
	forwards Stack
	backends Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwards: NewArrayStack(),
		backends: NewLinkedStack(),
	}
}

func (b *Browser) CanForward() bool {
	return !b.forwards.IsEmpty()
}

func (b *Browser) CanBackend() bool {
	return !b.backends.IsEmpty()
}

func (b *Browser) Open(url string) {
	b.backends.Push(url)
	b.forwards.Flush()
}

func (b *Browser) Push(url string) {
	b.backends.Push(url)
}

func (b *Browser) Forward() {
	if b.forwards.IsEmpty() {
		return
	}
	b.backends.Push(b.forwards.Pop())
}

func (b *Browser) Backend() {
	if b.backends.IsEmpty() {
		return
	}
	b.forwards.Push(b.backends.Pop())
}

/**
 ** 3、模拟算数表达式计算加减乘除运算
 *
 * 需要两个stack,一个存储操作数，一个用来存储操作符
 * 遍历计算表达式，遇到操作数将操作数压入 操作数stack
 * 如果遇到操作符则需要判断；如果操作符的stack为空则直接压入stack
 * 如果操作符的stack不为空，则当前的操作符需要跟操作符stack的顶部操作符做比较
 * 1、如果当前操作符的优先级比stack顶的操作符优先级高 则直接压入操作符stack
 * 2、如果当前操作符的优先级小于等于stack顶操作符的优先级
 *    取出两个操作数跟stack顶操作符做计算；将计算结果压入操作数stack
 *    继续使用当前操作符跟stack顶的操作符做比较，直到条件不成立
 */
var (
	supportOperators = map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
	}
)

type Calculator struct {
	numberStack      Stack
	operatorStack    Stack
	supportOperators map[string]int
	reg              *regexp.Regexp
}

func NewCalculator() *Calculator {
	return &Calculator{
		numberStack:      NewArrayStack(),
		operatorStack:    NewArrayStack(),
		supportOperators: supportOperators,
		reg:              regexp.MustCompile(`[0-9]+|[+\-*/]|\s`),
	}
}

func (c *Calculator) split(express string) (ret []string) {
	for len(express) > 0 {
		split := c.reg.FindString(express)
		express = strings.TrimPrefix(express, split)
		if strings.TrimSpace(split) == "" {
			continue
		}
		ret = append(ret, strings.TrimSpace(split))
	}
	return
}

// 必须传入合法的数字跟指定的操作符
func (c *Calculator) Calculate(express string) (float64, error) {
	splits := c.split(express)
	if len(splits) < 3 {
		return 0, errors.New("invalid express")
	}

	for _, obj := range splits {
		curIndex, ok := c.supportOperators[obj]
		if !ok {
			number, err := strconv.ParseFloat(obj, 64)
			if err != nil {
				return 0, err
			}
			c.numberStack.Push(number)
			continue
		}

		// operator
		c.innerCal(curIndex)
		c.operatorStack.Push(obj)
	}
	return c.result(), nil
}

func (c *Calculator) innerCal(index int) {
	for !c.operatorStack.IsEmpty() {
		topOperator := c.operatorStack.Top().(string)
		topIndex := c.supportOperators[topOperator]
		if topIndex < index {
			break
		}

		c.operatorStack.Pop()
		number2 := c.numberStack.Pop().(float64)
		number1 := c.numberStack.Pop().(float64)
		result := c.operate(number1, number2, topOperator)
		c.numberStack.Push(result)
	}
}

func (c *Calculator) result() float64 {
	for !c.operatorStack.IsEmpty() {
		operator := c.operatorStack.Pop().(string)
		number2 := c.numberStack.Pop().(float64)
		number1 := c.numberStack.Pop().(float64)
		c.numberStack.Push(c.operate(number1, number2, operator))
	}
	return c.numberStack.Pop().(float64)
}

func (c *Calculator) operate(num1, num2 float64, operator string) float64 {
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}
