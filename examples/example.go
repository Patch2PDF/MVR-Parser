package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	GDTFMeshReader "github.com/Patch2PDF/GDTF-Mesh-Reader"
	STL "github.com/Patch2PDF/GDTF-Parser/examples/stl"
	MVRParser "github.com/Patch2PDF/MVR-Parser"
	MVRTypes "github.com/Patch2PDF/MVR-Parser/pkg/types"
)

var config = MVRTypes.MVRParserConfig{
	MeshHandling:      MVRTypes.BuildStageModel,
	ReadThumbnail:     true,
	GDTFParserWorkers: 4,
	StageMeshWorkers:  4,
	ModelConfig: MVRTypes.ModelConfig{
		Global: MVRTypes.GlobalModelConfig{
			RenderOnlyAddressedFixture: true,
		},
		Individual: map[string]MVRTypes.ModelNodeConfig{
			"FA992217-CB18-D844-9D42-5B791B2BF05E": {
				Exclude:                    &MVRTypes.FalsePtr,
				RenderOnlyAddressedFixture: &MVRTypes.TruePtr,
			},
		},
	},
}

func main() {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	start := time.Now()

	mvr, err := zip.OpenReader("test.mvr")
	if err != nil {
		log.Fatal(err)
	}
	defer mvr.Close()

	GDTFMeshReader.LoadPrimitives()

	mvrData, err := MVRParser.ParseMVRZipReader(&mvr.Reader, config)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v\n", mvrData)
	}
	fmt.Println("Elapsed:", time.Since(start))
	pprof.StopCPUProfile()

	// write mesh as STL
	meshFile, _ := os.Create("Test.stl")
	STL.WriteBinary(meshFile, mvrData.StageModel)
	f.Close()
}
