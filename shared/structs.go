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
	"encoding/hex"
	"image/color"
)

type CustomImage struct {
	Width, Height           int
	BgColor, TxtColor, Text string
	Data                    []byte
}

// parse a color, or default to the fallback color
func (i CustomImage) parseColor(toParse string, fallback color.RGBA) color.RGBA {
	var c color.RGBA = fallback
	if txt, _ := hex.DecodeString(toParse); len(txt) == 3 {
		c = color.RGBA{txt[0], txt[1], txt[2], 255}
	}

	return c
}

// get the background color
func (i CustomImage) GetBgColor() color.RGBA {
	return i.parseColor(i.BgColor, color.RGBA{171, 171, 171, 255})
}

// get the text color
func (i CustomImage) GetTxtColor() color.RGBA {
	return i.parseColor(i.TxtColor, color.RGBA{255, 255, 255, 255})
}
