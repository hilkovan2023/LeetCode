package main

import "fmt"

// 通过接口，实现高内聚低耦合
type Op interface {
	Key() string
	Value() interface{}
}

// ch、m都是局部变量，ch、m被其他协程使用，发生内存逃逸，防止ch、m被回收
// 数据存放到 m 中，通过channel进行数据的set和get
func server() chan<- Op {
	ch := make(chan Op)
	m := map[string]string{}
	go func() {
		for op := range ch {
			k, v := op.Key(), op.Value()
			s, set := v.(string)
			if set {
				m[k] = s
				continue
			}
			v.(chan string) <- m[k]
			continue
		}
	}()
	return ch
}

type s struct {
	key   string
	value string
}

func (x *s) Key() string        { return x.key }
func (x *s) Value() interface{} { return x.value }

func setOp(key, value string) Op {
	return &s{key, value}
}

type g struct {
	key   string
	value chan string
}

func (x *g) Key() string        { return x.key }
func (x *g) Value() interface{} { return x.value }

func getOp(key string) Op {
	return &g{key, make(chan string)}
}

func main() {
	srv := server()
	srv <- setOp("key", "value")
	get := getOp("key")
	srv <- get
	fmt.Println(<-(get.(*g).value))
}
