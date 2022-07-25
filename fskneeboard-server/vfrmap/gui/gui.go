package gui

import (
	"vfrmap-for-vr/vfrmap/application/globals"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/gui/res"
	"vfrmap-for-vr/vfrmap/gui/tabs/consolepanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/controlpanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/hotkeyspanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/pdfimportpanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/settingspanel"
	"vfrmap-for-vr/vfrmap/gui/tabs/supportpanel"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var w fyne.Window
var guiTourStartedBinding binding.NewBoolean()

func UpdateGuiTourStarted(started bool) {
	guiTourStartedBinding.Set(started)
}

func InitGui() {
	utils.Println("Starting FSKneeboard GUI...")

	a := app.New()

	logger.LogDebugVerboseOverride("Loading icon...", false)
	iconAsset, err := res.Asset("icon.png")
	if err == nil {
		iconResource := fyne.NewStaticResource("icon.png", iconAsset)
		logger.LogDebugVerboseOverride("Icon loaded", false)
		a.SetIcon(iconResource)
	} else {
		logger.LogWarnVerboseOverride("Icon could not be loaded!", false)
	}

	title := globals.ProductName + " v" + globals.BuildVersion

	w = a.NewWindow(title)

	logger.LogDebugVerboseOverride("Initializing tabs...", false)
	welcomeTab := conatiner.NewTabItem("Welcome", controlpanel.WelcomePanel());

	tabs := container.NewAppTabs(
		welcomeTab,
		container.NewTabItem("Control Panel", controlpanel.ControlPanel()),
		container.NewTabItem("Settings", settingspanel.SettingsPanel()),
		container.NewTabItem("Hotkeys", hotkeyspanel.HotkeysPanel()),
		container.NewTabItem("PDF Import", pdfimportpanel.PdfImportPanel()),
		container.NewTabItem("Console", consolepanel.ConsolePanel()),
		container.NewTabItem("Get Support", supportpanel.SupportPanel()),
	)

	guiTourStartedBinding.AddListener(binding.NewDataListener(func() {
		guiTourStarted, _ := guiTourStartedBinding.Get()

		if guiTourStarted {
			welcomeTab.Show()
		} else {
			welcomeTab.Hide()
		}
	}))

	logger.LogDebugVerboseOverride("Tabs initialized", false)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(800, 600))

	dialogs.ParentWindow = &w
}

func ShowAndRun() {
	if w != nil {
		logger.LogDebug("Showing window...")
		w.ShowAndRun()
	}
}
