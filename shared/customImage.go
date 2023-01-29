/*
 * Produced: Thu Jan 26 2023
 * Author: Alec M., James A.
 * GitHub: https://github.com/placeholder-app
 * Copyright: (C) 2023 Alec M., James A.
 * License: License GNU Affero General Public License v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package shared

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"github.com/placeholder-app/go-fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type CustomImage struct {
	// Image dimensions in pixels
	Width, Height, BorderWidth int
	// Background and text colors, text string
	BgColor, TxtColor, BorderColor, Text string
	// Font family
	FontFamily string
}

// BuildImage builds an image from the CustomImage struct
//
// i: CustomImage struct
//
// Example: BuildImage(&CustomImage{Width: 100, Height: 100, Text: "Hello, World!"})
func (i *CustomImage) Build() (data []byte, err error) {
	// Build the image base
	img := i.GenerateBase()

	// Draw the text
	fontface := fonts.CalSansSemiBold.GetFontFace(78, 32)
	textData := fonts.CalSansSemiBold.GetTextData(fontface, i.Text)
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{i.GetTxtColor()},
		Face: fontface.Face,
		Dot:  fixed.P((i.Width-textData.Width)/2, (i.Height/2)+(textData.Height/2)),
	}
	drawer.DrawString(i.Text)

	// Encode the image
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("failed to encode image: %v", err)
	} else {
		return buf.Bytes(), nil
	}
}

// Parse a color from a hex string
//
// toParse: hex string to parse
//
// author: James-Elicx
func (i CustomImage) parseColor(toParse string, fallback color.RGBA) color.RGBA {
	if txt, _ := hex.DecodeString(toParse); len(txt) == 3 {
		return color.RGBA{txt[0], txt[1], txt[2], 255}
	}

	return fallback
}

// Convert CustomImage BgColor to a color.RGBA
//
// Example: GetBgColor("ffffff") = color.RGBA{255, 255, 255, 255}
//
// author: James-Elicx
func (i CustomImage) GetBgColor() color.RGBA {
	return i.parseColor(i.BgColor, color.RGBA{171, 171, 171, 255})
}

// Convert CustomImage TxtColor to a color.RGBA
//
// Example: GetTxtColor("ffffff") = color.RGBA{255, 255, 255, 255}
//
// author: James-Elicx
func (i CustomImage) GetTxtColor() color.RGBA {
	return i.parseColor(i.TxtColor, color.RGBA{255, 255, 255, 255})
}

// Convert CustomImage BorderColor to a color.RGBA
//
// Example: GetBorderColor("ffffff") = color.RGBA{255, 255, 255, 255}
//
// author: amattu2
func (i CustomImage) GetBorderColor() color.RGBA {
	return i.parseColor(i.BorderColor, color.RGBA{171, 171, 171, 255})
}

// GenerateBase draws the base of the image including the background and border
//
// img: image.RGBA to draw to
//
// author: amattu2
func (i CustomImage) GenerateBase() *image.RGBA {
	// Create the image base
	img := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))

	// Draw the border
	if i.BorderWidth > 0 {
		borderUniform := image.NewUniform(i.GetBorderColor())
		draw.Draw(img, img.Bounds(), borderUniform, image.Point{}, draw.Src)
	}

	// Draw the background
	backgroundRect := image.Rect(i.BorderWidth, i.BorderWidth, i.Width-i.BorderWidth, i.Height-i.BorderWidth)
	draw.Draw(img, backgroundRect, image.NewUniform(i.GetBgColor()), image.Point{}, draw.Src)

	return img
}
