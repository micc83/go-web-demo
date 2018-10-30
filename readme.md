# Iris web demo

First steps with Go Iris web framework.

## Links

- [Go Lang](https://golang.org/)
- [Iris](https://iris-go.com) - Go web framework.
- [gin](https://github.com/codegangsta/gin) - Simple command line utility for live-reloading Go web applications.

## Commands

**Install GO on MAC with brew**

```
brew install go
```

**Change `$GOPATH`**

Add to your `.bash_profile` or `.zshrc`:

```
export GOPATH=$HOME/your/path
```

If using ZSH in VS Code add the following to your user `settings.json`:

```
"go.gopath": "~/your/path"
```

**Run the project**
```bash
go run main.go
```
Then open <http://localhost:3001>

**Run the project with gin (Auto rebuilding proxy)**

Install GIN:
```bash
go get github.com/codegangsta/gin
```

Run GIN:
```bash
../bin/gin main.go
```

Then open <http://localhost:3000>



