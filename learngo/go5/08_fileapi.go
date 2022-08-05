package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	fp, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintln(fp, "hello world")
	defer fp.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		fp.WriteString(buf)
	}

}
func ReadFile(path string) (res string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	buf := make([]byte, 1024*2)
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		fmt.Println(err1.Error())
		return
	}
	res = string(buf[:n])
	fmt.Println("buf = ", string(buf[:n]))
	return
}
func ReadLine(path string) (res []string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		} else if err == io.EOF {
			break
		}
		res = append(res, string(buf))
	}
	return
}

func copy(dst, src string) (err error) {
	fs, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer fs.Close()
	buf := make([]byte, 8*1024)
	n, err := fs.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if n > 8*1024 {
		fmt.Println("The file is too large")
		err = errors.New("The file is too large")
		return
	}

	fd, err := os.Create(dst)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer fd.Close()
	n, err = fd.Write(buf[:n])
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return
}
func main() {
	// filename := "test.txt"
	// fp, err := os.Create(filename)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Fprintln(fp, "hello world")

	// // os.Stdout.Close()
	// // fmt.Println("hello stdout") //无法输出

	path := "./demo.txt"
	// WriteFile(path)

	// ReadFile(path)

	res := ReadLine(path)

	for i, line := range res {
		fmt.Printf("%d:%s", i, line)
	}

	srcpath := "./demo.txt"
	dstpath := "./democopy.txt"
	copy(dstpath, srcpath)

}
