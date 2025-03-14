package sync

import (
	"fmt"
	"github.com/petermattis/goid"
	"go-base/sync/hello"
	"sync"
	"testing"
)

type item struct {
	value int
}

func TestName(t *testing.T) {
	var coll = &sync.Pool{New: func() interface{} {

		return "hello, shenzhen"

	}}

	a := "xiasha"

	coll.Put(a)

	b := coll.Get()

	fmt.Println("第一次取出的东西是：", b)

	c := coll.Get()

	fmt.Println("第二次取出的东西是：", c)

	d := coll.Get()

	fmt.Println("第三次取出的东西是：", d)
}

func TestHello(t *testing.T) {
	hello.Hello()
}

type Obj struct {
	mu sync.Mutex
	// ... 其他字段
}

func (o Obj) Lock() {
	//o.mu.Lock()
	fmt.Printf("Current goroutine ID: %d\n", goid.Get())
}
func (o Obj) Dosomething() {}
func (o Obj) Unlock() {
	fmt.Printf("Current goroutine ID: %d\n", goid.Get())
	//o.mu.Unlock()
	fmt.Printf("Current goroutine ID: %d\n", goid.Get())
}
func TestCopy(t *testing.T) {
	var wg sync.WaitGroup
	fmt.Printf("Current goroutine ID: %d\n", goid.Get())
	wg.Add(1)
	doSomething(&wg)
	wg.Wait()

}

func doSomething(wg *sync.WaitGroup) {
	fmt.Printf("Current goroutine ID: %d\n", goid.Get())
	defer wg.Done()
	// ...
}
