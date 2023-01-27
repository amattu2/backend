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
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"strconv"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/google/gxui/gxfont"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// RoundTo rounds a number to the nearest multiple of a positive number
// x: non-negative number to round
// to: non-negative number to round by
// Example: RoundTo(102, 5) = 100
func RoundTo(x int, to int) int {
	return (x + 2) / to * to
}

// SplitSize splits a string into two integers
// size: string to split, delimited by "x"
// Example: SplitSize("100x100") = 100, 100
func SplitSize(size string) (int, int) {
	a := strings.Split(size, "x")
	w, _ := strconv.Atoi(a[0])
	h, _ := strconv.Atoi(a[1])

	return RoundTo(w, 5), RoundTo(h, 5)
}

func (i *CustomImage) Build() (err error) {
	img := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))

	if bg, _ := hex.DecodeString(i.BgColor); len(bg) == 3 {
		draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{bg[0], bg[1], bg[2], 255}}, image.Point{}, draw.Src)
	}
	// TODO - default bg color

	var c color.Color = color.RGBA{255, 255, 255, 255}
	if txt, _ := hex.DecodeString(i.TxtColor); len(txt) == 3 {
		c = color.RGBA{txt[0], txt[1], txt[2], 255}
	}

	var face font.Face
	if f, _ := truetype.Parse(gxfont.Default); f != nil {
		face = truetype.NewFace(f, &truetype.Options{
			Size:    32, // TODO - make this dynamic
			DPI:     72,
			Hinting: font.HintingFull,
		})
	} else {
		return fmt.Errorf("failed to parse font")
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: face,
	}
	d.Dot = fixed.P(10, 10+int(face.Metrics().Ascent.Ceil()))
	d.DrawString(i.Text)
	// TODO - center text in image

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		log.Fatal(err)
	}

	i.Data = buf.Bytes()
	return nil
}
