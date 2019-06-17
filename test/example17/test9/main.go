/*
Author: lipengwei
Date: 2019/5/15
Description: 
*/
package main

import (
	"sync"
	"bytes"
	"fmt"
	"io"
)
//Deprecated: use buffPool instead
var bufPool sync.Pool

type Buffer interface {
	//用于获取数据块之间的定界符
	Delimiter()
	//Write 用于写一个数据块
	Write(contents string)(err error)
	//Read用于读一个数据块
	Read()(contents string,err error)
	//Free用于释放当前缓冲区
	Free()
}
type myBuffer struct {
	buf bytes.Buffer
	delimiter byte
}

func (b *myBuffer)Delimiter() byte  {
	return b.delimiter
}
func (b *myBuffer)Write(contents string)(err error)  {
	if _,err = b.buf.WriteString(contents);err != nil {
		return
	}
	return b.buf.WriteByte(b.delimiter)
}
func (b *myBuffer) Read()(contents string,err error)  {
	return b.buf.ReadString(b.delimiter)
}
func (b *myBuffer) Free()  {
	bufPool.Put(b)
}
var delimiter = byte('\n')

func init()  {
	bufPool = sync.Pool{
		New: func() interface{} {
			return &myBuffer{delimiter: delimiter}
		},
	}
}
func GetBuffer() Buffer  {
	return bufPool.Get().(Buffer)
}
func main()  {
	buf :=GetBuffer()
	defer buf.Free()
	buf.Write("A Pool is a set of temporary objects that" +
		"may be individually saved and retrieved.")
	buf.Write("A Pool is safe for use by multiple goroutines simultaneously.")
	buf.Write("A Pool must not be copied after first use.")

	fmt.Println("The data blocks in buffer:")
	for {
		block, err := buf.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("unexpected error: %s", err))
		}
		fmt.Print(block)
	}
}