package main

import (
	"fmt"
	"math/bits"
	"os"
)

func main() {

	// f := excelize.NewFile()

	// f.SetCellValue("Sheet1", "B2", 100)
	// f.SetCellValue("Sheet1", "A1", 50)

	// now := time.Now()

	// f.SetCellValue("Sheet1", "A4", now.Format(time.ANSIC))

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		str := fmt.Sprintf("%c%d", 'A'+i, j+1)
	// 		fmt.Println(str)
	// 		f.SetCellValue("Sheet1", str, i*j)
	// 	}
	// }

	// if err := f.SaveAs("simple.xls"); err != nil {
	// 	log.Fatal(err)
	// }
	// f, err := excelize.OpenFile("./simple.xls")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // f.setcell
	// val := f.GetCellValue("Sheet1", "C5")
	// fmt.Println(val)
	// types := f.GetCellStyle("Sheet1", "C5")
	// fmt.Println(types)

	// file := Read("simple.xls")
	// fmt.Printf("%s", file)

	// fmt.Println(bsr(16))
	// fmt.Println(isPowerOfTwo(10))
	// f, _ := os.Create("test")
	// for i := 0; i < 300000; i++ {
	// 	n, err := f.WriteString("abcdefghijklmnopqrst")
	// 	if n != 20 || err != nil {
	// 		break
	// 	}
	// }
	// f.Close()
}

func Read(path string) (data []byte) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	buf := make([]byte, 1024*4*8)
	n, err1 := f.Read(buf)
	if err1 != nil {
		fmt.Println(err1)
		return buf[:n]
	}

	return buf[:n]
}

func bsr(x int) int {
	return bits.Len(uint(x)) - 1
}

func isPowerOfTwo(x int) bool {
	fmt.Printf("%o,%o,%o\n", x, -x, (x & (-x)))
	return (x & (-x)) == x
}
