package repository

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	MaxUploadBytes  = 50 << 20  // 50 MB compressed
	MaxExtractBytes = 200 << 20 // 200 MB uncompressed (zip-bomb guard)
)

var (
	ErrFileTooLarge = errors.New("file exceeds 50 MB limit")
	ErrInvalidMIME  = errors.New("file must be a ZIP archive")
)

// zipMagic is the ZIP local file header signature.
var zipMagic = []byte{0x50, 0x4B, 0x03, 0x04}

func validateZipHeader(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := make([]byte, 4)
	if _, err := io.ReadFull(f, buf); err != nil {
		return ErrInvalidMIME
	}
	for i, b := range zipMagic {
		if buf[i] != b {
			return ErrInvalidMIME
		}
	}
	return nil
}

func extractZip(src, destDir string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("open zip: %w", err)
	}
	defer r.Close()

	var totalBytes int64
	for _, f := range r.File {
		target, err := safeJoin(destDir, f.Name)
		if err != nil {
			return err
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(target, 0755)
			continue
		}
		n, err := writeZipFile(f, target)
		if err != nil {
			return err
		}
		totalBytes += n
		if totalBytes > MaxExtractBytes {
			return fmt.Errorf("extracted content exceeds %d MB limit", MaxExtractBytes>>20)
		}
	}
	return nil
}

// safeJoin prevents path traversal by ensuring target stays within destDir.
func safeJoin(base, name string) (string, error) {
	target := filepath.Join(base, filepath.Clean("/"+name))
	if !strings.HasPrefix(target, filepath.Clean(base)+string(os.PathSeparator)) {
		return "", fmt.Errorf("invalid path in zip: %s", name)
	}
	return target, nil
}

func writeZipFile(f *zip.File, target string) (int64, error) {
	if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
		return 0, err
	}
	out, err := os.Create(target)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	rc, err := f.Open()
	if err != nil {
		return 0, err
	}
	defer rc.Close()
	// LimitedReader guards against zip-bomb per-file as well.
	n, err := io.Copy(out, io.LimitReader(rc, MaxExtractBytes))
	return n, err
}
