package errUtil

import (
	"fmt"
	"path"
	"runtime"
)

func WithCaller(err error) error {
	if err == nil {
		return nil
	}
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file)

	return fmt.Errorf("fileName:%s, line:%d, funcName:%s, msg: %w", fileName, line, funcName, err)
}

func MsgWithCaller(statement string, arg ...any) error {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf(statement, arg...)
	}

	msg := fmt.Sprintf(statement, arg...)
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file)
	return fmt.Errorf("fileName:%s, line:%d, funcName:%s, msg: %s", fileName, line, funcName, msg)
}
