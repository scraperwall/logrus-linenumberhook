# logrus-linenumberhook

logrus-linenumberhook is a hook for logrus that automatically adds the source file and line number where the log method was called to each log entry.

## Installation

```bash
go get github.com/scraperwall/logrus-linenumberhook
```

## Usage

```go
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/scraperwall/logrus-linenumberhook"
)

func main() {
	log.AddHook(lineNumberHook.New())

	// ... your code
}
