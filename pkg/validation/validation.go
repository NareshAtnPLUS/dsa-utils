package validation

import (
	"github.com/NareshAtnPLUS/dsa-utils/pkg/datastructures"
	"github.com/NareshAtnPLUS/dsa-utils/pkg/utils"
)

func IsOutOfBounds(pos datastructures.Coordinates, gridSize *datastructures.GridSize) bool {

	if utils.Min(pos.R, pos.C) < 0 ||
		pos.R >= gridSize.Rows || pos.C >= gridSize.Cols {
		return true
	}

	return false
}
