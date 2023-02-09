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

package main

import (
	"fmt"
	"log"
	"placeholder-app/backend/middlewares"
	"placeholder-app/backend/routes"
	"placeholder-app/backend/shared"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	cert       = shared.GetEnv("SSLCERT", "")
	key        = shared.GetEnv("SSLKEY", "")
	address    = shared.GetEnv("ADDR", "")
	port       = shared.GetEnv("PORT", "8080")
	requestMax = shared.GetEnv("REQUESTMAX", "10")
	env        = shared.GetEnv("APP_ENV", "prod")
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.Default()

	engine.SetTrustedProxies(nil)
	engine.Use(middlewares.RateLimit(time.Minute, uint(shared.CoerceInt(requestMax))))
	engine.Use(middlewares.Cors())

	routes.InitRouter(engine)

	var err error
	if cert != "" && key != "" {
		err = engine.RunTLS(fmt.Sprintf("%s:%s", address, port), cert, key)
	} else {
		err = engine.Run(fmt.Sprintf("%s:%s", address, port))
	}

	log.Fatalf("Failed to start server: %s", err)
}
