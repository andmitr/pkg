package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/h2non/bimg"
)

// Convert converts the source image to the specified format.
//
// Supported outType values: jpg, jpeg, png, avif, webp.
func Convert(file, outDir, outType string) error {
	op := "Convert"

	validTypes := []string{"jpg", "jpeg", "png", "avif", "webp"}
	if !slices.Contains(validTypes, outType) {
		return fmt.Errorf("%s: unsupported output type %q", op, outType)
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("%s.ReadFile: failed to read file %q: %w", op, file, err)
	}

	var options bimg.Options
	switch outType {
	case "png":
		options = bimg.Options{Type: bimg.PNG}
	case "jpg", "jpeg":
		options = bimg.Options{Type: bimg.JPEG}
	case "webp":
		options = bimg.Options{Type: bimg.WEBP}
	case "avif":
		options = bimg.Options{Type: bimg.AVIF}
	}

	result, err := bimg.NewImage(bytes).Process(options)
	if err != nil {
		return fmt.Errorf("%s.NewImage.Process: failed to convert %q: %w", op, file, err)
	}

	outName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)) + "." + strings.ToLower(outType)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("%s.MkdirAll: failed to create output directory %q: %w", op, outDir, err)
	}
	if err := os.WriteFile(filepath.Join(outDir, outName), result, 0644); err != nil {
		return fmt.Errorf("%s.WriteFile: failed to write converted %q: %w", op, file, err)
	}

	return nil
}
