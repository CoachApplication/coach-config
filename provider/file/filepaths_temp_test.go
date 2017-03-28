package file

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func makeSomeTempFiles(t *testing.T, p *TempFilePaths) chan bool {
	names := []string{}
	for _, key := range []string{"A", "B", "C"} {
		for _, scope := range []string{"1", "2", "3"} {
			path, _ := p.Path(key, scope)

			if f, err := os.Create(path); err == nil {
				names = append(names, f.Name())
				t.Logf("Making temp file : %s", f.Name())
				f.WriteString(strings.Join([]string{key, scope}, DELIM))
				f.Close()
			}

		}
	}

	remove := make(chan bool)
	go func(names []string, remove chan bool) {
		<-remove
		for _, name := range names {
			t.Logf("Removing temp file : %s", name)
			os.Remove(name)
		}
	}(names, remove)

	time.Sleep(1000) // just give the OS some time to make the files
	return remove
}

func TestTempFilePaths_Keys(t *testing.T) {
	ps := NewTempFilePaths("coach-base-filepaths-build", "_keys")
	rem := makeSomeTempFiles(t, ps)

	keys := ps.Keys()
	if len(keys) != 3 {
		t.Error("TempFilePath returned the wrong number of Keys()")
	}

	rem <- true
}

func TestTempFilePaths_Scopes(t *testing.T) {
	ps := NewTempFilePaths("coach-base-filepaths-build", "_scopes")
	rem := makeSomeTempFiles(t, ps)

	scopes := ps.Scopes()
	if len(scopes) != 3 {
		t.Error("TempFilePath returned the wrong number of Scopes() : ", scopes)
	}

	rem <- true
}

func TestTempFilePaths_Path(t *testing.T) {
	ps := NewTempFilePaths("coach-base-filepaths-build", "_path")
	rem := makeSomeTempFiles(t, ps)

	if _, err := ps.Path("no", "no"); err == nil {
		t.Error("TempFilePaths return no error on an invalid key-scope pair")
	}

	if p, err := ps.Path("A", "1"); err != nil {
		t.Error("TempFilePaths return an error on a key-scope pair: ", err.Error())
	} else if contents, err := ioutil.ReadFile(p); err != nil {
		t.Error("TempFilePaths provided a filename that couldn't be opened: ", err.Error())
	} else if string(contents) != strings.Join([]string{"A", "1"}, DELIM) {
		t.Error("TempFilePath given path had the wrong content", string(contents))
	}

	rem <- true
}
