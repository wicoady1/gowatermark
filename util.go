package gowatermark

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// processImageType converts file into image variable
// based on image's format (JPEG or PNG)
func processImageType(file *os.File, imageType int) (image.Image, error) {
	var result image.Image
	var err error

	if imageType == ImageJPEG {
		result, err = jpeg.Decode(file)
		if err != nil {
			return nil, err
		}
	} else if imageType == ImagePNG {
		result, err = png.Decode(file)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// validateImagePosition checks position flags and return error if flag is invalid
func validateImagePosition(position int) error {
	if position < TopLeftCorner || position > Free {
		return fmt.Errorf("Invalid Position")
	}
	return nil
}

// validatePercent checks input must be around 0 to 100
func validatePercent(input float64) error {
	if input < float64(0) || input > float64(100) {
		return fmt.Errorf("Invalid Number")
	}
	return nil
}

// calcCoordinateStartingPoint calculates OverheadImage final anchor position
// based on WatermarkImage size, OverheadImage positions and offset
func calcCoordinateStartingPoint(wm *WatermarkImage, oh OverheadImage, position int, offsetX int, offsetY int) image.Point {

	return image.Point{0, 0}
}
