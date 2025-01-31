package future

import (
	"testing"
)

func TestFuture(t *testing.T) {
	// Test code here
	//get, err := RunAsync(func() error {
	//	for i := 0; i < 10; i++ {
	//		t.Log("Hello", i)
	//		time.Sleep(1 * time.Second)
	//	}
	//	panic("Panic")
	//}).Exceptionally(func(err error) {
	//	t.Log(err)
	//}).Get()
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//t.Logf("%T\n", get)
	//
	//time.Sleep(30 * time.Second)
}
