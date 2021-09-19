package calhounAlgorithms

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var nums = []struct {
	n        int
	expected string
}{
	{10, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz"},
	{5, "1, 2, Fizz, 4, Buzz"},
	{10, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, 10"},
	{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz"},
	{17, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17"},
	{20, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz"},
}

func TestFizzBuzz(t *testing.T) {

	for _, tt := range nums {

		t.Run(fmt.Sprintf("N=%d", tt.n), func(t *testing.T) {
			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			osStdout := os.Stdout
			os.Stdout = writer
			defer func(){
				os.Stdout = osStdout
			}()

			FizzBuzz(tt.n)
			writer.Close()

			var buf bytes.Buffer
			io.Copy(&buf, testStdout)
			got := buf.String()
			if got != tt.expected{
				t.Fatalf("FizzBuzz(%d) = %q; want %q",tt.n,got,tt.expected)
			}
		})
	}
}
