/*
 * Produced: Wed Feb 08 2023
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

package middlewares

import (
	"net/http"
	"placeholder-app/backend/shared"

	"github.com/gin-gonic/gin"
)

// GenerateEtag generates an ETag based on the query string
//
// c: gin context to generate ETag from
func generateEtag(c *gin.Context) string {
	from := []string{
		c.Param("size"),
		c.Param("bgColor"),
		c.Param("txtColor"),
		c.Query("text"),
		c.Query("borderWidth"),
		c.Query("borderColor"),
		c.Query("font"),
		c.Query("format"),
	}
	str := ""

	for _, v := range from {
		str += v
	}

	return shared.GenerateHash(str)
}

func ETag() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sum = generateEtag(c)

		if c.Request.Header.Get("If-None-Match") == sum {
			c.AbortWithStatus(http.StatusNotModified)
		}

		c.Writer.Header().Set("ETag", sum)

		c.Next()
	}
}
