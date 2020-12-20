package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"runtime/debug"

	"github.com/1lann/imagequant"
	quantdimage "github.com/1lann/quantd/image"
	"github.com/tinylib/msgp/msgp"
)

func usage() {
	fmt.Println("quantd is a standalone helper binary for programs wishing to use libimagequant without linking to C.")
	fmt.Println("quantd when built contains libimagequant bindings which contains GPLv3 licensed code.")
	fmt.Println("View the license with `quantd license`")
	fmt.Println("Refer to: https://github.com/1lann/quantd")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "license":
		fmt.Println(imagequant.License())
		os.Exit(0)
	case "info":

	}

	if os.Args[1] == "version" {
		buildInfo, ok := debug.ReadBuildInfo()

		quantdVersion := "unknown version"
		imagequantVersion := "unknown version"

		if ok {
			if buildInfo.Main.Sum != "" {
				quantdVersion = buildInfo.Main.Version + " (" + buildInfo.Main.Sum + ")"
			} else {
				quantdVersion = buildInfo.Main.Version
			}

			for _, dep := range buildInfo.Deps {
				switch dep.Path {
				case "github.com/1lann/quantd/image":
					quantdVersion = dep.Version + " (" + dep.Sum + ")"
				case "github.com/1lann/imagequant":
					imagequantVersion = dep.Version + " (" + dep.Sum + ")"
				}
			}
		}

		fmt.Println("quantd " + quantdVersion)
		fmt.Println("with github.com/1lann/imagequant " + imagequantVersion)
		os.Exit(0)
	}

	if os.Args[1] != "start" {
		fmt.Println("Unknown subcommand: " + os.Args[1])
		usage()
		os.Exit(1)
	}

	in := msgp.NewReader(os.Stdin)
	out := msgp.NewWriter(os.Stdout)
	var req quantdimage.ImageRequest
	for {
		err := req.DecodeMsg(in)
		if err != nil {
			log.Fatalf("error decoding request: %v", err)
		}

		resp, err := Quantize(&req)
		if err != nil {
			log.Fatalf("error quantizing image: %v", err)
		}

		err = resp.EncodeMsg(out)
		if err != nil {
			log.Fatalf("error encoding response: %v", err)
		}
		err = out.Flush()
		if err != nil {
			log.Fatalf("error flushing response: %v", err)
		}
	}
}

func toRawPalette(palette color.Palette) []quantdimage.RGBA {
	result := make([]quantdimage.RGBA, len(palette))
	for i, col := range palette {
		result[i] = quantdimage.RGBA(col.(color.RGBA))
	}

	return result
}

// Quantize quantizes an image into a maximum of 16 colors with the given
// parameters.
func Quantize(req *quantdimage.ImageRequest) (*quantdimage.ImageResponse, error) {
	attr, err := getAttributes(req.Speed, req.MaxColors)
	if err != nil {
		return nil, fmt.Errorf("Attribues: %w", err)
	}
	defer attr.Release()

	quant, err := imagequant.NewImage(attr, req.Image, req.DX, req.DY, 0)
	if err != nil {
		return nil, fmt.Errorf("ref NewImage: %w", err)
	}
	defer quant.Release()

	res, err := quant.Quantize(attr)
	if err != nil {
		return nil, fmt.Errorf("Quantize: %w", err)
	}

	defer res.Release()

	// outputImg, err := imagequant.NewImage(attr, imagequant.GoImageToRgba32(img),
	// 	img.Bounds().Dx(), img.Bounds().Dy(), 0)
	// if err != nil {
	// 	return nil, fmt.Errorf("img NewImage: %w", err)
	// }
	// defer outputImg.Release()

	// res.SetOutputImage(outputImg)

	err = res.SetDitheringLevel(float32(req.Dither))
	if err != nil {
		return nil, fmt.Errorf("SetDitheringLevel: %w", err)
	}

	rgb8data, err := res.WriteRemappedImage()
	if err != nil {
		return nil, fmt.Errorf("WriteRemappedImage: %w", err)
	}

	return &quantdimage.ImageResponse{
		Image:   rgb8data,
		DX:      req.DX,
		DY:      req.DY,
		Palette: toRawPalette(res.GetPalette()),
	}, nil
}

func getAttributes(speed, maxColors int) (*imagequant.Attributes, error) {
	attr, err := imagequant.NewAttributes()
	if err != nil {
		return nil, fmt.Errorf("NewAttributes: %w", err)
	}

	err = attr.SetSpeed(speed)
	if err != nil {
		return nil, fmt.Errorf("SetSpeed: %w", err)
	}

	err = attr.SetMaxColors(maxColors)
	if err != nil {
		return nil, fmt.Errorf("SetMaxColors: %w", err)
	}

	return attr, nil
}
