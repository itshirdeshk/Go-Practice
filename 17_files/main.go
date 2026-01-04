package main

import (
	"bufio"
	"os"
)

func main() {
	// f, err := os.Open("newFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("FileName: ", fileInfo.Name())
	// fmt.Println("File Size: ", fileInfo.Size())

	// defer f.Close()

	// buf := make([]byte, 10)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := range buf {
	// 	fmt.Println(d, string(buf[i]))
	// }

	// data, err := os.ReadFile("newFile.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))
	//
	// dir, err := os.Open("..")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(0)
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(fileInfo)
	// for i, file := range fileInfo {
	// 	fmt.Println(i, file.Name())
	// }
	//

	// file, err := os.Create("newFile2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// file.WriteString("Hello Go Again")
	// file.Write([]byte("Hello Go "))
	//
	sourceFile, err := os.Open("newFile.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	
	destFile, err := os.Create("newFile2.txt")
	if err != nil {
		panic(err)
	}
	
	defer destFile.Close()
	
	readBuf := bufio.NewReader(sourceFile)
	writeBuf := bufio.NewWriter(destFile)
	
	for {
		b, err := readBuf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writeBuf.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}
	writeBuf.Flush()
}
