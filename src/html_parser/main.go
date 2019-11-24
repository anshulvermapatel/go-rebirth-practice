package main

    import (
            "fmt"
            "github.com/timolinn/html-parser"
        )

    func main() {
        tmpl := []byte(`<html><body><p id="hw">Hello, world</p></body></html>`)
        parsed := html.Parse(tmpl)
        fmt.Printf("%+v\n", parsed)
    }
