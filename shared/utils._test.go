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

import "testing"

func TestRoundTo1(t *testing.T) {
	r := RoundTo(100, 5)
	if r != 100 {
		t.Error("Did not leave 100 at 100")
	}
}

func TestRoundTo2(t *testing.T) {
	r := RoundTo(106, 10)
	if r != 100 {
		t.Error("Did not round 106 to 110")
	}
}

func TestRoundTo3(t *testing.T) {
	r := RoundTo(94, 10)
	if r != 90 {
		t.Error("Did not round 94 to 90")
	}
}

func TestRoundTo4(t *testing.T) {
	r := RoundTo(0, 10)
	if r != 0 {
		t.Error("Did not leave 0 as 0")
	}
}

func TestSplitSize1(t *testing.T) {
	w, h := SplitSize("100x100")
	if (w != 100) || (h != 100) {
		t.Error("Did not split 100x100 into 100x100")
	}
}

func TestSplitSize2(t *testing.T) {
	w, h := SplitSize("100x100x5")
	if (w != 100) || (h != 100) {
		t.Error("Did not split 100x100x5 into 100x100")
	}
}

func TestSplitSize3(t *testing.T) {
	w, h := SplitSize("302x949")
	if (w != 300) || (h != 950) {
		t.Error("Did not split 302x949 into 300x950")
	}
}

func TestSplitSize4(t *testing.T) {
	w, h := SplitSize("500x")
	if (w != 500) || (h != 0) {
		t.Error("Did not split 500x nothing into 500x0")
	}
}
