// cmd/blockchainnftminterapiultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftminterapiultra/internal/blockchainnftminterapiultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftminterapiultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
