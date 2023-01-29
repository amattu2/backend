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
	"placeholder-app/backend/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	// Root Routes
	engine.GET("/status", controllers.GetStatus)
	engine.GET("/fonts", controllers.GetFonts)

	// Image Routes
	imageRoutes := engine.Group("/image/:size")
	imageRoutes.GET("/", controllers.GetImage)
	imageRoutes.GET("/:bgColor", controllers.GetImage)
	imageRoutes.GET("/:bgColor/:txtColor", controllers.GetImage)
}
