package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

func chackErr(err error)  {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\n go\n")
	err := ioutil.WriteFile("text.txt",d1,0644)
	chackErr(err)

	f, err := os.Create("tmp/dat2")
	chackErr(err)

	defer f.Close()

	d2 := []byte{115,111,109,101,10} // some
	n2,err := f.Write(d2)
	chackErr(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writer\n")
	fmt.Printf("wrote %d bytes\n",n3)
	f.Sync()

	w := bufio.NewWriter(f)
	n4 ,err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes \n", n4)

	w.Flush()
}
