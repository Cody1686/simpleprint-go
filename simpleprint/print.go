package simpleprint

import (
	"fmt"
	"io"
)

// Println 包装了fmt.Println函数，可以直接使用
func println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

// Printf 包装了fmt.Printf函数，可以直接使用
func printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

// Print 包装了fmt.Print函数，可以直接使用
func print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

// Sprintf 包装了fmt.Sprintf函数，可以直接使用
func sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Fprintf 包装了fmt.Fprintf函数，可以直接使用
func fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
} 
