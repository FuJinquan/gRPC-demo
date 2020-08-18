package ErrorCode

var (
	ErrSuccess      = StandardError{0, "成功"}
	ErrUnrecognized = StandardError{-1, "未知错误"}
	ErrTitleIsNil   = StandardError{1000, "标题为空错误"}
	ErrRecordIsNil  = StandardError{1001, "记录为空错误"}
	ErrDbIsNil      = StandardError{1002, "数据库为空错误"}
	ErrCreateTable  = StandardError{2000, "创建数据表错误"}
	ErrCreateDb     = StandardError{2001, "创建数据库失败"}
	ErrCreateRecord = StandardError{2002, "创建记录错误"}
	ErrSearch       = StandardError{3000, "查询数据库错误"}
	ErrDelete       = StandardError{4000, "删除记录错误"}
	ErrUpdate       = StandardError{5000, "记录更新错误"}
)

type StandardError struct {
	ErrorCode int32
	ErrorMsg  string
}
