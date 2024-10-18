package main

import (
	"fmt"
	"inter/evaluator"
	"inter/lexer"
	"inter/object"
	"inter/parser"
	"inter/repl"
	"io"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) > 1 {
		executeScript(os.Args[1])
		os.Exit(0)
	}
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func executeScript(filename string) {
	out := os.Stdout
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	script := string(bytes)
	l := lexer.New(script)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(out, p.Errors())
		os.Exit(1)
	}
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}
