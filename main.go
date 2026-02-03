package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	VERSION         = "0.2.0"
	RUN_SH          = "run.sh"
	BOOTSTRAP       = "bootstrap"
	DEFAULT_EXEC    = "bootstrap,run.sh"
	EXEC_PERMISSION = 0o755
)

type packOptions struct {
	execFiles string
	lfFiles   string
}

var opts = packOptions{
	execFiles: DEFAULT_EXEC,
	lfFiles:   RUN_SH,
}

func parseExecFiles(execFiles string) []string {
	if execFiles == "" {
		return []string{}
	}
	parts := strings.Split(execFiles, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func createFileHeader(info os.FileInfo, baseDir, path string) (*zip.FileHeader, error) {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return nil, err
	}

	relPath, err := filepath.Rel(baseDir, path)
	if err != nil {
		return nil, err
	}
	header.Name = filepath.ToSlash(relPath)

	if info.IsDir() {
		header.Name += "/"
	} else {
		header.Method = zip.Deflate
	}

	return header, nil
}

func applyExecutablePermission(header *zip.FileHeader, info os.FileInfo, execFiles []string) bool {
	for _, execFile := range execFiles {
		if info.Name() == execFile {
			header.SetMode(EXEC_PERMISSION)
			return true
		}
	}
	return false
}

func shouldConvertLineEndings(info os.FileInfo, lfFiles []string) bool {
	for _, lfFile := range lfFiles {
		if info.Name() == lfFile {
			return true
		}
	}
	return false
}

func writeFileContent(writer io.Writer, path string, shouldConvertLF bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "Error closing file:", err)
		}
	}()

	if shouldConvertLF {
		content, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		crlf := []byte{'\r', '\n'}
		lf := []byte{'\n'}
		content = bytes.ReplaceAll(content, crlf, lf)

		_, err = writer.Write(content)
		if err != nil {
			return err
		}
	} else {
		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func warnAboutMissingExecFiles(execFiles []string, foundExecFiles map[string]bool) {
	for _, execFile := range execFiles {
		if !foundExecFiles[execFile] {
			fmt.Fprintf(
				os.Stderr,
				"Warning: exec file '%s' not found in source directory\n",
				execFile,
			)
		}
	}
}

func initFlags() {
	flag.StringVar(
		&opts.execFiles,
		"exec",
		DEFAULT_EXEC,
		"Comma-separated list of files to set executable permissions",
	)
	flag.StringVar(
		&opts.lfFiles,
		"lf",
		RUN_SH,
		"Comma-separated list of files to convert CRLF line endings to LF",
	)
}

// zipDirectory compresses the entire contents of the sourceDir directory
// into a zip file named targetZipFile.
//
// If files are specified via --exec flag, they are given Unix executable
// permissions (0755). The default is "bootstrap,run.sh".
//
// If files are specified via --lf flag, their line endings are converted
// from CRLF to LF. The default is "run.sh".
func zipDirectory(sourceDir, targetZipFile string, execFiles, lfFiles []string) error {
	zipFile, err := os.Create(targetZipFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := zipFile.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "Error closing zip file:", err)
		}
	}()

	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if err := zipWriter.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "Error closing zip writer:", err)
		}
	}()

	baseDir := filepath.Clean(sourceDir)
	foundExecFiles := make(map[string]bool)

	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == baseDir {
			return nil
		}

		header, err := createFileHeader(info, baseDir, path)
		if err != nil {
			return err
		}

		if !info.IsDir() && applyExecutablePermission(header, info, execFiles) {
			foundExecFiles[info.Name()] = true
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			shouldConvertLF := shouldConvertLineEndings(info, lfFiles)
			if err := writeFileContent(writer, path, shouldConvertLF); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	warnAboutMissingExecFiles(execFiles, foundExecFiles)

	return nil
}

func main() {
	initFlags()
	flag.Parse()
	// get cmd arguments
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: lzpb [--exec=files] [--lf=files] <source_dir> <target_zip>")
		return
	}
	sourceDir := args[0]
	targetZip := args[1]
	execFiles := parseExecFiles(opts.execFiles)
	lfFiles := parseExecFiles(opts.lfFiles)
	if err := zipDirectory(sourceDir, targetZip, execFiles, lfFiles); err != nil {
		fmt.Fprintln(os.Stderr, "Error creating zip file:", err)
		return
	}
}
