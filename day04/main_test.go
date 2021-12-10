package main

import "testing"

func TestBingo_HasColumnBingo(t *testing.T) {
	nums := []int{0,1,2,3,4}

	for _, num := range nums {
		t.Run("Column", func(t *testing.T) {
			board := Board{}
			for i := 0; i < 25; i++ {
				value := 0
				if i % 5 == num {
					value = i / 5
					if num == 0 {
						value++
					}
				}

				board[i] = &Box{value, false}
			}

			bingo := NewBingo(board)

			for i := 1; i <= 5; i++ {
				bingo.CheckNumber(i)
			}

			hasBingo := bingo.HasColumnBingo(num)
			if ! hasBingo {
				t.Errorf("Expected a Bingo, did not get any. col=%d\n%+v",num, bingo.Board.String())
			}
		})
	}
}
