package charts

type PdfFileInfo struct {
	SourcePath string
	TargetPath string
	FileName   string
}

func CreatePdfFileList() ([]PdfFileInfo, error) {
	return []PdfFileInfo{}, nil
}

func ImportPdfFolder(updateStatusCallback func(string)) error {
	return nil
}

func HasGhostscript() bool {
	return false
}

func DownloadGhostscript() error {
	return nil
}

func OpenPdfSourceFolder() {}

func OpenPdfOutFolder() {}

func ClearPdfImportFolder() error {
	return nil
}
