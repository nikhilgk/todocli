package ui

import (
	"github.com/gdamore/tcell"
	"github.com/kitagry/go-todotxt"
	"github.com/kitagry/todocli/todo"
	"github.com/rivo/tview"
)

const (
	headerColor = tcell.ColorYellow
)

// Table is a view of todo tasks.
type Table struct {
	tview.Table
}

// NewTable returns a Table in which headers are written.
func NewTable() *Table {
	view := tview.NewTable()

	// Set Header
	view.SetCell(0, 1, tview.NewTableCell("").SetTextColor(headerColor)).
		SetCell(0, 2, tview.NewTableCell("Priority").SetTextColor(headerColor)).
		SetCell(0, 3, tview.NewTableCell("Task").SetTextColor(headerColor)).
		SetCell(0, 4, tview.NewTableCell("Creation Date").SetTextColor(headerColor)).
		SetCell(0, 5, tview.NewTableCell("Completion Date").SetTextColor(headerColor))

	return &Table{
		*view,
	}
}

func (t *Table) getColor(todo *todotxt.Task) tcell.Color {
	switch {
	case todo.Completed:
		return tcell.ColorGray
	case todo.Priority() == 'A':
		return tcell.ColorPeachPuff
	case todo.Priority() == 'B':
		return tcell.ColorAqua
	case todo.Priority() == 'C':
		return tcell.ColorPaleGreen
	}
	return tcell.ColorWhite
}

// WriteTask writes a single task to table.
func (t *Table) WriteTask(task *todotxt.Task, row int) {
	tableColor := t.getColor(task)

	completionText := "☐ "
	if task.Completed {
		completionText = "☑ "
	}

	createdAt := ""
	if !task.CreationDate.IsZero() {
		createdAt = task.CreationDate.Format("2006-01-02")
	}

	completedAt := ""
	if !task.CompletionDate.IsZero() {
		completedAt = task.CompletionDate.Format("2006-01-02")
	}

	t.SetCell(row, 1, tview.NewTableCell(completionText).SetTextColor(tableColor)).
		SetCell(row, 2, tview.NewTableCell(string(task.Priority())).SetTextColor(tableColor)).
		SetCell(row, 3, tview.NewTableCell(task.Description()).SetTextColor(tableColor)).
		SetCell(row, 4, tview.NewTableCell(createdAt).SetTextColor(tableColor)).
		SetCell(row, 5, tview.NewTableCell(completedAt).SetTextColor(tableColor))
}

// WriteTasks writes multiple tasks to table.
func (t *Table) WriteTasks(s *todo.Service) {
	for i := 0; i < s.Length(); i++ {
		todo, err := s.GetTask(i)
		if err != nil {
			continue
		}
		t.WriteTask(todo, i+1)
	}
}
