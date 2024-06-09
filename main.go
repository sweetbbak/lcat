package main

import (
	"log"
	"os"
	"github.com/sweetbbak/lcat/icat"

	"github.com/jessevdk/go-flags"
)

type Opts struct {
	Align              string  `short:"a" long:"align" description:"Horizontal alignment for the displayed image. Choices: center, left, right"`
	Place              string  `short:"p" long:"place" description:"Choose where on the screen to display the image. The image will be scaled to fit into the specified rectangle. The syntax for specifying rectangles is <width>x<height>@<left>x<top>."`
	ScaleUp            bool    `short:"S" long:"scale" description:"scale up images if they are smaller than the given area specified by --place"`
	Background         string  `short:"b" long:"background" description:"add a color background to an image"`
	Mirror             string  `short:"m" long:"mirror" description:"mirror imagae on X or Y axis. Choices none, both, horizontal, vertical"`
	TransferMode       string  `short:"t" long:"transfer-mode" description:"choices: detect, file, memory, stream"`
	Clear              bool    `short:"c" long:"clear" description:"clear all images placed"`
	DetectSupport      bool    `short:"D" long:"detect-support" description:"detect terminal support"`
	DetectionTimeout   float64 `short:"T" long:"detection-timeout" description:"detection timeout for detect support flag"`
	UseWindowSize      string  `short:"w" long:"win-size" description:"set window size instead of detecting it. Format: width_in_cells,height_in_cells,width_in_pixels,height_in_pixels"`
	PrintWindowSize    bool    `short:"x" long:"print-window-size" description:"print the window size"`
	Stdin              string  `short:"s" long:"stdin" description:"read image data from stdin - yes, no, detect"`
	Engine             string  `short:"e" long:"engine" description:"image engine - auto, builtin, magick"`
	ZIndex             string  `short:"z" long:"z-index" description:"z index of the image"`
	Loop               int     `short:"l" long:"loop" description:"loop GIF N number of times [-1 for infinite]"`
	Hold               bool    `short:"H" long:"hold" description:"wait for a keypress before exiting"`
	UnicodePlaceholder bool    `short:"U" long:"unicode" description:"use the unicode placeholder for images"`
	Passthrough        string  `short:"g" long:"passthrough" description:"terminal passthrough - detect, none, tmux"`
	ImageId            int     `short:"I" long:"id" description:"set the image ID"`
}

func cat(args []string, opt Opts) error {
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
	return nil
}

var opts Opts

func init() {
	opts.TransferMode = "memory"
	opts.Align = "center"
	opts.ZIndex = "0"
}

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}

	err = cat(args, opts)
	if err != nil {
		log.Fatal(err)
	}
}
