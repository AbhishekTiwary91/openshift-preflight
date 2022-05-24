package artifacts

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// WriteFile will write contents of the string to a file in
// the artifacts directory.
// Returns the full path (including the artifacts dir)
func WriteFile(filename, contents string) (string, error) {
	fullFilePath := filepath.Join(Path(), filename)

	err := os.WriteFile(fullFilePath, []byte(contents), 0o644)
	if err != nil {
		return fullFilePath, err
	}
	return fullFilePath, nil
}

func createArtifactsDir(artifactsDir string) (string, error) {
	if !strings.HasPrefix(artifactsDir, "/") {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Error(fmt.Errorf("unable to get current directory: %w", err))
			return "", err
		}

		artifactsDir = filepath.Join(currentDir, artifactsDir)
	}

	err := os.MkdirAll(artifactsDir, 0o777)
	if err != nil {
		log.Error(fmt.Errorf("unable to create artifactsDir: %w", err))
		return "", err
	}
	return artifactsDir, nil
}

// Path will return the artifacts path from viper config
func Path() string {
	artifactDir := viper.GetString("artifacts")
	artifactDir, err := createArtifactsDir(artifactDir)
	if err != nil {
		log.Fatal(fmt.Errorf("could not retrieve artifact path: %w", err))
		// Fatal does an os.Exit
	}
	return filepath.Join(artifactDir)
}
