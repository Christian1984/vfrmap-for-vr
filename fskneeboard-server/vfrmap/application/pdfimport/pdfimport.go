package pdfimport

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"vfrmap-for-vr/vfrmap/gui/dialogs"
	"vfrmap-for-vr/vfrmap/logger"
	"vfrmap-for-vr/vfrmap/utils"
)

var importerBasePath, _ = filepath.Abs("pdf-importer")
var importerCallerPath, _ = filepath.Abs(importerBasePath + "\\" + "pdf-import-runner.exe")
var importerExePath, _ = filepath.Abs(importerBasePath + "\\" + "gswin64c.exe")
var importerDllPath, _ = filepath.Abs(importerBasePath + "\\" + "gsdll64.dll")
var importerLicensePath, _ = filepath.Abs(importerBasePath + "\\" + "THIRD-PARTY-LICENSE.md")

const importerBaseUrl = "https://github.com/Christian1984/pdf-importer/releases/download/v1.1.1/"
const importerCallerUrl = importerBaseUrl + "pdf-import-runner.exe"
const importerExeUrl = importerBaseUrl + "gswin64c.exe"
const importerDllUrl = importerBaseUrl + "gsdll64.dll"
const importerLicenseUrl = importerBaseUrl + "THIRD-PARTY-LICENSE.md"

const sourceFolder = "charts\\!import"
const outFolder = "charts\\imported"

type PdfFileInfo struct {
	SourcePath string
	TargetPath string
	FileName   string
}

func importPdfChart(sourcePath string, targetBasePath string, filename string) error {
	logger.LogDebug("Enter importPdfChart...")

	documentTargetPath, _ := filepath.Abs(targetBasePath + "\\" + strings.TrimSuffix(filename, ".pdf"))

	in, _ := filepath.Abs(sourcePath + "\\" + filename)
	out, _ := filepath.Abs(documentTargetPath + "\\" + strings.TrimSuffix(filename, ".pdf") + "-%03d.png")

	logger.LogInfoVerbose("Starting PDF Import of " + in)
	logger.LogDebug("Creating out path at " + documentTargetPath)
	mkDirErr := os.MkdirAll(documentTargetPath, os.ModePerm)

	if mkDirErr != nil {
		logger.LogErrorVerbose("Could not create target folder, reason: " + mkDirErr.Error())
		return mkDirErr
	}

	cmdParams := []string{
		"--out",
		out,
		"--in",
		in,
		"--gspath",
		".\\pdf-importer",
		"--verbose",
	}

	cmd := exec.Command(importerCallerPath, cmdParams...)
	logger.LogDebug("Import command is: " + cmd.String())

	s, importErr := cmd.Output()
	result := string(s)
	logger.LogDebug("Import output:\n" + result)

	if importErr != nil {
		logger.LogErrorVerbose("Could not import PDF file, reason: " + importErr.Error())
		return importErr
	} else {
		logger.LogInfoVerbose("Import successful!")
	}

	return nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if errors.Is(err, os.ErrNotExist) {
		logger.LogWarn(filePath + " not found, reason: " + err.Error())
		return false
	} else if err != nil {
		logger.LogWarn("Something when wrong when searching for " + filePath + ", reason: " + err.Error())
		return false
	}

	logger.LogInfo(filePath + " found!")
	return true
}

