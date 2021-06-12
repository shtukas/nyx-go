package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview" // https://pkg.go.dev/github.com/rivo/tview
)

type Nx19 struct {
	announce string
	uuid     string
}

type Nx27 struct {
	uuid     string
	datetime string
	type_    string
	payload  string
}
type NxEvent struct {
	uuid        string
	datetime    string
	date        string
	description string
}

type NxListing struct {
	uuid        string
	datetime    string
	description string
}

// NxEntity is representing a super type of Nx27 and NxEvent
type NxEntity struct {
	uuid       string
	entityType string
	datetime   string
}

func updateGridContents(grid *tview.Grid, mainElement tview.Primitive, commandField tview.Primitive) {
	grid.
		Clear().
		SetRows(-1, 1).
		SetColumns(-1).
		SetBorders(true).
		AddItem(mainElement, 0, 0, 1, 1, 0, 0, false).
		AddItem(commandField, 1, 0, 1, 1, 0, 0, true)
}

func main() {
	fmt.Println("Nyx")
	fmt.Println("Nx19", Nx19{"announce", "6c249256-2379-4683-bc31-23bbcce4fd39"})
	fmt.Println("Nx27", Nx27{"e1cec0c2-f1ad-4411-9f59-d25cf6bdfa4b", "2021-05-16T17:41:45Z", "unique-string", "a301c45a-e0d1"})
	fmt.Println("NxEvent", NxEvent{"6c249256-2379-4683-bc31-23bbcce4fd39", "2021-05-16T18:14:12Z", "2021-05-16", "description"})
	fmt.Println("NxListing", NxListing{"6c249256-2379-4683-bc31-23bbcce4fd39", "2021-05-16T18:14:12Z", "description"})
	fmt.Println("NxEntity", NxEntity{"6c249256-2379-4683-bc31-23bbcce4fd39", "Nx27", "2021-05-16T18:14:12Z"})

	app := tview.NewApplication()

	grid := tview.NewGrid()

	textView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetText("Hello World")

	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	commandField := tview.NewInputField()

	commandField.
		SetLabel("> ").
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.NewHexColor(0)).
		SetChangedFunc(func(text string) {
			if text == "exit" {
				app.Stop()
				return
			}
			if text == "list" {
				updateGridContents(grid, list, commandField)
				return
			}
			updateGridContents(grid, textView, commandField)
			textView.SetText(text)
		}).
		SetDoneFunc(func(key tcell.Key) {

		})

	updateGridContents(grid, textView, commandField)

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

}
