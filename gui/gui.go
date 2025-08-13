package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/kikiluvv/clockr/export"
	"github.com/kikiluvv/clockr/session"
)

// Start launches the Clockr GUI
func Start() {
	a := app.New()
	w := a.NewWindow("Clockr")
	w.Resize(fyne.NewSize(450, 500))
	w.SetMaster()

	// labels
	statusLabel := widget.NewLabelWithStyle(session.StatusString(), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	summaryLabel := widget.NewLabel(session.SummaryString())

	// helper to refresh both labels
	refreshLabels := func() {
		statusLabel.SetText(session.StatusString())
		summaryLabel.SetText(session.SummaryString())
	}

	// buttons
	inBtn := widget.NewButton("Clock In", func() { session.ClockIn(); refreshLabels() })
	outBtn := widget.NewButton("Clock Out", func() { session.ClockOut(); refreshLabels() })
	breakStartBtn := widget.NewButton("Start Break", func() { session.BreakStart(); refreshLabels() })
	breakEndBtn := widget.NewButton("End Break", func() { session.BreakEnd(); refreshLabels() })
	summaryBtn := widget.NewButton("Show Summary", func() { summaryLabel.SetText(session.SummaryString()) })
	exportBtn := widget.NewButton("Export Hours", func() {
		if err := export.ExportHours(); err != nil {
			summaryLabel.SetText("Export failed: " + err.Error())
		} else {
			summaryLabel.SetText("Export completed!")
		}
	})

	// group buttons by purpose
	clockButtons := container.NewHBox(inBtn, outBtn)
	breakButtons := container.NewHBox(breakStartBtn, breakEndBtn)
	actionButtons := container.NewVBox(summaryBtn, exportBtn)

	buttons := container.NewVBox(
		clockButtons,
		breakButtons,
		widget.NewSeparator(),
		actionButtons,
	)

	// wrap content with padding
	content := container.NewVBox(
		statusLabel,
		layout.NewSpacer(),
		buttons,
		layout.NewSpacer(),
		summaryLabel,
	)

	w.SetContent(container.NewPadded(content))
	w.ShowAndRun()
}
