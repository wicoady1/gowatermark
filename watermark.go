package gowatermark

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// New is used to initialize WatermarkImage struct
// mainFilePath parameter is main image file path
// imageType is image format, either ImageJPEG or ImagePNG
func New(mainFilePath string, imageType int) (WatermarkImage, error) {
	imageFile, err := os.Open(mainFilePath)
	if err != nil {
		return WatermarkImage{}, fmt.Errorf("Image Load Fail: %+v", err)
	}

	image, err := processImageType(imageFile, imageType)
	if err != nil {
		return WatermarkImage{}, fmt.Errorf("Image Process Fail: %+v", err)
	}

	return WatermarkImage{
		MainImage: image,
	}, nil
}

// AddOverheadImage for append your image onto top of main image
// overheadImage parameter is overhead image file path
// imageType is image format, either ImageJPEG or ImagePNG
func (wm *WatermarkImage) AddOverheadImage(overheadImage string, imageType int) error {
	ohImageFile, err := os.Open(overheadImage)
	if err != nil {
		return fmt.Errorf("Overhead Image Load Fail: %+v", err)
	}

	image, err := processImageType(ohImageFile, imageType)
	if err != nil {
		return fmt.Errorf("Overhead Image Process Fail: %+v", err)
	}

	newOHImage := OverheadImage{
		Image:   image,
		OffsetX: 0,
		OffsetY: 0,
		Alpha:   1,
		Size:    1,
	}

	wm.OverlayImage = append(wm.OverlayImage, newOHImage)

	return nil
}

// OutputImage to process all stored image into a variable of image.Image
func (wm *WatermarkImage) OutputImage() (image.Image, error) {
	//draw image which is used to draw main picture
	r := image.Rectangle{image.Point{0, 0}, wm.MainImage.Bounds().Size()}
	imageResult := image.NewRGBA(r)
	draw.Draw(imageResult, r, wm.MainImage, image.Point{0, 0}, draw.Over)

	for _, v := range wm.OverlayImage {
		//for now assume all input is 0,0
		startPoint := calcCoordinateStartingPoint(wm, v, v.Position, v.OffsetX, v.OffsetY)
		rOH := image.Rectangle{image.Point{0, 0}, v.Image.Bounds().Size()}

		draw.Draw(imageResult, rOH, v.Image, startPoint, draw.Over)
	}

	return imageResult, nil
}

// OutputImageToFile to process all stored image and create a new image file based on result
func (wm *WatermarkImage) OutputImageToFile(path string, filename string, outputType int) error {
	imageResult, err := wm.OutputImage()
	if err != nil {
		return err
	}

	extension := "jpeg"
	if outputType == ImagePNG {
		extension = "png"
	}
	outputFilename := fmt.Sprintf("%s%s.%s", path, filename, extension)

	out, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println(err)
	}

	if outputType == ImagePNG {
		png.Encode(out, imageResult)
	} else {
		jpeg.Encode(out, imageResult, nil)
	}
	return nil
}

// SetPosition adjusts OverheadImage anchor
func (oh *OverheadImage) SetPosition(position int) error {
	if err := validateImagePosition(position); err != nil {
		return fmt.Errorf("SetPosition Fail: %+v", err)
	}
	oh.Position = position

	return nil
}

// SetOffset adjusts x and y offset of OverheadImage based on the anchor
func (oh *OverheadImage) SetOffset(x int, y int) {
	oh.OffsetX = x
	oh.OffsetY = y
}

// SetAlpha sets image transparancy (not functional yet)
func (oh *OverheadImage) SetAlpha(input float64) error {
	if err := validatePercent(input); err != nil {
		return fmt.Errorf("SetOpacity Fail: %+v", err)
	}
	oh.Alpha = input

	return nil
}
