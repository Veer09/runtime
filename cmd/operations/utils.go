package operations

import (
	"encoding/json"
	"os"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func readSpec() (*specs.Spec, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	spec := &specs.Spec{}
	if err := json.NewDecoder(f).Decode(spec); err != nil {
		return nil, err
	}
	return spec, nil
}