func downloadFile(filepath string, url string) error {
	logger.LogInfo("Downloading from " + url + " to " + filepath)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func createAndOpenFolder(path string) {
	absPath, _ := filepath.Abs(path)
	os.MkdirAll(absPath, os.ModePerm)

	logger.LogDebug("Trying to open folder [" + path + "]...")
	err := utils.OpenExplorer(path)

	if err != nil {
		logger.LogErrorVerbose("Could not open folder [" + path + "]")
		dialogs.ShowError("Folder could not be opened! Reason: " + err.Error())
	}
}

func CreatePdfFileList() ([]PdfFileInfo, error) {
	list := []PdfFileInfo{}

	absPath, _ := filepath.Abs(sourceFolder)
	sourceFolderErr := os.MkdirAll(absPath, os.ModePerm)

	if sourceFolderErr != nil {
		logger.LogErrorVerbose("Import failed! Could not find or create import source folder, reason: " + sourceFolderErr.Error())
		return list, sourceFolderErr
	}

	walkErr := filepath.Walk(sourceFolder, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			logger.LogDebug("Walking import dir, enter directory: " + filePath + ", info: " + fmt.Sprintf("%v", info))
		} else {
			// check if pdf
			if !strings.HasSuffix(strings.ToLower(info.Name()), ".pdf") {
				logger.LogDebug("Skipping file " + info.Name() + ", reason: Extension [.pdf] missing!")
				return nil
			}

			logger.LogDebug("Walking import dir, current file: " + filePath + ", info: " + fmt.Sprintf("%v", info))
			sourcePath := strings.TrimSuffix(filePath, "\\"+info.Name()) // e.g. charts\!import\folder 1
			relPath := strings.TrimPrefix(sourcePath, sourceFolder)      // e.g. \folder 1
			relPath = strings.TrimPrefix(relPath, "\\")                  // e.g. folder 1
			targetPath := outFolder + "\\" + relPath                     // e.g. charts\imported\folder 1

			logger.LogDebug("filePath: " + filePath)
			logger.LogDebug("sourcePath: " + sourcePath)
			logger.LogDebug("sourceRoot: " + sourceFolder)
			logger.LogDebug("relPath: " + relPath)
			logger.LogDebug("targetPath: " + targetPath)

			list = append(list, PdfFileInfo{SourcePath: sourcePath, TargetPath: targetPath, FileName: info.Name()})
		}

		return nil
	})

	if walkErr != nil {
		logger.LogErrorVerbose("Could not import folder " + sourceFolder + ", reason: " + walkErr.Error())
		return list, walkErr
	}

	return list, nil
}

func ImportPdfFolder(updateStatusCallback func(string)) error {
	logger.LogDebug("Enter ImportPdfFolder...")

	pdfList, listErr := CreatePdfFileList()

	if listErr != nil {
		logger.LogErrorVerbose("Could not create list of PDF files, reason: " + listErr.Error())
		return listErr
	}

	if len(pdfList) == 0 {
		logger.LogWarnVerbose("No PDFs found in import folder!")
		return errors.New("No PDFs found in import folder!")
	}

	for _, fileInfo := range pdfList {
		if updateStatusCallback != nil {
			updateStatusCallback("Importing PDF: [" + fileInfo.FileName + "]")
		}

		importErr := importPdfChart(fileInfo.SourcePath, fileInfo.TargetPath, fileInfo.FileName)

		if importErr != nil {
			logger.LogErrorVerbose("Could not import file " + fileInfo.FileName + ", reason: " + importErr.Error())
			return importErr
		}
	}

	logger.LogInfoVerbose("Import process finished!")
	return nil
}

func HasImporter() bool {
	if !fileExists(importerCallerPath) || !fileExists(importerExePath) || !fileExists(importerDllPath) {
		logger.LogWarnVerbose("Local importer binaries not found or incomplete!")
		return false
	}

	logger.LogInfoVerbose("Local importer binaries found!")
	return true
}

func DownloadImporter() error {
	logger.LogInfoVerbose("Downloading Importer...")

	mkDirErr := os.MkdirAll(importerBasePath, os.ModePerm)

	if mkDirErr != nil {
		logger.LogError("Could not create importer base directory, reason: " + mkDirErr.Error())
		return mkDirErr
	}

	if err := downloadFile(importerCallerPath, importerCallerUrl); err != nil {
		logger.LogError("Could not download importer caller, reason: " + err.Error())
		return err
	}

	if err := downloadFile(importerExePath, importerExeUrl); err != nil {
		logger.LogError("Could not download importer executable, reason: " + err.Error())
		return err
	}

	if err := downloadFile(importerDllPath, importerDllUrl); err != nil {
		logger.LogError("Could not download importer dll, reason: " + err.Error())
		return err
	}

	if err := downloadFile(importerLicensePath, importerLicenseUrl); err != nil {
		logger.LogError("Could not download importer license, reason: " + err.Error())
		return err
	}

	logger.LogInfoVerbose("Importer successfully downloaded...")

	return nil
}

func OpenPdfSourceFolder() {
	logger.LogDebug("Trying to open PDF import source folder [" + sourceFolder + "]...")
	createAndOpenFolder(sourceFolder)
}

func OpenPdfOutFolder() {
	logger.LogDebug("Trying to open PDF import out folder [" + outFolder + "]...")
	createAndOpenFolder(outFolder)
}

func ClearPdfImportFolder() error {
	absSourcePath, _ := filepath.Abs(sourceFolder)
	err := os.RemoveAll(absSourcePath)

	if err != nil {
		logger.LogErrorVerbose("PDF import source folder [" + sourceFolder + "] could not be cleared, reason: " + err.Error())
		return err
	}

	os.MkdirAll(absSourcePath, os.ModePerm)

	return nil
}
