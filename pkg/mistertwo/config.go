package mistertwo

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Configuration contains all the data from a configuration file.
type Configuration struct {
	Backup struct {
		Overwrite bool `yaml:"overwrite"`
	} `yaml:"backup"`
	Restore struct {
		Overwrite bool `yaml:"overwrite"`
	} `yaml:"restore"`
	Sync SyncOpts `yaml:"sync"`
	// Spec is a map of source to target, where source/target are
	// the path to a file or a directory.
	Spec map[string]string `yaml:"spec"`
}

// SyncOpts contains different options that change the sync task's behavior.
type SyncOpts struct {
	Clean     bool `yaml:"clean"`
	Overwrite bool `yaml:"overwrite"`
}

// NewConfiguration reads and validates a configuration file.
// Errors are written to os.Stderr and will result in program termination.
func NewConfiguration(path string) (config *Configuration) {
	path = strings.Replace(path, "~", os.Getenv("HOME"), 1)

	b, err := ioutil.ReadFile(path)
	fatalIfError(err, "could not load config file")

	err = yaml.Unmarshal(b, &config)
	fatalIfError(err, "could not parse config file")

	if !config.validate() {
		os.Exit(1)
	}

	return
}

// validate is a helper function that checks if a Configuration is valid.
func (config *Configuration) validate() (isValid bool) {
	isValid = true // until proven otherwise

	missing := func(str string) {
		fmt.Fprintf(os.Stderr, "bonclay: error: invalid config file: %s\n", str)
		isValid = false
	}

	if len(config.Spec) == 0 {
		missing("no files/directories specified in the config file's spec.")
	}

	return
}
