package screenshot

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/go-vgo/robotgo"
)

func TakeScreenshot() (*bytes.Buffer, error) {
	fmt.Println("Taking screenshot")
	bitmap := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(bitmap)

	img := robotgo.ToImage(bitmap)
	imageBuffer := new(bytes.Buffer)
	err := png.Encode(imageBuffer, img)
	if err != nil {
		return nil, err
	}
	return imageBuffer, nil

}
