package gowatermark

import (
	"image"
)

// WatermarkImage is main struct
type WatermarkImage struct {
	MainImage    image.Image
	TypeImage    int
	OverlayImage []OverheadImage
}

// OverheadImage is struct for saving overehead image properties
type OverheadImage struct {
	Image     image.Image
	TypeImage int
	OffsetX   int
	OffsetY   int
	Alpha     float64
	Size      float64
	Position  int
}

// Position Anchors contants
const (
	TopLeftCorner = iota + 1
	TopRightCorner
	BottomLeftCorner
	BottomRightCorner
	Fill
	Stretch
	Free
)

// Image Type contants
const (
	ImageJPEG int = iota + 1
	ImagePNG
)
