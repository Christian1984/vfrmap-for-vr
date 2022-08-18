package pdfimport

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"vfrmap-for-vr/vfrmap/logger"
)

var importToolBasePath, _ = filepath.Abs("pdf-importer")

var importToolPath, _ = filepath.Abs(importToolBasePath + "\\" + "pdf-importer.exe")

const importerBaseUrl = "https://github.com/Christian1984/pdf-import-tool/releases/download/v0.0.1/"

type BinFileInfo struct {
	Url      string
	FilePath string
	Checksum string
}

var DownloadFiles = []BinFileInfo{
	MakeBinFileInfo("pdf-importer.exe", ""),
	MakeBinFileInfo("gswin64c.exe", ""),
	MakeBinFileInfo("gsdll64.dll", ""),
	MakeBinFileInfo("THIRD-PARTY-LICENSE.md", ""),
}

type PdfFileInfo struct {
	SourcePath string
	TargetPath string
	FileName   string
}

func MakeBinFileInfo(filename string, checksum string) BinFileInfo {
	path, _ := filepath.Abs(importToolBasePath + "\\" + filename)
	return BinFileInfo{Url: importerBaseUrl + filename, FilePath: path, Checksum: checksum}
}

func StartImporter() error {
	logger.LogDebug("Enter RunImporter...")

	absInFolder, _ := filepath.Abs("charts\\!import")
	absOutFolder, _ := filepath.Abs("charts\\imported")

	cmdParams := []string{
		"--in",
		absInFolder,
		"--out",
		absOutFolder,
		"--lib",
		importToolBasePath,
	}

	cmd := exec.Command(importToolPath, cmdParams...)
	logger.LogDebug("Import command is: " + cmd.String())

	startErr := cmd.Run()

	if startErr != nil {
		logger.LogErrorVerbose("Could not start PDF-Import-Tool, reason: " + startErr.Error())
		return startErr
	}

	logger.LogInfoVerbose("PDF-Import-Tool started!")
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

func checksumValid(fi BinFileInfo) bool {
	f, err := os.Open(fi.FilePath)
	defer f.Close()

	if err != nil {
		logger.LogErrorVerbose("Cannot calculate checksum for file " + fi.FilePath + ", cannot open file!")
		return false
	}

	hash := md5.New()

	_, copyErr := io.Copy(hash, f)

	if copyErr != nil {
		logger.LogErrorVerbose("Cannot calculate checksum for file " + fi.FilePath + ", cannot copy to hasher!")
	}

	isHash := hash.Sum(nil)
	isHashString := fmt.Sprintf("%x", isHash)
	logger.LogDebugVerbose("File " + fi.FilePath + " has md5 hash of [" + isHashString + "], expected hash is [" + fi.Checksum + "]")

	return true
}

func HasValidImporter() bool {
	for _, fi := range DownloadFiles {
		if !fileExists(fi.FilePath) {
			logger.LogWarnVerbose("Local importer binaries not found or incomplete, file: " + fi.FilePath)
			return false
		}

		if !checksumValid(fi) {
			logger.LogWarnVerbose("File integrity check failed for file: " + fi.FilePath)
			return false
		}
	}

	logger.LogInfoVerbose("Local importer binaries found!")
	return true
}

func DownloadImporter() error {
	logger.LogInfoVerbose("Downloading Importer...")

	mkDirErr := os.MkdirAll(importToolBasePath, os.ModePerm)

	if mkDirErr != nil {
		logger.LogError("Could not create importer base directory, reason: " + mkDirErr.Error())
		return mkDirErr
	}

	for _, fi := range DownloadFiles {
		if err := downloadFile(fi.FilePath, fi.Url); err != nil {
			logger.LogError("Could not download file, reason: " + err.Error())
			return err
		}
	}

	logger.LogInfoVerbose("Importer successfully downloaded...")

	return nil
}
