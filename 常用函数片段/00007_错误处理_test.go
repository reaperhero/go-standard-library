package command

import "errors"

const (
	ErrCodeOK = 1000
)

var (
	ErrDbInvalid = errors.New("数据库出错")
)

func GetErrorCode(err error) int32 {
	switch err {
	case ErrDbInvalid:
		return ErrCodeOK + 1
	case nil:
		return ErrCodeOK
	default:
		return 0
	}

}

func GetErrorMap(err error) map[string]interface{} {
	var msg = "OK"
	if err != nil {
		msg = err.Error()
	}

	return map[string]interface{}{
		"errcode": GetErrorCode(err),
		"msg":     msg,
	}
}
