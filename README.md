# :deciduous_tree: find

This is `find`. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

## Installation

If you just want to try it out, grab the release binary with curl and put into your PATH:

```
curl -L https://github.com/tr00datp00nar/find/releases/latest/download/find-linux-amd64 -o ~/.local/bin/find
curl -L https://github.com/tr00datp00nar/find/releases/latest/download/find-darwin-amd64 -o ~/.local/bin/find
curl -L https://github.com/tr00datp00nar/find/releases/latest/download/find-darwin-arm64 -o ~/.local/bin/find
curl -L https://github.com/tr00datp00nar/find/releases/latest/download/find-windows-amd64 -o ~/.local/bin/find
```

Or with `go`:

```shell
go install github.com/tr00datp00nar/find/cmd/find@latest
```

Composed

```go
package c

import (
	Z "github.com/rwxrob/bonzai/z"
    "github.com/tr00datp00nar/find"
)

var Cmd = &Z.Cmd{
	Name:     'c',
    Commands: []*Z.Cmd{help.Cmd, find.Cmd},
}
```

## Resources

To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z
