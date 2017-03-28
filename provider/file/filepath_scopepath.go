package file

import (
	"errors"
	"io/ioutil"
	"path"
	"strings"

	coach_utils "github.com/james-nesbitt/coach-utils"
)

/**
 * Scoped paths
 */

// FilePaths FilePaths that uses the an arrangement of files where different scopes are in different folders
type ScopedFilePaths struct {
	scopes       []string
	scopePathMap map[string]string

	prefix string
	suffix string
}

// NewFilePaths FilePaths constructor
func NewScopedFilePaths(scopes []string, scopePathMap map[string]string, prefix, suffix string) *ScopedFilePaths {
	return &ScopedFilePaths{
		scopes:       scopes,
		scopePathMap: scopePathMap,
		prefix:       prefix,
		suffix:       suffix,
	}
}

func (sfp *ScopedFilePaths) FilePaths() FilePaths {
	return FilePaths(sfp)
}

// Keys return a list of all keys found (*this is not reliable as it may change easily after scanning)
func (sfp *ScopedFilePaths) Keys() []string {
	keys := coach_utils.UniqueStringSlice{}

	for _, scope := range sfp.scopes {
		if sp, found := sfp.scopePathMap[scope]; found {
			if ps, err := ioutil.ReadDir(sp); err == nil {
				for _, p := range ps {
					n := p.Name()

					if sfp.prefix != "" {
						if !strings.HasPrefix(n, sfp.prefix) {
							continue
						}
						n = strings.TrimPrefix(n, sfp.prefix)
					}
					if sfp.suffix != "" {
						if !strings.HasSuffix(n, sfp.suffix) {
							continue
						}
						n = strings.TrimSuffix(n, sfp.suffix)
					}
					keys.Add(n)
				}
			}
		}
	}

	return keys.Slice()
}

func (sfp *ScopedFilePaths) Scopes() []string {
	return sfp.scopes
}

func (sfp *ScopedFilePaths) Path(key, scope string) (string, error) {
	if p, found := sfp.scopePathMap[scope]; found {
		n := sfp.prefix + key + sfp.suffix
		return path.Join(p, n), nil
	}
	return "", errors.New("Invalid scope :" + scope)
}
