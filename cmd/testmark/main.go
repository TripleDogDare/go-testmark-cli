package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/serum-errors/go-serum"
	"github.com/warpfork/go-testmark"
)

var helptext = strings.TrimSpace(`
testmark <subcommand>
	list    : for given files prints filename:hunk:line
	extract : prints the given hunk or writes the hunk paths to disk
`)

type subcommand func() error

var commands = map[string]subcommand{
	"find":    cmdFind,
	"list":    cmdList,
	"extract": cmdExtract,
	"exec":    cmdExec,
}

const (
	ECodeIo             = "tmcli-error-io"
	ECodeNotImplemented = "tmcli-error-not-implemented"
	ECodeInvalidArgs    = "tmcli-error-invalid-args"
	ECodeHideOutput     = "tmcli-error-hide-output"
)

func argN(i int) string {
	if len(os.Args) < i+1 {
		return ""
	}
	return os.Args[i]
}

func main() {
	log.SetFlags(0)
	cmd, ok := commands[argN(1)]
	if !ok {
		log.Printf("invalid subcommand: %q", argN(1))
		log.Println(helptext)
		os.Exit(1)
	}
	if err := cmd(); err != nil {
		if serum.Code(err) == ECodeHideOutput {
			os.Exit(1)
		}
		log.Println(err)
		os.Exit(1)
	}

}

func cmdFind() error {
	return serum.Error(ECodeNotImplemented, serum.WithMessageLiteral("find not implemented: coming soon TM"))
}

// parseDocs will return a map of strings to testmark Documents
// It does not return an error, instead it logs errors and will not create an
// entry in the result map for the file that failed to read.
func parseDocs(files []string) map[string]*testmark.Document {
	docs := make(map[string]*testmark.Document, len(files))
	for _, filename := range files {
		doc, err := testmark.ReadFile(filename)
		if err != nil {
			log.Println(serum.Error(ECodeIo,
				serum.WithMessageTemplate("unable to read file: {{filename}}"),
				serum.WithCause(err),
				serum.WithDetail("filename", filename),
			))
			continue
		}
		docs[filename] = doc
	}
	return docs
}

func cmdList() error {
	docs := parseDocs(os.Args[2:])
	for filename, doc := range docs {
		for _, hunk := range doc.DataHunks {
			fmt.Printf("%s:%s:%d\n", filename, hunk.Hunk.Name, hunk.LineStart)
		}
	}
	return nil
}

func cmdExtract() error {
	hunkName := argN(2)
	if hunkName == "" {
		log.Println("extract command requires a hunk name argument")
		log.Println("extract [hunk name] [files...]")
		return serum.Error(ECodeInvalidArgs, serum.WithMessageLiteral("extract command requires a hunk name argument"))
	}
	files := os.Args[3:]
	docs := parseDocs(files)
	for filename, doc := range docs {
		if hunk, ok := doc.HunksByName[hunkName]; ok {
			log.Printf("%s:%s:%d", filename, hunk.Hunk.Name, hunk.LineStart)
			fmt.Println(string(hunk.Hunk.Body))
		}
	}
	return nil
}

func cmdExec() error {
	return serum.Error(ECodeNotImplemented, serum.WithMessageLiteral("exec not implemented: requires changing testexec to take an interface instead of a testing.T"))
}
