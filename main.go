package image

import (
	b64 "encoding/base64"
	"encoding/binary"
	"fmt"
	"image"
	"io/ioutil"
	"os"
)

// Size :
type Size struct {
	Width     int
	Height    int
	MaxWidth  int
	MaxHeight int
	Unit      string
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
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

	if size.MaxWidth > 0 || size.MaxHeight > 0 {
		imageWidth, imageHeight := getImageDimension(href)

		if size.MaxWidth > 0 && imageWidth > size.MaxWidth {
			sizeParams += fmt.Sprintf(";width=%dpx", size.MaxWidth)
		}

		if size.MaxHeight > 0 && imageHeight > size.MaxHeight {
			sizeParams += fmt.Sprintf(";height=%dpx", size.MaxHeight)
		}

	} else {
		if size.Width > 0 {
			sizeParams += fmt.Sprintf(";width=%d%s", size.Width, size.Unit)
		}
		if size.Height > 0 {
			sizeParams += fmt.Sprintf(";height=%d%s", size.Height, size.Unit)
		}
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
