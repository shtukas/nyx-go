package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview" // https://pkg.go.dev/github.com/rivo/tview
)

type Nx10 struct {
	description string
	uuid        string
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

// NxEntity is representing a super type of Nx10, Nx27, NxEvent and NxListing
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

func makeTextViewFromStrings(strs []string, selectedLineNumber int) *tview.TextView {
	// text := strings.Join(strs[:], "\n")

	text := ""
	for i, v := range strs {
		separator := "\n"
		if i == 0 {
			separator = ""
		}
		if i == selectedLineNumber {
			v = v + " (this one)"
		}
		text = text + separator + v
	}

	return tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetText(text)
}

func main() {
	fmt.Println("Nyx")
	fmt.Println("Nx19", Nx10{"node description", "6c249256-2379-4683-bc31-23bbcce4fd39"})
	fmt.Println("Nx27", Nx27{"e1cec0c2-f1ad-4411-9f59-d25cf6bdfa4b", "2021-05-16T17:41:45Z", "unique-string", "a301c45a-e0d1"})
	fmt.Println("NxEvent", NxEvent{"6c249256-2379-4683-bc31-23bbcce4fd39", "2021-05-16T18:14:12Z", "2021-05-16", "event description"})
	fmt.Println("NxListing", NxListing{"6c249256-2379-4683-bc31-23bbcce4fd39", "2021-05-16T18:14:12Z", "listing description"})
	fmt.Println("NxEntity", NxEntity{"6c249256-2379-4683-bc31-23bbcce4fd39", "Nx27", "2021-05-16T18:14:12Z"})

	app := tview.NewApplication()

	grid := tview.NewGrid()

	textView1 := makeTextViewFromStrings([]string{"Hello World"}, -1)

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
				txV := makeTextViewFromStrings([]string{"List Item 1", text}, 1)
				updateGridContents(grid, txV, commandField)
				return
			}
			updateGridContents(grid, textView1, commandField)
			textView1.SetText(text)
		}).
		SetDoneFunc(func(key tcell.Key) {

		})

	updateGridContents(grid, textView1, commandField)

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

}
