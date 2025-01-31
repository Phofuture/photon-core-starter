package future

import "fmt"

type Void struct{}

type Future[T any] interface {
	Get() (T, error)
	Exceptionally(func(error)) Future[T]
}

type NonResultFuture struct {
	done       chan int
	errorChain chan error
}

func (f *NonResultFuture) Get() (void Void, err error) {
	<-f.done

	select {
	case err = <-f.errorChain:
		return void, err
	default:
		return
	}
}

func (f *NonResultFuture) Exceptionally(fn func(error)) Future[Void] {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic: ", r)
			}
		}()
		if err := <-f.errorChain; err != nil {
			fn(err)
		}
	}()
	return f
}

func RunAsync(runnable Runnable) Future[Void] {
	res := make(chan int)
	errChan := make(chan error)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
			close(res)
		}()
		err := runnable()
		if err != nil {
			errChan <- err
		}
	}()

	return &NonResultFuture{done: res, errorChain: errChan}
}
