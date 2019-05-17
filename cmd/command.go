package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Code struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

func bind(req *http.Request) Code {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)

	code := Code{}
	json.Unmarshal(buf.Bytes(), &code)
	return code
}

type CodeRunner func(string, io.Writer) error

var (
	coderRunners = map[string]CodeRunner{
		"go":     tempRun("go", "run", "main.go"),
		"java":   tempRun("jshell", "-v", "main.java"),
		"ruby":   evaluator("ruby", "-e"),
		"python": evaluator("python", "-c"),
		"js":     evaluator("node", "-e"),
	}
)

//
// Code Runner
//

func evaluator(cmd, opt string) CodeRunner {
	return func(code string, w io.Writer) error {
		out, err := exec.Command(cmd, opt, code).CombinedOutput()
		if err != nil {
			log.Printf("%s err %v %v\n%s", cmd, err, string(out), code)
			return err
		}

		log.Printf("%s %s %v", cmd, opt, string(out))
		_, err = io.Copy(w, bytes.NewBuffer(out))
		return err
	}
}

// playground code see https://github.com/golang/playground/blob/master/sandbox.go#L291
func tempRun(cmd, opt, tmp string) CodeRunner {
	return func(code string, w io.Writer) error {
		tmpDir, err := ioutil.TempDir("", "sandbox")
		if err != nil {
			return fmt.Errorf("error creating temp directory: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		src := []byte(code)
		in := filepath.Join(tmpDir, tmp)
		if err := ioutil.WriteFile(in, src, 0400); err != nil {
			return fmt.Errorf("error creating temp file %q: %v", in, err)
		}

		log.Printf("%s run file %v", cmd, in)
		out, err := exec.Command(cmd, opt, in).CombinedOutput()
		if err != nil {
			return err
		}

		_, err = io.Copy(w, bytes.NewBuffer(out))
		return err
	}
}

//
// Command run web server
//

func commandRun(w http.ResponseWriter, r *http.Request) {
	//lang := r.FormValue("lang")
	//code, err := url.PathUnescape(r.FormValue("code"))
	//log.Printf("[%v]run: %v", code.Lang, code.Code)
	code := bind(r)
	if r, ok := coderRunners[code.Lang]; ok {
		if err := r(code.Code, w); err != nil {
			fmt.Fprintf(w, "Error [%s]", err)
		}
	} else {
		fmt.Fprintf(w, "Unsupported Language [%s]", code.Lang)
	}
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/run", commandRun)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
