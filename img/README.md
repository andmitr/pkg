# img

Functions for optimizing, resizing, and converting images using libvips.

[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square;logo=opensource)](../LICENSE)

## Sponsorship

[![Boosty](https://img.shields.io/badge/Boosty-F15F2C?style=for-the-badge&logo=boosty&logoColor=white)![Support](https://img.shields.io/badge/Support%20me-grey?style=for-the-badge)](https://boosty.to/andmitr/donate) 

![Bitcoin](https://img.shields.io/badge/Bitcoin-F7931A?style=flat&logo=bitcoin&logoColor=white&logoSize=auto) 
```
1CCnwAvJYEoDVGM7vsBg2Q99cF9EHtBVaY
```

![Tether](https://img.shields.io/badge/Tether%20(USDT%20ETH)-168363?style=flat&logo=tether&logoColor=white&logoSize=auto) 
```
0x54f0ccc6b2987de454f69f2814fc9202bcfb74fe
```

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Supported Formats](#supported-formats)
- [Performance Notes](#performance-notes)
- [License](#license)

## Features

- Parallel processing for batch operations 
- Automatic format detection based on file content
- Preserves original if optimization does not reduce file size
- Stripped metadata 
- Compress images with format-specific settings tuned for web

## Prerequisites

- Go 1.25 or later
- libvips 

**Arch Linux**:
```bash
sudo pacman -S libvips
```

**macOS**:
```bash
brew install vips
```

**Ubuntu/Debian**:
```bash
sudo apt install libvips-dev
```

See [libvips install instructions](https://www.libvips.org/install.html) for other platforms.

## Installation

```bash
go get github.com/andmitr/pkg/img
```

## Usage

### func Optimize

```go
func Optimize(file string, outDir string) error
```

Optimize compresses a single image and saves it as `filename_min.ext` in outDir.

Parameters:
- `file` - path to input image
- `outDir` - directory for output file

Returns error if file cannot be read or processed.

### Example
```go
err := img.Optimize("inputDir/photo.jpg", "outputDir")
// Creates: outputDir/photo_min.jpg
```

The function applies format-specific compression:
- JPEG: quality 70, interlaced
- PNG: compression 6, palette mode, quality 85
- WebP: quality 80
- AVIF: quality 70, speed 7
- SVG: minified via tdewolff/minify

**Note**  
Compression options are hardcoded for optimal web performance.
If you need custom settings, fork the repository and modify the options in `optimize.go`.

### func OptimizeAll

```go
func OptimizeAll(inDir string, outDir string) error
```

OptimizeAll processes all supported images in inDir and saves optimized versions to outDir.

Processes files in parallel using all available CPU cores.
Skips files with unsupported extensions.

#### Example

```go
err := img.OptimizeAll("inputDir", "outputDir")
// Processes all supported images in parallel
// Creates: outputDir/photo1_min.jpg, outputDir/photo2_min.png, etc.
```

### func Resize

```go
func Resize(file string, outDir string, width int, height int) error
```

Resize creates a new image stretched to the specified dimensions.

Parameters:
- `file` - path to input image
- `outDir` - directory for output file
- `width` - target width in pixels
- `height` - target height in pixels

#### Example
```go
err := img.Resize("input/photo.jpg", "output", 800, 600)
// Creates: output/photo_800x600.jpg
// Stretches the image to exactly 800x600 pixels.
```

### func Convert

```go
func Convert(file string, outDir string, outType string) error
```

Convert transforms an image to a different format.

Parameters:
- `file` - path to input image
- `outDir` - directory for output file
- `outType` - target format: jpg, jpeg, png, webp, avif

#### Example
```go
err := img.Convert("input/photo.png", "output", "webp")
// Creates: output/photo.webp
```

## Supported Formats

| Format | Optimize | Resize | Convert |
|--------|----------|--------|---------|
| JPEG   | Yes      | Yes    | Yes     |
| PNG    | Yes      | Yes    | Yes     |
| WebP   | Yes      | Yes    | Yes     |
| AVIF   | Yes      | Yes    | Yes     |
| SVG    | Yes      | No     | No      |

## Performance Notes

- OptimizeAll uses goroutine pooling with a limit of `runtime.NumCPU()` to prevent excessive memory usage
- Processing order is non-deterministic due to parallel execution
- AVIF encoding is slower but provides excellent compression
- SVG optimization uses a different pipeline (minify) rather than libvips

## License
MIT Licensed. See [LICENSE](../LICENSE) for details.
