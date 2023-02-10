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

package routes

import (
	"net/http"
	"placeholder-app/backend/controllers"
	"placeholder-app/backend/middlewares"
	"placeholder-app/backend/shared"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	/**
	 * Root Route
	 *
	 * @GET /
	 */
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, shared.GetEnv("APP_URL", "https://placeholder.app"))
	})

	/**
	 * Utility Routes
	 *
	 * @GET /status
	 */
	engine.GET("/status", controllers.GetStatus)

	/**
	 * Image Routes
	 *
	 * Middleware: ETag
	 *
	 * @GET /image/fonts
	 * @GET /image/formats
	 * @GET /image/{size}
	 * @GET /image/{size}/{bgColor}
	 * @GET /image/{size}/{bgColor}/{txtColor}
	 */
	imageRoutes := engine.Group("/image", middlewares.ETag())
	imageRoutes.GET("/fonts", controllers.GetFonts)
	imageRoutes.GET("/formats", controllers.GetFormats)
	imageRoutes.GET("/:size/", controllers.GetImage)
	imageRoutes.GET("/:size/:bgColor", controllers.GetImage)
	imageRoutes.GET("/:size/:bgColor/:txtColor", controllers.GetImage)
}
