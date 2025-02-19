package future

import "fmt"

type Void struct{}

type Future[T any] interface {
	Get() (T, error)
	Exceptionally(func(error)) Future[T]
}

type NonResultFuture struct {
	errorChain chan error
}

func (f *NonResultFuture) Get() (void Void, err error) {

	if err = <-f.errorChain; err != nil {
		return void, err
	}
	return
}

func (f *NonResultFuture) Exceptionally(fn func(error)) Future[Void] {

	errChan := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()
		if err := <-f.errorChain; err != nil {
			fn(err)
			errChan <- err
		}
	}()
	return &NonResultFuture{errorChain: errChan}
}

func RunAsync(runnable Runnable) Future[Void] {
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()
		err := runnable()
		if err != nil {
			errChan <- err
		}
	}()

	return &NonResultFuture{errorChain: errChan}
}
