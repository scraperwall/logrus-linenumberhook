package linenumberhook

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Hook struct {
	skip int
}

func (hook *Hook) Levels() []log.Level {
	return log.AllLevels
}

func (hook *Hook) Fire(entry *log.Entry) error {
	// determine the call stack skip level for logrus to print the calling file/function/line number
	if hook.skip == -1 {
		i := 0
		for {
			pc, file, _, ok := runtime.Caller(i)

			if !ok {
				hook.skip = -2
				break
			}

			fname := runtime.FuncForPC(pc).Name()
			if !strings.Contains(file, "sirupsen/logrus") && !strings.Contains(fname, "linenumberhook") {
				hook.skip = i
				break
			}

			i++
		}
	}

	// don't try to add the file/func/line number info if the skip level couldn't be determined
	if hook.skip < 0 {
		return nil
	}

	// add the file, func name and line number in each log entry
	if pc, file, line, ok := runtime.Caller(hook.skip + 1); ok {
		funcName := runtime.FuncForPC(pc).Name()

		entry.Data["src"] = fmt.Sprintf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
	}

	return nil
}
