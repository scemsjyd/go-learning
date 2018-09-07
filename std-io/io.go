package std_io

import (
	"strings"
	"fmt"
	"os"
	"bufio"
	"bytes"
	"io"
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
	otherFile, err2 := os.Create("writeAt2.txt")
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

func ByteWriter() {
	var ch byte
	fmt.Scanf("%c\n", &ch)

	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功！准备读取该字节……")
		newCh, _ := buffer.ReadByte()
		fmt.Printf("读取的字节：%c\n", newCh)
	} else {
		fmt.Println("写入错误")
	}
}

func ByteScanner() {
	buffer := bytes.NewBuffer([]byte{'a', 'b'})
	buffer.ReadByte()
	err := buffer.UnreadByte()

	err = buffer.UnreadByte()

	if err != nil {
		panic(err)
	}
}

func LimitReader() {
	content := "This Is LimitReader Example"
	reader := strings.NewReader(content)
	limitReader := io.LimitedReader{R: reader, N: 8}
	for limitReader.N > 0 {
		tmp := make([]byte, 2)
		limitReader.Read(tmp)
		fmt.Printf("%s", tmp)
	}
}

type Person struct {
	Name string
	Age  int
	Sex  int
}

func Stringer() {
	p := &Person{"polaris", 28, 0}
	fmt.Println(p)
}

func Scanner() {
	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// 将文件 offset 设置到文件开头
	file.Seek(0, os.SEEK_SET)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}