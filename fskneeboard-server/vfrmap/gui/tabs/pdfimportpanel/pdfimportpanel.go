package pdfimportpanel

import (
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/application/pdfimport"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/gui/tabs/panelcommons"
	"vfrmap-for-vr/vfrmap/logger"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var importRunningBinding = binding.NewBool()
var statusBinding = binding.NewString()
var fileListBinding = binding.NewStringList()

func updateStatus(status string) {
	statusBinding.Set(status)
}

func runImport() {
	logger.LogInfoVerbose("Starting PDF import...")
	updateStatus("Preparing PDF batch import...")

	err := pdfimport.ImportPdfFolder(updateStatus)

	importRunningBinding.Set(false)

	if err != nil {
		logger.LogErrorVerbose("Something went wrong, reason: " + err.Error())
		updateStatus("PDF batch import failed!")

		dialogs.ShowError("PDF Import failed! Please refer to the Console Panel and/or logs for details!")
	} else {
		logger.LogInfoVerbose("Import finished!")
		updateStatus("PDF batch import finished!")

		dialogs.ShowInformation("The PDF Import finished!")
	}
}

func processImporterDownloadPromptCallback(proceed bool) {
	if !proceed {
		importRunningBinding.Set(false)
		return
	}

	updateStatus("Downloading importer module...")

	downloadErr := pdfimport.DownloadImporter()

	if downloadErr != nil {
		logger.LogErrorVerbose("Could not download the importer, reason: " + downloadErr.Error())
		dialogs.ShowError("Could not download the importer. Please refer to the Console Panel and/or logs for details!")
		updateStatus("Idle...")
		importRunningBinding.Set(false)
		return
	}

	runImport()
}

func clearImportFolderPromptCallback(proceed bool) {
	if proceed {
		updateStatus("Clearing PDF import folder...")
		err := pdfimport.ClearPdfImportFolder()

		if err != nil {
			logger.LogErrorVerbose("Could not clear PDF import folder, reason: " + err.Error())
			updateStatus("Clearing PDF import folder failed!")

			dialogs.ShowError("PDF import folder could not be cleared. Please refer to the Console Panel and/or logs for details!")
		}

		go refreshImportDir()

		updateStatus("PDF import folder cleared!")
	}
}

func refreshImportDir() {
	fileListBinding.Set([]string{})
	list, err := pdfimport.CreatePdfFileList()

	if err != nil {
		logger.LogErrorVerbose("Could not refresh PDF import folder, reason: " + err.Error())
		updateStatus("Refreshing PDF import folder failed!")

		dialogs.ShowError("PDF import folder could not be refreshed. Please refer to the Console Panel and/or logs for details!")
	}

	sList := []string{}
	for _, info := range list {
		sList = append(sList, info.FileName)
	}

	fileListBinding.Set(sList)
}

func PdfImportPanel() *fyne.Container {
	logger.LogDebug("Initializing PDF Import Panel...")

	// top
	refreshImportDirBtn := widget.NewButtonWithIcon("Refresh Import Directory", theme.ViewRefreshIcon(), func() {
		go func() {
			updateStatus("Refreshing PDF import folder...")
			refreshImportDir()
			updateStatus("PDF import folder refreshed!")
		}()
	})

	clearImportDirBtn := widget.NewButtonWithIcon("Clear Import Directory", theme.ContentClearIcon(), func() {
		dialogs.ShowClearImportFolderPrompt(clearImportFolderPromptCallback)
	})

	openImportDirBtn := widget.NewButtonWithIcon("Open Import Directory", theme.FolderOpenIcon(), func() {
		pdfimport.OpenPdfSourceFolder()
	})

	top := container.NewHBox(refreshImportDirBtn, clearImportDirBtn, openImportDirBtn)

	// bottom
	progressBar := widget.NewProgressBarInfinite()
	progressBar.Stop()

	statusLabel := widget.NewLabelWithData(statusBinding)
	statusLabel.Alignment = fyne.TextAlignCenter
	statusBinding.Set("Idle...")

	startImportBtn := widget.NewButtonWithIcon("Start Import", theme.MediaPlayIcon(), func() {
		go func() {
			importRunningBinding.Set(true)

			refreshImportDir()

			if pdfimport.HasImporter() {
				runImport()
			} else {
				logger.LogWarnVerbose("Importer module not found!")
				dialogs.ShowImporterDownloadPrompt(processImporterDownloadPromptCallback)
			}
		}()
	})

	importRunningBinding.AddListener(binding.NewDataListener(func() {
		importRunning, _ := importRunningBinding.Get()

		if importRunning {
			startImportBtn.Disable()
			progressBar.Start()
		} else {
			startImportBtn.Enable()
			progressBar.Stop()
		}
	}))

	openOutputDirBtn := widget.NewButtonWithIcon("Open Output Directory", theme.FolderOpenIcon(), func() {
		pdfimport.OpenPdfOutFolder()
	})

	bottomButtons := container.NewHBox(startImportBtn, openOutputDirBtn)
	bottom := container.NewVBox(progressBar, statusLabel, bottomButtons)

	// border layout
	border := layout.NewBorderLayout(top, bottom, nil, nil)
	var resContainer *fyne.Container

	// middle
	if globals.Pro {
		fileList := widget.NewListWithData(
			fileListBinding,
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			func(i binding.DataItem, o fyne.CanvasObject) {
				o.(*widget.Label).Bind(i.(binding.String))
			})
		fileList.OnSelected = func(id widget.ListItemID) {
			fileList.UnselectAll()
		}

		resContainer = container.New(border, top, bottom, fileList)
	} else {
		info := panelcommons.PremiumInfo("The Charts Viewer and PDF Importer are available to PRO supporters!")
		resContainer = container.New(border, top, bottom, info)
	}

	if globals.Pro {
		go refreshImportDir()
	} else {
		refreshImportDirBtn.Disable()
		clearImportDirBtn.Disable()
		openImportDirBtn.Disable()
		startImportBtn.Disable()
		openOutputDirBtn.Disable()
	}

	logger.LogDebug("PDF Import Panel initialized")

	return resContainer
}
