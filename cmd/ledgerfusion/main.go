// cmd/ledgerfusion/main.go
package main

import (
"flag"
"log"
"os"

"ledgerfusion/internal/ledgerfusion"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ledgerfusion.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
