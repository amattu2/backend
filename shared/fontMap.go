/*
 * Produced: Sun Jan 29 2023
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
	"github.com/placeholder-app/go-fonts"
	"github.com/placeholder-app/go-fonts/util"
)

// FontMap is a map of font names to font structs
//
// font name => *util.Font
var FontMap = map[string]util.Font{
	"CalSansSemiBold": *fonts.CalSansSemiBold,
}

// Get Key Names from FontList
//
// Example: GetFontList() = ["Arial", "TimesNewRoman", "ComicSansMS"]
func GetFontList() []string {
	var fontList []string

	for key := range FontMap {
		fontList = append(fontList, key)
	}

	return fontList
}

// Get Font from FontMap by Name
//
// name: name of font to get
//
// Example: GetFont("Arial") = util.Font
func GetFontStruct(name string) util.Font {
	if font, ok := FontMap[name]; ok {
		return font
	}

	return FontMap["CalSansSemiBold"]
}
