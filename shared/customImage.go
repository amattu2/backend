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
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"placeholder-app/backend/fonts"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type CustomImage struct {
	Width, Height           int
	BgColor, TxtColor, Text string
	Data                    []byte
}

// BuildImage builds an image from the CustomImage struct
//
// i: CustomImage struct
//
// Example: BuildImage(&CustomImage{Width: 100, Height: 100, Text: "Hello, World!"})
func (i *CustomImage) Build() (err error) {
	img := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))

	text := fonts.CalSansSemibold.GetTextData(i.Text, 32)

	// Add content to image
	draw.Draw(img, img.Bounds(), &image.Uniform{i.GetBgColor()}, image.Point{}, draw.Src)
	drawer := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{i.GetTxtColor()},
		Face: text.Face,
		Dot:  fixed.P((i.Width-text.Width)/2, (i.Height/2)+(text.Height/2)),
	}
	drawer.DrawString(i.Text)

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	} else {
		i.Data = buf.Bytes()
	}

	return nil
}
