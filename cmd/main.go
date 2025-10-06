package main

import (
	"flag"

	"github.com/vkotsiuba99/compose-to-nomad/pkg/compose"
	"github.com/vkotsiuba99/compose-to-nomad/pkg/nomad"
	"github.com/vkotsiuba99/compose-to-nomad/pkg/utils"
)

func main() {
	composeFilePath := flag.String("compose-file", "",
		"Path to the Docker Compose YAML file")
	outputDirPath := flag.String("output-dir", ".",
		"Directory to output the generated Nomad job files")
	flag.Parse()

	composeFile := compose.ReadComposeFile(*composeFilePath)

	tmpl := nomad.ParseTemplate()
	utils.CreateOutputDirectory(*outputDirPath)

	for name, service := range composeFile.Services {
		nomad.GenerateNomadJob(name, service, tmpl, *outputDirPath)
	}
}
