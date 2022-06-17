package pdfimportpanel

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PdfImportPanel() *fyne.Container {
	logger.LogDebugVerboseOverride("Initializing PDF Import Panel...", false)

	// grid and centerContainer
	button := widget.NewButton("Import", func() {
		logger.LogDebug("Starting PDF Import...")

		in, _ := filepath.Abs("gs\\in\\test.pdf")
		out, _ := filepath.Abs("gs\\out\\test--%03d.png")
		//out := "gs\\test--%03d.png"

		cmdParams := []string{
			//"-q",
			//"-dQUIET",
			"-dSAFER",
			"-dBATCH",
			"-dNOPAUSE",
			"-dNOPROMPT",
			"-dMaxBitmap=500000000",
			"-dAlignToPixels=0",
			"-dGridFitTT=2",
			"-sDEVICE=png16m",
			"-dTextAlphaBits=4",
			"-dGraphicsAlphaBits=4",
			"-r150x150",

			"-o",
			out,
			in,
		}

		cmd := exec.Command(".\\gs\\gswin64c.exe", cmdParams...)
		logger.LogDebugVerbose("Import command is: " + cmd.String())

		s, importErr := cmd.Output()
		result := string(s)
		fmt.Println(result)

		/*cmd := exec.Command(".\\gs\\gswin64c.exe", cmdParams...)
		logger.LogDebugVerbose("Import command is: " + cmd.String())

		importErr := cmd.Run()*/

		if importErr != nil {
			logger.LogErrorVerbose("Could not import PDF file, reason: " + importErr.Error())
		} else {
			logger.LogInfoVerbose("Import successful!")
		}

		logger.LogInfoVerbose("Import process finished!")
	})

	centerContainer := container.NewCenter(button)

	logger.LogDebugVerboseOverride("PDF Import Panel initialized", false)

	return centerContainer

}
