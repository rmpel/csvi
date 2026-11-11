package csviapp

import (
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRcFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "csvirc")
	content := "# comment\n\n-rv\nautofit\n--w 20,0:5\n  exteditor  code --wait  \n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	args, err := readRcArgs(path)
	if err != nil {
		t.Fatal(err)
	}
	expect := []string{"-rv", "-autofit", "-w=20,0:5", "-exteditor=code --wait"}
	if !reflect.DeepEqual(args, expect) {
		t.Fatalf("expected %#v, got %#v", expect, args)
	}

	t.Setenv("CSVIRC", path)
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	opt := NewOptions().Bind(fs)
	if err := applyRcFile(fs); err != nil {
		t.Fatal(err)
	}
	// the command line overrides the rc file
	if err := fs.Parse([]string{"-rv=false", "file.csv"}); err != nil {
		t.Fatal(err)
	}
	if opt.ReverseVideo || !opt.AutoFit || opt.CellWidth != "20,0:5" || opt.ExtEditor != "code --wait" {
		t.Fatalf("unexpected options: %+v", opt)
	}
	if fs.Arg(0) != "file.csv" {
		t.Fatalf("expected file.csv as argument, got %q", fs.Arg(0))
	}
}

func TestRcFileMissing(t *testing.T) {
	args, err := readRcArgs(filepath.Join(t.TempDir(), "nope"))
	if err != nil || args != nil {
		t.Fatalf("missing rc file should be ignored, got %v, %v", args, err)
	}
}
