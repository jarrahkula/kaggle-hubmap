package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

var (
// DataPath      string = "./input/hubmap-kidney-segmentation"
// TrainPath     string = "./input/hubmap-kidney-segmentation/train"
// TestPath      string = "./input/hubmap-kidney-segmentation/test"
// MaskPath      string = "./input/mask"
// TileImagePath string = "./input/tile/image"
// TestImagePath string = "./input/test/image"
// TileMaskPath  string = "./input/tile/mask"
// TestMaskPath  string = "./input/test/mask"
// Reduction     int    = 4
// TileSize      int    = 256
)

// flag variables
var (
	DataPath  string
	OptStr    string
	ModelPath string
	Device    string
	task      string
)

// hyperparameters
var (
	Reduction    int     // image resolution reduction times
	TileSize     int     // image tile size
	LR           float64 // learning rate
	BatchSize    int     // batch size
	ValidateSize int     // number of iterations that triggers running validation task
)

func init() {
	flag.StringVar(&DataPath, "input", "./input", "specify input data directory")
	flag.StringVar(&ModelPath, "model", "./model/resnet34.ot", "specify full path to model weight '.ot' file.")
	flag.StringVar(&Device, "device", "CPU", "specify device i.e. CPU or GPU to run.")
	flag.StringVar(&task, "task", "train", "specify task to run")
	flag.Float64Var(&LR, "lr", 0.001, "specify learning rate")
	flag.IntVar(&Reduction, "reduction", 4, "specify image resolution reduction times")
	flag.IntVar(&BatchSize, "batch", 16, "specify batch size")
	flag.IntVar(&ValidateSize, "validate", 10, "specify size of validation cycles.")
	flag.IntVar(&TileSize, "tile", 256, "specify tile image size")
	flag.StringVar(&OptStr, "opt", "SGD", "specify optimizer type")
}

func main() {
	flag.Parse()

	DataPath = absPath(DataPath)
	ModelPath = absPath(ModelPath)

	switch task {
	case "train":
		runTrain()
	case "eda":
		runEDA()
	case "image":
		processImage()
	default:
		err := fmt.Errorf("Unknown 'task' name. Please specify valid 'task' flag to run.\n")
		panic(err)
	}
}

// helper to get absolute file path
func absPath(p string) string {
	fullpath, err := filepath.Abs(p)
	if err != nil {
		log.Fatal(err)
	}
	return fullpath
}
