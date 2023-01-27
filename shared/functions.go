/*
 * Produced: Thu Jan 26 2023
 * Author: Alec M.
 * GitHub: https://amattu.com/links/github
 * Copyright: (C) 2023 Alec M.
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
	"strconv"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// RoundTo rounds a number to the nearest multiple of a positive number
//
// x: non-negative number to round
//
// to: non-negative number to round by
//
// Example: RoundTo(102, 5) = 100
func RoundTo(x int, to int) int {
	return (x + 2) / to * to
}

// SplitSize splits a string into two integers
//
// size: string to split, delimited by "x"
//
// Example: SplitSize("100x100") = 100, 100
func SplitSize(size string) (int, int) {
	a := strings.Split(size, "x")
	w, _ := strconv.Atoi(a[0])
	h, _ := strconv.Atoi(a[1])

	return RoundTo(w, 5), RoundTo(h, 5)
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
