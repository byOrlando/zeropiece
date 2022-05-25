package common

import "fmt"

// Error is a function that will prevent the error from being thrown out.
func Error() func() {
	return func() {
		if err := recover(); err != nil {
			LOG.Error(fmt.Sprintf("发生了错误%v", err))
		}
	}
}
