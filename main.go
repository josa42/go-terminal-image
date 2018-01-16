package image

import (
	b64 "encoding/base64"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

// Size :
type Size struct {
	Width  int
	Height int
	Unit   string
}

// Create :
func Create(href string) string {
	return CreateWithSize(href, Size{})
}

// CreateWithSize :
func CreateWithSize(href string, size Size) string {
	fileName := b64.StdEncoding.EncodeToString([]byte("filename.jpg"))
	data, _ := ioutil.ReadFile(href)
	image := b64.StdEncoding.EncodeToString(data)
	fileSize := binary.Size(data)

	sizeParams := ""
	if size.Width > 0 {
		sizeParams += fmt.Sprintf(";width=%d%s", size.Width, size.Unit)
	}
	if size.Height > 0 {
		sizeParams += fmt.Sprintf(";height=%d%s", size.Height, size.Unit)
	}

	return fmt.Sprintf("\033]1337;File=name=%s;size=%d%s;inline=1:%s\a\n", fileName, fileSize, sizeParams, image)
}

// Print :
func Print(href string) {
	link := Create(href)
	fmt.Print(link)
}

// func main() {
// 	// fmt.Print(Create("/Users/josa/Desktop/download.png"))
// 	fmt.Print(CreateWithSize("/Users/josa/Desktop/download.png", Size{Width: 10}))
// 	fmt.Print(CreateWithSize("/Users/josa/Desktop/download.png", Size{Height: 2}))
// 	fmt.Print(CreateWithSize("/Users/josa/Desktop/download.png", Size{Height: 2}))
// 	fmt.Print(CreateWithSize("/Users/josa/Desktop/download.png", Size{Height: 2}))
// }
