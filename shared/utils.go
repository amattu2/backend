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

package shared

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strconv"
	"strings"
)

// RoundTo rounds a number to the nearest multiple of a positive number
//
// x: non-negative number to round
//
// to: non-negative number to round by
//
// Example: RoundTo(102, 5) = 100
func RoundTo(x int, to int) int {
	return (x + 2) / to * to
}

// SplitSize splits a string into two integers
//
// size: string to split, delimited by "x"
//
// Example: SplitSize("100x100") = 100, 100
func SplitSize(size string) (int, int) {
	a := strings.Split(size, "x")
	w := CoerceInt(a[0])
	h := CoerceInt(a[1])

	return RoundTo(w, 5), RoundTo(h, 5)
}

// CoerceInt attempts to convert a string to an integer, defaulting to 0
//
// s: string to convert
//
// Example: CoerceInt("100") = 100
func CoerceInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	return 0
}

// GetEnv attempts to get an environment variable, defaulting to a fallback
//
// key: environment variable to get
//
// fallback: fallback value
//
// Example: GetEnv("PORT", "8080") = "8080"
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// GenerateHash generates an MD5 hash of a string
//
// str: string to hash
//
// Example: GenerateHash("example") = "1a79a4d60de6718e8e5b326e338ae533"
//
// Note: This is not designed to be cryptographically secure
func GenerateMD5Hash(str string) string {
	hash := md5.Sum([]byte(str))

	return hex.EncodeToString(hash[:])
}
