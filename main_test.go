package main

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

func TestParseExecFiles(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single file",
			input:    "script.sh",
			expected: []string{"script.sh"},
		},
		{
			name:     "multiple files",
			input:    "script.sh,bin/app,tools/build",
			expected: []string{"script.sh", "bin/app", "tools/build"},
		},
		{
			name:     "files with spaces",
			input:    "script.sh, bin/app ,tools/build",
			expected: []string{"script.sh", "bin/app", "tools/build"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "default value",
			input:    "bootstrap,run.sh",
			expected: []string{"bootstrap", "run.sh"},
		},
		{
			name:     "whitespace only",
			input:    "   ",
			expected: []string{},
		},
		{
			name:     "consecutive commas",
			input:    "file1,,file2",
			expected: []string{"file1", "file2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseExecFiles(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf(
					"parseExecFiles(%q) returned %d files, expected %d",
					tt.input,
					len(result),
					len(tt.expected),
				)
				return
			}
			for i, file := range result {
				if file != tt.expected[i] {
					t.Errorf(
						"parseExecFiles(%q)[%d] = %q, expected %q",
						tt.input,
						i,
						file,
						tt.expected[i],
					)
				}
			}
		})
	}
}

func TestZipDirectory(t *testing.T) {
	t.Run("single exec file", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcDir := filepath.Join(tmpDir, "src")
		if err := os.MkdirAll(srcDir, 0o755); err != nil {
			t.Fatalf("failed to create src dir: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "script.sh"),
			[]byte("#!/bin/bash\necho hello"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create script.sh: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "data.txt"),
			[]byte("data"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create data.txt: %v", err)
		}

		zipPath := filepath.Join(tmpDir, "out.zip")
		err := zipDirectory(srcDir, zipPath, []string{"script.sh"})
		if err != nil {
			t.Fatalf("zipDirectory failed: %v", err)
		}

		verifyZipPermissions(t, zipPath, "script.sh", 0o755)
		verifyZipPermissions(t, zipPath, "data.txt", 0o644)
	})

	t.Run("multiple exec files", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcDir := filepath.Join(tmpDir, "src")
		if err := os.MkdirAll(srcDir, 0o755); err != nil {
			t.Fatalf("failed to create src dir: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "bootstrap"),
			[]byte("#!/bin/bash"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create bootstrap: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "run.sh"),
			[]byte("#!/bin/bash"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create run.sh: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "lib.sh"),
			[]byte("#!/bin/bash"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create lib.sh: %v", err)
		}

		zipPath := filepath.Join(tmpDir, "out.zip")
		err := zipDirectory(srcDir, zipPath, []string{"bootstrap", "run.sh"})
		if err != nil {
			t.Fatalf("zipDirectory failed: %v", err)
		}

		verifyZipPermissions(t, zipPath, "bootstrap", 0o755)
		verifyZipPermissions(t, zipPath, "run.sh", 0o755)
		verifyZipPermissions(t, zipPath, "lib.sh", 0o644)
	})

	t.Run("default exec behavior", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcDir := filepath.Join(tmpDir, "src")
		if err := os.MkdirAll(srcDir, 0o755); err != nil {
			t.Fatalf("failed to create src dir: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "bootstrap"),
			[]byte("#!/bin/bash"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create bootstrap: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "run.sh"),
			[]byte("#!/bin/bash\r\n"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create run.sh: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "data.txt"),
			[]byte("data"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create data.txt: %v", err)
		}

		zipPath := filepath.Join(tmpDir, "out.zip")
		err := zipDirectory(srcDir, zipPath, []string{"bootstrap", "run.sh"})
		if err != nil {
			t.Fatalf("zipDirectory failed: %v", err)
		}

		verifyZipPermissions(t, zipPath, "bootstrap", 0o755)
		verifyZipPermissions(t, zipPath, "run.sh", 0o755)
		verifyZipPermissions(t, zipPath, "data.txt", 0o644)
	})

	t.Run("empty exec list", func(t *testing.T) {
		tmpDir := t.TempDir()
		srcDir := filepath.Join(tmpDir, "src")
		if err := os.MkdirAll(srcDir, 0o755); err != nil {
			t.Fatalf("failed to create src dir: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "bootstrap"),
			[]byte("#!/bin/bash"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create bootstrap: %v", err)
		}
		if err := os.WriteFile(
			filepath.Join(srcDir, "run.sh"),
			[]byte("#!/bin/bash"),
			0o644,
		); err != nil {
			t.Fatalf("failed to create run.sh: %v", err)
		}

		zipPath := filepath.Join(tmpDir, "out.zip")
		err := zipDirectory(srcDir, zipPath, []string{})
		if err != nil {
			t.Fatalf("zipDirectory failed: %v", err)
		}

		verifyZipPermissions(t, zipPath, "bootstrap", 0o644)
		verifyZipPermissions(t, zipPath, "run.sh", 0o644)
	})
}

func verifyZipPermissions(t *testing.T, zipPath, filename string, expectedMode os.FileMode) {
	t.Helper()
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		t.Fatalf("failed to open zip: %v", err)
	}
	defer func() {
		if err := r.Close(); err != nil {
			t.Errorf("error closing zip: %v", err)
		}
	}()

	for _, f := range r.File {
		if f.Name == filename {
			actualMode := f.Mode()
			unixAttrs := f.ExternalAttrs >> 16
			hasExec := actualMode&0o111 != 0
			expectedHasExec := expectedMode&0o111 != 0
			if hasExec != expectedHasExec {
				t.Errorf(
					"file %s: exec bit mismatch - mode=%o expected=%o unix_external_attrs=0x%04x (%o) hasExec=%v expectedHasExec=%v",
					filename,
					actualMode,
					expectedMode,
					unixAttrs,
					unixAttrs,
					hasExec,
					expectedHasExec,
				)
			}
			return
		}
	}
	t.Errorf("file %s not found in zip", filename)
}
