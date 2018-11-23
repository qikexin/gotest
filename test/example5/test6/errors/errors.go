package errors

import (
	"strings"
	"bytes"
	"fmt"
)

//ErrorType 代表错误类型
type ErrorType string
//错误类型常量
const (
	//下载器错误
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	//分析器错误
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	//条目处理管道错误
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	//调度器错误
	ERROR_TYPE_SCHEDULER ErrorType = "scheduler error"

)

//CrawlerError代表爬虫错误的接口类型
type CrawlerError interface {
	Type() ErrorType
	Error() string
}
type myCrawlerError struct {
	//errType代表错误的类型
	errType ErrorType
	//errMsg代表错误的提示信息
	errMsg string
	//fullErrMsg代表完整的错误提示信息
	fullErrMsg string
}

func NewCrawlerError(errType ErrorType,errMsg string) CrawlerError{
	return &myCrawlerError{
		errType: errType,
		errMsg: strings.TrimSpace(errMsg),
	}
}

func NewCrawlerErrorBy(errType ErrorType,err error) CrawlerError{
	return NewCrawlerError(errType,err.Error())
}

func (ce *myCrawlerError) Type() ErrorType  {
	return ce.errType
}
func (ce *myCrawlerError) Error() string  {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if ce.errType != ""{
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s",buffer.String())
	return
}