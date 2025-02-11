// Copyright (C) 2021 Tweag IO
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

package decider

import (
	"fmt"
)

type PercentageDecider struct {
	Minimum int `toml:"minimum"`
}

func (s *PercentageDecider) Validate() error {
	if s.Minimum < 0 || s.Minimum > 100 {
		return fmt.Errorf("Minimum percentage decider out of bounds (%d)", s.Minimum)
	}
	return nil
}
