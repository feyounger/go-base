package hello

import "fmt"
import "os"

func init() {
	fmt.Println("imp-init() come here.")
}

func Print() {
	fmt.Println("Hello!")
}

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("D:\\GoProject\\go-base\\text.txt")
	defer f.Close()
	for{
		n, _ := f.Read(buf)
		if n == 0{
			break
		}
		os.Stdout.Write(buf[:n])
	}
}