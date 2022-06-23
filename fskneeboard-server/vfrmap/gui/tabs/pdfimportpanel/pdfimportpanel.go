package pdfimportpanel

import (
	"vfrmap-for-vr/_vendor/premium/charts"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var importRunningBinding = binding.NewBool()

func runImport() {
	logger.LogInfoVerbose("Starting PDF import...")

	//err := charts.ImportPdfChart("charts\\!import", "charts\\imported", "test.pdf")
	err := charts.ImportPdfFolder("charts\\!import", "charts\\imported")

	importRunningBinding.Set(false)

	if err != nil {
		logger.LogErrorVerbose("Something went wrong, reason: " + err.Error())
		dialogs.ShowError("PDF Import failed! Please refer to the Console Panel and/or logs for details!")
	} else {
		logger.LogInfoVerbose("Import finished!")
		dialogs.ShowInformation("The PDF Import finished!")
	}
}

func processGhostScriptDownloadPromptCallback(proceed bool) {
	if !proceed {
		importRunningBinding.Set(false)
		return
	}

	downloadErr := charts.DownloadGhostscript()

	if downloadErr != nil {
		logger.LogErrorVerbose("Could not download ghostscript, reason: " + downloadErr.Error())
		dialogs.ShowError("Could not download Ghostscript. Please refer to the Console Panel and/or logs for details!")
		importRunningBinding.Set(false)
		return
	}

	runImport()

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
				dialogs.ShowGhostscriptDownloadPrompt(processGhostScriptDownloadPromptCallback)
			}
		}()
	})

	importRunningBinding.AddListener(binding.NewDataListener(func() {
		importRunning, _ := importRunningBinding.Get()

		if importRunning {
			button.Disable()
			progressBar.Show()
			progressBar.Start()
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
