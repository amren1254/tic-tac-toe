package main

import (
	"reflect"
	"testing"
)

func TestDisplayBoardSuccess(t *testing.T) {
	t.Run("Test Board Display ", func(t *testing.T) {
		board := createNewBoard()
		board.displayCurrentBoard()
		reflect.DeepEqual(board.displayCurrentBoard(), nil)
	})
}
