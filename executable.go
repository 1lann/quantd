package quantd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	_ "github.com/1lann/quantd/assets"

	"github.com/markbates/pkger"
)

type Executable struct {
	path string
}

func SetupExecutable(savePath ...string) (*Executable, error) {
	var file *os.File
	if len(savePath) == 0 {
		var err error
		file, err = ioutil.TempFile("", "quantd")
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		file, err = os.Create(savePath[0])
		if err != nil {
			return nil, err
		}
	}

	defer file.Close()

	f, err := pkger.Open("/assets/quantd_" + runtime.GOOS + "_" + runtime.GOARCH)
	if err != nil {
		return nil, fmt.Errorf("quantd: os and architecture unsupported: %w", err)
	}

	_, err = io.Copy(file, f)
	if err != nil {
		return nil, err
	}

	finalPath, err := filepath.Abs(file.Name())
	if err != nil {
		return nil, fmt.Errorf("quantd: could not determine path to executable: %w", err)
	}

	exe := &Executable{
		path: finalPath,
	}

	err = exe.test()
	if err != nil {
		return nil, fmt.Errorf("quantd: executable failed test, report this to 1lann: %w", err)
	}

	return exe, nil
}

func (e *Executable) test() error {
	out, err := exec.Command(e.path, "version").CombinedOutput()
	if err != nil {
		return err
	}

	if !bytes.HasPrefix(out, []byte("quantd")) {
		return fmt.Errorf("quantd: unexpected version response")
	}

	return nil
}
