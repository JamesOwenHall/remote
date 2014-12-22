# remote

The remote package provides an API to run commands on a running process without needing access to its standard input or output.

## Example

You have started a web server process in the background and would like to disable view caching.  In your web server code, you include:

```Go
r, _ := remote.New("/tmp/myserver.sock")
r.Commands["dev-mode"] = func(args []string) string {
    switch args[0] {
    case "true":
        // disable view caching here
        return "dev-mode is now enabled"
    case "false":
        // enable view caching here
        return "dev-mode is disabled"
    default:
        return "dev-mode command expects either true or false"
    }
}
go r.ListenAndServe()
```

Then at any time, you can run the following from your shell.

    remotec -p /tmp/myserver.sock dev-mode true

which outputs

    dev-mode is now enabled

The remote package is based on Unix domain sockets which are only available locally.  This way you're not exposing any commands to the network.

## Installation

remote comes in two pieces: the remote package for the server and the remotec client.  For this reason, you need to build remotec in order to use the remote package.

1. `go get github.com/JamesOwenHall/remote`
2. `go install github.com/JamesOwenHall/remote/remotec`


## License

The MIT License (MIT)

Copyright (c) 2014 James Hall

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
