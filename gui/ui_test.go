package gui

import (
	//gui "accounter/gui"
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	acc:=NewApp()
	acc.LoadUI(test.NewApp())
	test.Type(acc.IncomeEntry, "0.01")
	test.Tap(acc.AddBtn)
}
