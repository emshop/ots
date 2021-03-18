package xerr

import "errors"

//ErrNOTEXISTS 数据不存在
var ErrNOTEXISTS = errors.New("NOT_EXISTS")

//ErrOrderTimeout 订单超时
var ErrOrderTimeout = errors.New("ORDER_TIME_OUT")
