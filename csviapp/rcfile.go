package csviapp

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// rcFileName is the per-user preference file in the home directory.
// Each non-empty line holds one command-line option, which is applied
// as a default before the actual command line is parsed:
//
//	# comment
//	-rv
//	-autofit
//	-w 20
//	exteditor code --wait
//
// The leading "-" is optional. Everything after the first run of
// whitespace is the option's value, so values may contain spaces.
const rcFileName = ".csvirc"

func rcFilePath() string {
	if p := os.Getenv("CSVIRC"); p != "" {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, rcFileName)
}

func readRcArgs(path string) ([]string, error) {
	fd, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	defer fd.Close()

	var args []string
	sc := bufio.NewScanner(fd)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || line[0] == '#' {
			continue
		}
		name, value, hasValue := strings.Cut(line, " ")
		if !hasValue {
			name, value, hasValue = strings.Cut(line, "\t")
		}
		name = "-" + strings.TrimLeft(name, "-")
		if hasValue {
			args = append(args, name+"="+strings.TrimSpace(value))
		} else {
			args = append(args, name)
		}
	}
	return args, sc.Err()
}

// applyRcFile parses the options in the rc file into fs so that they
// become defaults which the real command line can still override.
func applyRcFile(fs *flag.FlagSet) error {
	path := rcFilePath()
	if path == "" {
		return nil
	}
	args, err := readRcArgs(path)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	if len(args) == 0 {
		return nil
	}
	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("%s: unexpected argument %q", path, fs.Arg(0))
	}
	return nil
}
