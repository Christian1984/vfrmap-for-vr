package pdfimportpanel

import (
	"fmt"
	"vfrmap-for-vr/_vendor/premium/charts"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var importRunningBinding = binding.NewBool()

func PdfImportPanel() *fyne.Container {
	logger.LogDebugVerboseOverride("Initializing PDF Import Panel...", false)

	progressBar := widget.NewProgressBarInfinite()
	progressBar.Stop()

	button := widget.NewButton("Import", func() {
		go func() {
			importRunningBinding.Set(true)

			err := charts.ImportPdfChart("charts\\!import", "charts\\imported", "test.pdf")

			if err != nil {
				fmt.Println("Something went wrong:", err.Error()) // TODO: show dialog
			} else {
				fmt.Println("Import finished!") // TODO: show dialog
			}

			importRunningBinding.Set(false)
		}()
	})

	importRunningBinding.AddListener(binding.NewDataListener(func() {
		importRunning, _ := importRunningBinding.Get()

		if importRunning {
			button.Disable()
			progressBar.Start()
		} else {
			button.Enable()
			progressBar.Stop()
		}
	}))

	vBox := container.NewVBox(progressBar, button)
	centerContainer := container.NewCenter(vBox)

	logger.LogDebugVerboseOverride("PDF Import Panel initialized", false)

	return centerContainer

}
