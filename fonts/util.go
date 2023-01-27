package fonts

import (
	"bytes"
	"compress/flate"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// courtesy of https://github.com/google/gxui/blob/master/gxfont/gxfont.go
func inflateFontFile(src []byte) []byte {
	r := bytes.NewReader(src)
	b, err := ioutil.ReadAll(flate.NewReader(r))
	if err != nil {
		panic(err)
	}
	return b
}

type Font struct {
	Font *truetype.Font
}

// parse an inflated font file to a Font struct
func parseFont(src []byte) *Font {
	f, err := truetype.Parse(inflateFontFile(src))
	if err != nil {
		panic(err)
	}
	return (&Font{Font: f})
}

var (
	CalSansSemibold *Font = parseFont(cal_sans_semibold)
)

type TextInfo struct {
	Face   font.Face
	Width  int
	Height int
}

// retrieve the font face and calculate text size
func (f *Font) GetTextData(text string, size float64) *TextInfo {
	// get the font face
	face := truetype.NewFace(f.Font, &truetype.Options{
		Size:    size, // TODO - make this dynamic based on W/H
		DPI:     72,
		Hinting: font.HintingFull,
	})

	var textWidth int = 0
	var heightMin, heightMax fixed.Int26_6 = 0, 0

	for _, r := range text {
		bounds, advance, _ := face.GlyphBounds(r)
		textWidth += advance.Floor()

		// to get the total height we need to find to the lowest point and the highest point for the string.
		if bounds.Min.Y < heightMin {
			heightMin = bounds.Min.Y
		}
		if bounds.Max.Y > heightMax {
			heightMax = bounds.Max.Y
		}
	}
	// then we can subtract the two to get the value we can use when rendering the text.
	textHeight := (0 - heightMin - heightMax).Ceil()

	return &TextInfo{
		Face:   face,
		Width:  textWidth,
		Height: textHeight,
	}
}
