package gui

import (
	"testing"

	// "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
	// "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	//var newError = dialog.NewError()

	acc := NewApp()
	acc.LoadUI(test.NewApp())
	test.Type(acc.IncomeEntry, "")
	test.Tap(acc.AddBtn)
	t.Errorf("",)
}
