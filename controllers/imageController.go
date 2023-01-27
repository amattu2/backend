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

package controllers

import (
	"fmt"
	"placeholder-app/backend/shared"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {
	if c.Params == nil || !strings.Contains(c.Param("size"), "x") {
		c.AbortWithStatus(400)
		return
	}

	var width, height int = shared.SplitSize(c.Param("size"))
	if width < 30 || width > 4000 || height < 30 || height > 4000 {
		c.AbortWithStatus(400)
		return
	}

	var text string = strings.TrimSpace(c.Query("text"))
	if text == "" || len(text) > 100 {
		text = fmt.Sprintf("%d x %d", width, height)
	}

	img := &shared.CustomImage{
		Width:      width,
		Height:     height,
		Text:       text,
		FontFamily: c.Query("font"),
		BgColor:    c.Param("bgColor"),
		TxtColor:   c.Param("txtColor"),
	}

	if data, err := img.Build(); err != nil {
		c.AbortWithStatus(500)
	} else {
		c.Header("Cache-Control", "public, max-age=86400")
		c.Header("Content-Disposition", "inline; filename=image.png")
		c.Data(200, "image/png", data)
	}
}
