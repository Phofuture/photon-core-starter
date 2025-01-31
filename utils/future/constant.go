package future

type (
	Runnable                        func() error
	Consumer[T any]                 func(T) error
	Supplier[T any]                 func() (T, error)
	Function[T any, K any]          func(T) (K, error)
	BiFunction[T any, K any, R any] func(T, K) (R, error)
)
