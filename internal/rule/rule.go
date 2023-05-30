package rule

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func foo() {
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar("name", decls.String),
			decls.NewVar("group", decls.String)))
	if err != nil {
		panic(err)
	}

	ast, issues := env.Compile(`name.startsWith("/groups/" + group)`)
	if issues != nil && issues.Err() != nil {
		log.Fatalf("type-check error: %s", issues.Err())
	}
	prg, err := env.Program(ast)
	if err != nil {
		log.Fatalf("program construction error: %s", err)
	}

	out, details, err := prg.Eval(map[string]interface{}{
		"name":  "/groups/acme.co/documents/secret-stuff",
		"group": "acme.co"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v %#v\n", out, details)
}
