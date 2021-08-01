package internal

import "net/http"

// GlobalError 全局异常结构体
type GlobalError struct {
	Status  int `json:"-"`
	Code    int `json:"code"`
	Message string
}

//获取err的内容
func (err GlobalError) Error() string {
	return err.Message
}

const (
	IsEmpty         = 1001 //参数不能为空
	InternalService = 1002 //内部服务错误
)

//IsEmptyError  参数不能为空
func IsEmptyError(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    IsEmpty,
		Message: message,
	}
}

//InternalServiceError  内部服务错误
func InternalServiceError(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusForbidden,
		Code:    InternalService,
		Message: message,
	}
}
