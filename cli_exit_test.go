package fountain

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// These tests run the three programs as built binaries, the way a user does,
// with real files from testdata.

var programNames = []string{"fountain2html", "fountain2json", "fountainfmt"}

var builtDir string

// binaries builds the three programs once per test run.
func binaries(t *testing.T) string {
	t.Helper()
	if builtDir != "" {
		return builtDir
	}
	dir, err := os.MkdirTemp("", "fountain-cli-")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range programNames {
		if out, err := exec.Command("go", "build", "-o", filepath.Join(dir, name), "./cmd/"+name).CombinedOutput(); err != nil {
			t.Fatalf("%s: build failed: %v\n%s", name, err, out)
		}
	}
	builtDir = dir
	return dir
}

func TestMain(m *testing.M) {
	code := m.Run()
	if builtDir != "" {
		os.RemoveAll(builtDir)
	}
	os.Exit(code)
}

// run runs a built program with stdin and returns stdout, stderr and the exit code.
func run(t *testing.T, dir, name, stdin string, args ...string) (string, string, int) {
	t.Helper()
	cmd := exec.Command(filepath.Join(binaries(t), name), args...)
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(stdin)
	var out, errOut bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errOut
	code := 0
	if err := cmd.Run(); err != nil {
		var ee *exec.ExitError
		if !errors.As(err, &ee) {
			t.Fatalf("%s: %v", name, err)
		}
		code = ee.ExitCode()
	}
	return out.String(), errOut.String(), code
}

func readFile(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

/** TestOutputFileHoldsTheResult is the bug found while checking exit codes:
 * fountain2html declared a new `out` inside the -o branch, so the output file
 * was created empty and the result went to standard output.
 */
func TestOutputFileHoldsTheResult(t *testing.T) {
	in, _ := filepath.Abs("testdata/sample-01.fountain")
	for _, name := range programNames {
		want, _, code := run(t, t.TempDir(), name, "", "-i", in)
		if code != 0 || want == "" {
			t.Fatalf("%s -i: exit %d, %d bytes", name, code, len(want))
		}
		dir := t.TempDir()
		outFile := filepath.Join(dir, "out.txt")
		out, errOut, code := run(t, dir, name, "", "-i", in, "-o", outFile)
		if code != 0 || out != "" || readFile(t, outFile) != want {
			t.Errorf("%s -i -o: exit %d, stdout %d bytes, stderr %q; the file must hold the %d-byte result", name, code, len(out), errOut, len(want))
		}
		// The same input on standard input and as the first argument.
		if out, _, _ := run(t, t.TempDir(), name, readFile(t, in)); out != want {
			t.Errorf("%s: standard input and -i differ", name)
		}
		if out, _, _ := run(t, t.TempDir(), name, "", in); out != want {
			t.Errorf("%s: positional input and -i differ", name)
		}
	}
}

/** TestStandardOptions checks -h, -l and -v in both spellings, before any file
 * is opened, and that the help page is the man page structure.
 */
func TestStandardOptions(t *testing.T) {
	for _, name := range programNames {
		for _, o := range []string{"-h", "-help", "--help"} {
			out, errOut, code := run(t, t.TempDir(), name, "", o, "-i", "no-such-file")
			if code != 0 || errOut != "" || !strings.HasPrefix(out, "%"+name+"(1) user manual | version ") || !strings.Contains(out, "# EXIT STATUS") {
				t.Errorf("%s %s: exit %d, stderr %q, help starts %.50q", name, o, code, errOut, out)
			}
		}
		for _, o := range []string{"-l", "-license", "-v", "-version"} {
			out, _, code := run(t, t.TempDir(), name, "", o)
			if code != 0 || out == "" {
				t.Errorf("%s %s: exit %d, %d bytes", name, o, code, len(out))
			}
		}
	}
}

/** TestExitCodes covers DR-0003 for every program: a missing input is 66, an
 * output that cannot be created is 73, a bad command line is 2, a bad value
 * is 2, and a failing run does not create or empty the output file.
 */
func TestExitCodes(t *testing.T) {
	for _, name := range programNames {
		dir := t.TempDir()
		missing := filepath.Join(dir, "missing.fountain")
		good, _ := filepath.Abs("testdata/sample-01.fountain")
		noDir := filepath.Join(dir, "no", "such", "dir", "out")
		cases := []struct {
			what string
			args []string
			want int
		}{
			{"unknown option", []string{"-nosuchoption"}, 2},
			{"surplus argument", []string{good, filepath.Join(dir, "o1"), filepath.Join(dir, "o2")}, 2},
			{"bad width", []string{"-width", "wide", good}, 2},
			{"missing input file", []string{missing}, 66},
			{"missing input with -i", []string{"-i", missing}, 66},
			{"output cannot be created", []string{"-i", good, "-o", noDir}, 73},
		}
		for _, c := range cases {
			out, errOut, code := run(t, dir, name, "", c.args...)
			if code != c.want {
				t.Errorf("%s, %s: exit %d, want %d (stdout %.60q, stderr %q)", name, c.what, code, c.want, out, errOut)
			}
			if !strings.HasPrefix(errOut, name+": ") || strings.Count(errOut, "\n") != 1 {
				t.Errorf("%s, %s: stderr %q is not one line starting with the program name", name, c.what, errOut)
			}
		}
		// A missing input does not create or empty the output file.
		outFile := filepath.Join(dir, "keep.txt")
		if err := os.WriteFile(outFile, []byte("keep"), 0o644); err != nil {
			t.Fatal(err)
		}
		run(t, dir, name, "", "-i", missing, "-o", outFile)
		if readFile(t, outFile) != "keep" {
			t.Errorf("%s: a missing input changed the existing output file", name)
		}
		// -quiet drops the message, not the exit code.
		out, errOut, code := run(t, dir, name, "", "-quiet", "-i", missing)
		if code != 66 || errOut != "" || out != "" {
			t.Errorf("%s -quiet: exit %d, stdout %q, stderr %q; want 66 and nothing printed", name, code, out, errOut)
		}
	}
}
