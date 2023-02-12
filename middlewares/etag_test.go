/*
 * Produced: Sun Feb 12 2023
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

package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"placeholder-app/backend/middlewares"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// Basically test against a bunch of different inputs and make sure the E Tag is correct
// Also assure it returns a 304 if the E Tag matches

func TestEmptyETag(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest("GET", "/", nil)

	middlewares.ETag()(context)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "W/\"da39a3ee5e6b4b0d3255bfef95601890afd80709\"", recorder.Header().Get("ETag"))
}

func TestPathETagGeneration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// hash => request params
	tests := map[string][]gin.Param{
		"216ef3494d19b18b406008b212082526c237d3f1": {
			{Key: "size", Value: "315x315"},
		},
		"c46b8c07757883fe5cb527398fe4d31774f30a05": {
			{Key: "size", Value: "400x400"},
			{Key: "bgColor", Value: "3B3B3B"},
		},
		"cf079f569bef41abbfe8b343656607bd82a68684": {
			{Key: "size", Value: "200x200"},
			{Key: "bgColor", Value: "F2F2F2"},
			{Key: "txtColor", Value: "3B3B3B"},
		},
	}

	for hash, params := range tests {
		// Test Without If-None-Match
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/", nil)
		context.Params = params
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "W/\""+hash+"\"", recorder.Header().Get("ETag"))

		// Test HTTP If-None-Match
		recorder = httptest.NewRecorder()
		context, _ = gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/", nil)
		context.Params = params
		context.Request.Header.Set("If-None-Match", hash)
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusNotModified, recorder.Code)
	}
}

func TestQueryETagGeneration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// hash => query string
	tests := map[int]map[string]string{
		0: {
			"hash":  "5dead348ca8321104b06f87fdcc6b9872781aed9",
			"query": "?borderWidth=24&borderColor=3B3B3B",
		},
		1: {
			"hash":  "5dead348ca8321104b06f87fdcc6b9872781aed9",
			"query": "?borderColor=3B3B3B&borderWidth=24",
		},
		2: {
			"hash":  "c588cc537aa3c78e875bcc80baec5cfb88cddf60",
			"query": "?text=Hello%2C%20World!&borderWidth=24&borderColor=3B3B3B",
		},
		3: {
			"hash":  "4d8c3c24a9ae9f84f89317c798150fe5b450ba75",
			"query": "?font=Sans-Serif&format=bmp&text=Hello%2C%20World!&borderWidth=24&borderColor=3B3B3B",
		},
	}

	for _, tst := range tests {
		// Test Without If-None-Match
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/"+tst["query"], nil)
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "W/\""+tst["hash"]+"\"", recorder.Header().Get("ETag"))

		// Test HTTP If-None-Match
		recorder = httptest.NewRecorder()
		context, _ = gin.CreateTestContext(recorder)
		context.Request = httptest.NewRequest("GET", "/"+tst["query"], nil)
		context.Request.Header.Set("If-None-Match", tst["hash"])
		middlewares.ETag()(context)

		assert.Equal(t, http.StatusNotModified, recorder.Code)
	}
}
