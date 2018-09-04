package std_io

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

// 从下标为2的地方读取指定大小的字节
// 打印： 234567, 6
func ReadAt() {
	reader := strings.NewReader("0123456789")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}

// 文件先写入0123456789，然后在下标为4的地方开始写入insert，会替换原来位置的字节
// 打印：插入的字节大小6
func WriterAt() {
	file, e := os.Create("writeAt.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	file.WriteString("0123456789 ")
	n, err := file.WriteAt([]byte("insert"), 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func ReaderFrom() {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

func WriterTo() {
	file, err := os.Open("writeAt.txt")
	otherFile, err2 := os.Create("write2.txt")
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	defer file.Close()
	defer otherFile.Close()
	reader := bufio.NewReader(file)
	reader.WriteTo(otherFile)
}
