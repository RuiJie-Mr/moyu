package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/eiannone/keyboard"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type log struct {
	Name string
	Line int
}

const (
	quit = 'm'
)

// 获取指定文件阅读日志
func ChangeLog(logPath, fileName string, nowLine int) {
	f, err := ioutil.ReadFile(logPath)
	if err != nil {
		fmt.Println("read log fail", err.Error())
	}
	var logObj []log
	err = json.Unmarshal(f, &logObj)
	if err != nil {
		fmt.Println("解析日志错误 ", err)
	}
	for k, v := range logObj {
		if v.Name == fileName {
			logObj[k].Line = nowLine
		}
	}
	bytedata, _ := json.Marshal(logObj)
	err = WriteFile(logPath, bytedata, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// 获取指定文件阅读日志
func ReadLog(logPath, fileName string) (line int) {
	f, err := ioutil.ReadFile(logPath)
	if err != nil {
		fmt.Println("read log fail", err.Error())
	}
	var logObj []log
	err = json.Unmarshal(f, &logObj)
	if err != nil {
		fmt.Println("解析日志错误 ", err)
	}
	var hasLog bool
	for _, v := range logObj {
		if v.Name == fileName {
			line = v.Line
			hasLog = true
		}
	}
	if hasLog {
		return line
	}
	var logThis = log{
		Name: fileName,
		Line: 0,
	}

	logObj = append(logObj, logThis)
	bytedata, _ := json.Marshal(logObj)
	err = WriteFile(logPath, bytedata, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

func validUTF8(buf []byte) bool {
	nBytes := 0
	for i := 0; i < len(buf); i++ {
		if nBytes == 0 {
			if (buf[i] & 0x80) != 0 {
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1
					nBytes++
				}
				if nBytes < 2 || nBytes > 6 {
					return false
				}
				nBytes--
			}
		} else {
			if buf[i]&0xc0 != 0x80 {
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Read2(path string, goLine int, logPath string) {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("请输入正确文件路径:", path)
		return
	}
	defer fi.Close()
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	r := bufio.NewReader(fi)
	var line = 1
	for {
		str, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("读取文件错误:", path)
			return
		}
		if err == io.EOF {
			fmt.Println("已读取到文章末尾")
			return
		}
		if line < goLine {
			line++
			continue
		}
		str = strings.Replace(str, "\n", "", -1)
		r := []rune(str)
		l := 120
		n := 1
		for i := 0; i < len(r); i += l {
			end := i + l
			if end > len(r) {
				end = len(r)
			}
			sub := r[i:end]
			s := fmt.Sprintf("%d.%d : %s", line, n, string(sub))
			if !validUTF8([]byte(s)) {
				fmt.Println(enc.ConvertString(s))
			} else {
				fmt.Println(s)
			}
			n++
		}
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if char == quit || char == 46 || char == 183 || char == 96 {
			defer ChangeLog(logPath, path, line)
			for i := 0; i < 1000; i++ {
				fmt.Printf("\n")
			}
			return
		}
		line++
	}
}
