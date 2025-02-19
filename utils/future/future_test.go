package future

import (
	"errors"
	"testing"
	"time"
)

type ErrType struct {
	err error
}

func (e *ErrType) Error() string {
	return e.err.Error()
}

func TestFuture(t *testing.T) {

	//Test code here
	get, err := RunAsync(func() error {
		for i := 0; i < 10; i++ {
			t.Log("Hello", i)
			time.Sleep(1 * time.Second)
		}
		return errors.New("test error")
	}).Exceptionally(func(err error) {
		t.Logf("Exceptionally: %v\n", err)
	}).Get()
	t.Logf("GetError:%v\n", err)
	t.Logf("GetResult:%T\n", get)

	time.Sleep(30 * time.Second)
}
