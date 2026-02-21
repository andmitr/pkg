package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/h2non/bimg"
)

// Resize creates a new image with the specified width and height.
//
// Stretches the image to the exact dimensions without preserving aspect ratio.
//
// Supported formats: jpg, jpeg, png, avif, webp.
func Resize(file, outDir string, width, height int) error {
	op := "Resize"

	bytes, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("%s.ReadFile: failed to read %q: %w", op, file, err)
	}

	img, err := bimg.NewImage(bytes).Process(bimg.Options{
		Width:  width,
		Height: height,
		Force:  true,
	})
	if err != nil {
		return fmt.Errorf("%s.NewImage.Process: failed to resize %q: %w", op, file, err)
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("%s: failed to create output directory %q: %w", op, outDir, err)
	}
	ext := filepath.Ext(file)
	outputPath := filepath.Join(
		outDir, strings.TrimSuffix(filepath.Base(file), ext)+"_"+strconv.Itoa(width)+"x"+strconv.Itoa(height)+ext,
	)
	if err := os.WriteFile(outputPath, img, 0644); err != nil {
		return fmt.Errorf("%s.WriteFile: failed to write resized %q: %w", op, file, err)
	}

	return nil
}
