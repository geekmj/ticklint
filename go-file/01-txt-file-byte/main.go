package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Text file to byte")
	t, e := TextFileToBytes("sample.txt")
	if e != nil {
		fmt.Printf("Error %s", e)
	}
	fmt.Printf("File to text %s", t)

	t, e = QuickRead("sample.txt")
	if e != nil {
		fmt.Printf("Error %s", e)
	}
	fmt.Printf("Quick read File to text %s", t)

}
func QuickRead(f string) (bs []byte, e error) {
	bs, e = os.ReadFile(f)
	return
}
func TextFileToBytes(f string) (bs []byte, e error) {

	file, e := os.Open(f)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer file.Close()

	// Get the file size
	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e)
		return
	}

	// Read the file into a byte slice
	bs = make([]byte, stat.Size())
	_, e = bufio.NewReader(file).Read(bs)
	if e != nil && e != io.EOF {
		fmt.Println(e)
		return
	}
	fmt.Println(bs)
	return
}
