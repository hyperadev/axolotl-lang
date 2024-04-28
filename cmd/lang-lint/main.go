package main

import (
	"fmt"
	"os"

	"hypera.dev/axolotl-lang/v2"
)

func main() {
	if err := lang.LoadLocales(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Locales loaded successfully", lang.Bundle().LanguageTags())
}
