package metanit

import (
	"fmt"
	"testing"
)

type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}
func closeStream(stream Stream) {
	stream.close()
}

type File struct {
	text string
}
type Folder struct{}

func (f *File) read() string {
	return f.text
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Is written", message)
}
func (f *Folder) close() {
	fmt.Println("Folder is closed")
}
func (f *File) close() {
	fmt.Println("File is closed")
}

func TestCorespondanse(t *testing.T) {
	myFile := &File{}
	myFolder := &Folder{}

	writeToStream(myFile, "Annna")
	closeStream(myFile)
	myFolder.close()
}


