package main

import (
  "net/http"
  "log"
  "fmt"
  "bytes"
  "os"
  "os/exec"
  "io"
  "io/ioutil"
  "path/filepath"
  "go/parser"
  "go/token"
)

type CodeRunner func(string, io.Writer) error

var (
  coderRunners = map[string]CodeRunner{
    "go": GoRun,
  }
)

//
// Code Runner
//

// see https://github.com/golang/playground/blob/master/sandbox.go#L291
func GoRun(code string, w io.Writer) error {
	tmpDir, err := ioutil.TempDir("", "sandbox")
	if err != nil {
		return fmt.Errorf("error creating temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

  src := []byte(code)
	in := filepath.Join(tmpDir, "main.go")
	if err := ioutil.WriteFile(in, src, 0400); err != nil {
		return fmt.Errorf("error creating temp file %q: %v", in, err)
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, in, nil, parser.PackageClauseOnly)
	if err == nil && f.Name.Name != "main" {
		return fmt.Errorf("error main func not found %v", code)
	} else if err != nil {
    return fmt.Errorf("error parse: %v", err)
  }

  log.Printf("go run file %v", in)
  out, err := exec.Command("go", "run", in).CombinedOutput()
  if err != nil {
		return err
	}

  _, err = io.Copy(w, bytes.NewBuffer(out))
  return err
}



//
// Command run web server
//

func commandRun(w http.ResponseWriter, r *http.Request) {
  lang := r.FormValue("lang")
  code := r.FormValue("code")

  log.Printf("[%v]run: %v", lang, code)
  if r, ok := coderRunners[lang]; ok {
    r(code, w)
  } else {
    fmt.Fprintf(w, "Unsupported Language [%s]", lang)
  }
}

func main() {

  http.HandleFunc("/run", commandRun)

  log.Fatal( http.ListenAndServe(":6000", nil))
}
