# lcat

A Kitty Terminal graphics CLI tool and Library.

## The lcat CLI tool

if you are familiar with `icat` then the options are basically the exact same.

```sh
lcat --align=left img.png
lcat --place=10x10@0x0 img.png
```

## As a Library

```go

package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/sweetbbak/lcat/icat"
)

func lcat(args []string, opt Opts) error {
	var optss = &icat.Options{
		Align:              opt.Align,
		Place:              opt.Place,
		ScaleUp:            opt.ScaleUp,
		Background:         opt.Background,
		Mirror:             opt.Mirror,
		TransferMode:       opt.TransferMode,
		Clear:              opt.Clear,
		DetectSupport:      opt.DetectSupport,
		DetectionTimeout:   opt.DetectionTimeout,
		UseWindowSize:      opt.UseWindowSize,
		PrintWindowSize:    opt.PrintWindowSize,
		Stdin:              opt.Stdin,
		Engine:             opt.Engine,
		ZIndex:             opt.ZIndex,
		Loop:               opt.Loop,
		Hold:               opt.Hold,
		UnicodePlaceholder: opt.UnicodePlaceholder,
		Passthrough:        opt.Passthrough,
		ImageId:            opt.ImageId,
	}

	_, err := icat.Icat(optss, args)
	if err != nil {
		return err
	}
}
```
