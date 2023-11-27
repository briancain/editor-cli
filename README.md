# Editor CLI

A simple editor package for Golang CLIs. Similar to `kubectl patch`.

## Usage

This repo is meant to be used as a package for other Golang CLI projects.

The example below is a shortened version of `main.go`:

```golang
// Read the original contents of the file
contents, err := os.ReadFile(filePathToEdit)
if err != nil {
        fmt.Println("File reading error: ", err)
        os.Exit(1)
}

// Run the editor to let the user edit the contents in a tmp file
edited, _, err := Run(contents, filePathToEdit)
if err != nil {
        fmt.Println("File editing error: ", err)
        os.Exit(1)
}

// If changes, overwrite the original existing file
file, err := os.OpenFile(filePathToEdit, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
if err != nil {
        fmt.Println("Failed to open original file to update: ", err)
        os.Exit(1)
}
defer file.Close()

// Write the content to the file
estr := string(edited)
_, err = file.WriteString(estr)
if err != nil {
        fmt.Println("Failed to write content to file: ", err)
        os.Exit(1)
}

fmt.Println(fmt.Sprintf("Successfully updated %q!", filePathToEdit))
os.Exit(0)
```

### Examples

Use the example CLI to see how this package can work for you:

```shell
# Build the CLI with the Makefile in this repo
brian@localghost:editor-cli λ make
# Run the CLI against some examples
brian@localghost:editor-cli λ ./bin/editor-cli --file examples/test.txt
```
