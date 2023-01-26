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

package utilities

import "testing"

func TestSplitSize(t *testing.T) {
	t.Log("Test 1")
	w, h := SplitSize("100x100")
	if (w != 100) || (h != 100) {
		t.Error("Did not correctly split 100x100")
	}

	t.Log("Test 2")
	w, h = SplitSize("100x100x5")
	if (w != 100) || (h != 100) {
		t.Error("Did not correctly split 100x100x5")
	}
}
