package paper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Fold struct {
	IsHorizontal bool
	Value        int
}

func FoldFromText(text string) (Fold, error) {
	parts := strings.Split(text[11:], "=")
	if len(parts) != 2 {
		return Fold{}, errors.New("not enough parts in fold string")
	}

	isHorizontal := parts[0] == "y"

	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return Fold{}, fmt.Errorf("FoldFromText: %v", err)
	}

	fold := Fold{
		IsHorizontal: isHorizontal,
		Value:        value,
	}

	return fold, nil
}
