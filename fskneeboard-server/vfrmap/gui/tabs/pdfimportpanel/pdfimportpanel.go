package pdfimportpanel

import (
	"vfrmap-for-vr/_vendor/premium/charts"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var importRunningBinding = binding.NewBool()

func runImport() {
	logger.LogInfoVerbose("Starting PDF import...") // TODO: show dialog

	err := charts.ImportPdfChart("charts\\!import", "charts\\imported", "test.pdf")

	if err != nil {
		logger.LogErrorVerbose("Something went wrong, reason: " + err.Error()) // TODO: show dialog
	} else {
		logger.LogInfoVerbose("Import finished!") // TODO: show dialog
	}
}

func PdfImportPanel() *fyne.Container {
	logger.LogDebug("Initializing PDF Import Panel...")

	progressBar := widget.NewProgressBarInfinite()
	progressBar.Stop()
	progressBar.Hide()

	button := widget.NewButton("Import", func() {
		go func() {
			importRunningBinding.Set(true)

			if charts.HasGhostscript() {
				runImport()
			} else {
				logger.LogWarnVerbose("Ghostscript not found!") // TODO
				downloadErr := charts.DownloadGhostscript()

				if downloadErr != nil {
					logger.LogErrorVerbose("Could not download ghostscript, reason: " + downloadErr.Error()) // TODO
				} else {
					runImport()
				}
			}

			importRunningBinding.Set(false)
		}()
	})

	importRunningBinding.AddListener(binding.NewDataListener(func() {
		importRunning, _ := importRunningBinding.Get()

		if importRunning {
			button.Disable()
			progressBar.Start()
			progressBar.Show()
		} else {
			button.Enable()
			progressBar.Stop()
			progressBar.Hide()
		}
	}))

	vBox := container.NewVBox(progressBar, button)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebug("PDF Import Panel initialized")

	return centerContainer

}
