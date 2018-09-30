package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type logProcess struct {
	read  Reader
	write Writer
	wc    chan []byte
	rc    chan []byte
}

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan []byte)
}

type ReaderFromFile struct {
	path string
}

type WriteInfluxDB struct {
	influxDBDsn string
}

func (w *WriteInfluxDB) Write(wc chan []byte) {
	fmt.Println(<-wc)
}

func (r *ReaderFromFile) Read(rc chan []byte) {

	// 打开文件
	f, err := os.OpenFile(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}

	//从文件末尾开始逐行 读取文件内容
	rd := bufio.NewReader(f)
    line , err := rd.ReadBytes(\n)

	line := "message"
	rc <- line
}

func (l *logProcess) Process() {
	//解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

func main() {
	lp := &logProcess{
		read:  &ReaderFromFile{path: "test.log"},
		write: &WriteInfluxDB{influxDBDsn: "username&password.."},
		rc:    make(chan string),
		wc:    make(chan string),
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1)

}
