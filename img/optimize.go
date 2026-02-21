package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"time"

	"github.com/h2non/bimg"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/svg"
	"golang.org/x/sync/errgroup"
)

// OptimizeAll optimizes all images in inDir and saves them to outDir.
//
// Applies the Optimize function to each image.
func OptimizeAll(inDir string, outDir string) error {
	op := "OptimizeAll"

	defer func(startTime time.Time) {
		fmt.Printf("%s: Execution time: %v\n", op, time.Since(startTime).Round(time.Millisecond))
	}(time.Now())

	info, err := os.Stat(inDir)
	switch {
	case err != nil:
		return fmt.Errorf("%s: %w", op, err)
	case !info.IsDir():
		return fmt.Errorf("%s.Stat: input path is not directory", op)
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("%s.MkdirAll: failed to create output directory: %w", op, err)
	}

	files, err := os.ReadDir(inDir)
	if err != nil {
		return fmt.Errorf("%s.ReadDir: failed to read directory %q: %w", op, inDir, err)
	}

	var g errgroup.Group
	g.SetLimit(runtime.NumCPU())

	var validExts = []string{".jpg", ".jpeg", ".png", ".avif", ".webp", ".svg"}
	for _, file := range files {
		g.Go(func() error {
			filename := file.Name()
			ext := strings.ToLower(filepath.Ext(filename))

			if !slices.Contains(validExts, ext) {
				return nil
			}

			if err := Optimize(filepath.Join(inDir, filename), outDir); err != nil {
				return fmt.Errorf("%s: failed to optimize %q: %w", op, filename, err)
			}

			return nil
		})
	}

	return g.Wait()
}

// Optimize creates a new compressed image named file_min.ext in outDir.
//
// Parameters are optimized for web and provide a balance between output image quality and processing speed.
//
// Supported formats: jpg, jpeg, png, avif, webp, svg.
func Optimize(file string, outDir string) error {
	op := "Optimize"

	bytes, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("%s.ReadFile: failed to read file %q: %w", op, file, err)
	}

	imageType := bimg.DetermineImageType(bytes)
	ext := strings.ToLower(filepath.Ext(file))

	if imageType == bimg.UNKNOWN {
		switch ext {
		case ".jpg", ".jpeg":
			imageType = bimg.JPEG
		case ".png":
			imageType = bimg.PNG
		case ".webp":
			imageType = bimg.WEBP
		case ".avif":
			imageType = bimg.AVIF
		case ".svg":
			imageType = bimg.SVG
		default:
			return fmt.Errorf("%s: unknown image type", op)
		}
	}

	var options = map[bimg.ImageType]bimg.Options{
		bimg.JPEG: {
			Quality:       70,
			StripMetadata: true,
			Interlace:     true,
		},
		bimg.PNG: {
			Compression: 6,
			// Remove Palette and Quality if you need good semi-transparency
			Quality:       85,   // only used with Palette, controls color quantization quality
			Palette:       true, // strong compression but reduces transparency to 1-bit, use high Quality (85-95) with it
			StripMetadata: true,
			Interlace:     true,
		},
		bimg.WEBP: {
			Quality:       80,
			StripMetadata: true,
		},
		bimg.AVIF: {
			Quality:       70,
			Speed:         7, // lower Speed = better compression but slower
			StripMetadata: true,
		},
	}
	var result []byte
	switch imageType {
	case bimg.SVG:
		m := minify.New()
		m.AddFunc("image/svg+xml", svg.Minify)
		m.AddFunc("text/css", css.Minify)
		if result, err = m.Bytes("image/svg+xml", bytes); err != nil {
			return fmt.Errorf("%s.svg.Minify: failed to process %q: %w", op, file, err)
		}
	default:
		if result, err = bimg.NewImage(bytes).Process(options[imageType]); err != nil {
			return fmt.Errorf("%s.Process: failed to process %q: %w", op, file, err)
		}
	}
	if len(result) >= len(bytes) {
		result = bytes
		log.Printf("Info: %q is already optimal, saving original\n", file)
	}

	outputPath := filepath.Join(outDir, strings.TrimSuffix(filepath.Base(file), ext)+"_min"+ext)
	if err := os.WriteFile(outputPath, result, 0644); err != nil {
		return fmt.Errorf("%s.WriteFile: failed to write result file for %q: %w", op, file, err)
	}

	return nil
}
