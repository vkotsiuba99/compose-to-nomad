package compose

import (
	"log"
	"os"

	"github.com/vkotsiuba99/compose-to-nomad/pkg/types"
	"gopkg.in/yaml.v2"
)

func ReadComposeFile(path string) types.ComposeFile {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read Docker Compose file: %v", err)
	}

	var composeFile types.ComposeFile
	err = yaml.Unmarshal(data, &composeFile)
	if err != nil {
		log.Fatalf("Failed to unmarshal Docker Compose YAML: %v", err)
	}
	return composeFile
}
