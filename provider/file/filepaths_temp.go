package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	coach_utils "github.com/CoachApplication/coach-utils"
)

const DELIM = "_"

/**
 * TempFilePaths a FilePaths implementation that maps scopes and keys to
 */
type TempFilePaths struct {
	uuid   string // A unique identifier for the files (like project name)
	suffix string // a file terminator which may make sense
}

// NewTempFilePaths TempFilePaths constructor
func NewTempFilePaths(uuid, suffix string) *TempFilePaths {
	return &TempFilePaths{
		uuid:   uuid,
		suffix: suffix,
	}
}

// Keys return a string list of available keys
func (tfp TempFilePaths) Keys() []string {
	keys := coach_utils.UniqueStringSlice{}

	files, err := ioutil.ReadDir(os.TempDir())
	if err == nil {
		for _, file := range files {
			if key, _, err := tfp.fileNameToKeyScope(file.Name()); err == nil {
				keys.Add(key)
			}
		}
	}

	return keys.Slice()
}

// Scopes return a string list of available scopes
func (tfp TempFilePaths) Scopes() []string {
	scopes := coach_utils.UniqueStringSlice{}

	files, err := ioutil.ReadDir(os.TempDir())
	if err == nil {
		for _, file := range files {
			if _, scope, err := tfp.fileNameToKeyScope(file.Name()); err == nil {
				scopes.Add(scope)
			}
		}
	}

	return scopes.Slice()
}

// Find the path for a key-scope pair, if they are valid
func (tfp TempFilePaths) Path(key, scope string) (string, error) {
	if name, err := tfp.fileNameFromKeyScope(key, scope); err != nil {
		return name, err
	} else {
		_, err := os.Stat(name)
		return name, err
	}
}

func (tfp TempFilePaths) fileNameToKeyScope(name string) (string, string, error) {
	if strings.HasPrefix(name, tfp.uuid+DELIM) && strings.HasSuffix(name, tfp.suffix) {
		name = strings.TrimPrefix(name, tfp.uuid+DELIM)
		name = strings.TrimSuffix(name, tfp.suffix)
		pieces := strings.Split(name, DELIM)

		if len(pieces) == 2 {
			key := pieces[0]
			scope := pieces[1]

			return key, scope, nil
		}
	}
	return "", "", errors.New("file is not a Config file")
}
func (tfp TempFilePaths) fileNameFromKeyScope(key, scope string) (string, error) {
	name := strings.Join([]string{tfp.uuid, key, scope}, DELIM) + tfp.suffix
	root := os.TempDir()
	return path.Join(root, name), nil
}
