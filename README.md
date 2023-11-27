# Editor CLI

A simple editor package for Golang CLIs. Similar to `kubectl patch`. Edit
content from a remote server, or content locally on disk, and overwrite the
results using your favorite editor.

![](img/example-edit.gif)

## Configuration

This package will respect the `$EDITOR` environment variable when launched. By
default if no env var is found, it will use `vim` to edit the requested file.

## Usage

This repo is meant to be used as a package for other Golang CLI projects.

It has two main functions you can use as a CLI author for editing files:

* Run(originalContent []byte, originalFilePath string)
    + `originalContent` is the content of the text to edit in bytes
    + `originalFilePath` is where the file originated, can be empty if remote. Used to determine the file ext for the tmp file.
* RunLocal(originalFilePath string)
    + `originalFilePath` is where the file originated, can be empty if remote. Used to determine the file ext for the tmp file.

The example below is a shortened version of `main.go` using the Run version:

```golang
// Run the editor to let the user edit the contents in a tmp file
edited, _, err := Run(contents, filePathToEdit)
if err != nil {
        fmt.Println("File editing error: ", err)
        os.Exit(1)
}
```

### Examples

Use the example CLI to see how this package can work for you:

```shell
# Build the CLI with the Makefile in this repo
brian@localghost:editor-cli λ make
# Run the CLI against some examples
brian@localghost:editor-cli λ ./bin/editor-cli --file examples/test.txt
```
