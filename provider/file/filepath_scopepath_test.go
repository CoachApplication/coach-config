package file_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	file "github.com/CoachApplication/coach-config/provider/file"
)

func testScopedPaths() *file.ScopedFilePaths {
	workingDir, _ := os.Getwd()
	rootPathTest := path.Join(workingDir, "test")

	scopes := []string{"A", "B", "C"}
	scopedPaths := map[string]string{}
	for _, scope := range scopes {
		scopedPaths[scope] = path.Join(rootPathTest, scope)
	}

	return file.NewScopedFilePaths(scopes, scopedPaths, "", ".txt")
}

func TestScopedFilePaths_Keys(t *testing.T) {
	sfp := testScopedPaths()

	keys := sfp.Keys()

	if len(keys) != 4 {
		t.Error("ScopedFilePaths returned the wrong number of keys : ", keys)
	}
	if !sliceHasVal(keys, "one") {
		t.Error("ScopeFilePaths keys is missing expected value: one")
	}
	if !sliceHasVal(keys, "two") {
		t.Error("ScopeFilePaths keys is missing expected value: two")
	}
	if !sliceHasVal(keys, "three") {
		t.Error("ScopeFilePaths keys is missing expected value: three")
	}
	if !sliceHasVal(keys, "four") {
		t.Error("ScopeFilePaths keys is missing expected value: four")
	}
}

func TestScopedFilePaths_Scopes(t *testing.T) {
	sfp := testScopedPaths()

	scopes := sfp.Scopes()
	if len(scopes) != 3 {
		t.Error("ScopedFilePaths returned the wrong number of scopes: ", scopes)
	}
	if !sliceHasVal(scopes, "A") {
		t.Error("ScopeFilePaths keys is missing expected scope: A")
	}
	if !sliceHasVal(scopes, "B") {
		t.Error("ScopeFilePaths keys is missing expected scope: B")
	}
	if !sliceHasVal(scopes, "C") {
		t.Error("ScopeFilePaths keys is missing expected scope: C")
	}
}

func TestScopedFilePaths_Path(t *testing.T) {
	sfp := testScopedPaths()

	if p, err := sfp.Path("no", "no"); err == nil {
		t.Error("ScopedPaths returned no error making a real path for an invalid key-scope pair: ", "no", "no", p, err.Error())
	}
	if p, err := sfp.Path("one", "A"); err != nil {
		t.Error("ScopedPaths returned an error making a real path: ", "A", "one", p, err.Error())
	} else {
		if contents, err := ioutil.ReadFile(p); err != nil {
			t.Error("Could not get content from scoped key file :", err.Error())
		} else if string(contents) != "A-one" {
			t.Error("Wrong content: ", p, string(contents))
		}
	}
}

func sliceHasVal(s []string, v string) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}
	return false
}
