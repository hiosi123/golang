package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Makefile() {
	file, _ := os.Create("file.txt")
	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello"))
	fmt.Println(n, err)
	n, err = io.WriteString(writer, "!")
	fmt.Println(n, err)

	file.Close()

	file, _ = os.Open("file.txt")
	reader := io.Reader(file)
	buffer := make([]byte, 10)
	reader.Read(buffer)
	fmt.Println(string(buffer))
	file.Close()

	//seeker interface - wraps the basic Seek method.
	//trying to read the file using io.Reader by using twice.

	seeker := reader.(io.Seeker)
	seeker.Seek(0, io.SeekStart)
	buffer, err = io.ReadAll(reader)
	reader.Read(buffer)
	fmt.Println(string(buffer))
}
