package lang

import (
	"better-dev-container/cmd/lang/node"
	"github.com/docker/distribution/reference"
)

func MapLanguageCommand(containerImage string, inputCommand []string) []string {
	ref, err := reference.ParseNormalizedNamed(containerImage)
	if err != nil {
		panic(err)
	}

	switch reference.FamiliarName(ref) {
	case "node":
		return node.RunNpmCommand(inputCommand)
	default:
		return inputCommand
	}
}
