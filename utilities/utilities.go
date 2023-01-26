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

package utilities

import (
	"bytes"
	"encoding/hex"
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

func SplitSize(size string) (int, int) {
	a := strings.Split(size, "x")
	w, _ := strconv.Atoi(a[0])
	h, _ := strconv.Atoi(a[1])

	return w, h
}

func GenerateImage(width, height int, bgColor string, txtColor string, text string) []byte {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	if bg, _ := hex.DecodeString(bgColor); len(bg) == 3 {
		draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{bg[0], bg[1], bg[2], 255}}, image.Point{}, draw.Src)
	}
	// TODO - default bg color

	var c color.Color = color.RGBA{255, 255, 255, 255}
	if txt, _ := hex.DecodeString(txtColor); len(txt) == 3 {
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
		log.Fatal("Failed to parse font")
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: face,
	}
	d.Dot = fixed.P(10, 10+int(face.Metrics().Ascent.Ceil()))
	d.DrawString(text)
	// TODO - center text in image

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
