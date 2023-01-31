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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestImageGenerateNoOptions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetImage(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestAcceptWithOversizedBorder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{
		{Key: "size", Value: "100x100"},
	}
	context.Request = httptest.NewRequest("GET", "/?borderWidth=26&format=gif", nil)

	GetImage(context)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "image/gif", recorder.Header().Get("Content-Type"))
}

func TestAcceptWithUndersizedBorder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{
		{Key: "size", Value: "100x100"},
	}
	context.Request = httptest.NewRequest("GET", "/?borderWidth=-1", nil)

	GetImage(context)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "image/png", recorder.Header().Get("Content-Type"))
}

func TestIncorrectImageSizeFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "size", Value: "fakesize"}}

	GetImage(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestInvalidImageSize1(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "size", Value: "0x0"}}

	GetImage(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestInvalidImageSize2(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "size", Value: "-450x500"}}

	GetImage(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestInvalidImageSize3(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "size", Value: "5555x240"}}

	GetImage(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestValidImageCreateFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = []gin.Param{
		{Key: "size", Value: "315x315"},
	}
	context.Request = httptest.NewRequest("GET", "/?format=bmp", nil)

	GetImage(context)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "image/bmp", recorder.Header().Get("Content-Type"))
}

func TestGetImageFonts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetFonts(context)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
	assert.Equal(t, true, json.Valid(recorder.Body.Bytes()))
	assert.Equal(t, true, strings.Contains(recorder.Body.String(), "status"))
	assert.Equal(t, true, strings.Contains(recorder.Body.String(), "fonts"))
}

func TestGetImageFormats(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	GetFormats(context)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
	assert.Equal(t, true, json.Valid(recorder.Body.Bytes()))
	assert.Equal(t, true, strings.Contains(recorder.Body.String(), "status"))
	assert.Equal(t, true, strings.Contains(recorder.Body.String(), "formats"))
}
