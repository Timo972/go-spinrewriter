# go-spinrewriter

A minimalistic but fully featured go client for the spinrewriter AI API.  
Currently in an experimental state, the package will soon be battle tested and hopefully proven as stable.

**API Stability**
- ✅ Quota
- ✅ Spintax
- 🚧 UniqueVariation
- 🚧 UniqueSpintaxVariation

> ✅ = Tested 🚧 = Untested

## ⚡️ Quickstart

```go
package main

import "github.com/timo972/go-spinrewriter"

func main() {
    client := spinrewriter.New("email", "apiKey")

	spintax, err := client.Spintax("This is a sentence.")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Text options: %d", spintax.NumOptions())
}
```

## ⚙️ Installation

Make sure you have Go installed ([download](https://go.dev/dl/)). Version `1.19` or higher is required.

Initialize your project by creating a folder and then running `go mod init github.com/your/repo` ([learn more](https://go.dev/blog/using-go-modules)) inside the folder. Then install the spinrewriter client with the [`go get`](https://pkg.go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
go get -u github.com/timo972/go-spinrewriter
```

## ⛔️ Limitations

The client does not circumvent limitations set by the Spinrewriter API.
Make sure to check out their [documentation](https://www.spinrewriter.com/api-documentation)!

