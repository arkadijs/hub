// +build !windows

package lifecycle

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/agilestacks/hub/cmd/hub/config"
)

const tailLines = 20

type tail struct {
	out   io.Writer
	cols  int
	lines int
	bytes int
}

type windowSize struct {
	rows uint16
	cols uint16
}

func newTail(out *os.File) io.Writer {
	var sz windowSize
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		out.Fd(), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	rows := int(sz.rows)
	cols := int(sz.cols)
	if config.Trace {
		log.Printf("Terminal rows: %d; cols: %d", rows, cols)
	}
	return &tail{out: out, cols: cols}
}

func (t *tail) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if bytes.IndexByte(p, '\n') == -1 {
		written, err := t.out.Write(p)
		t.bytes += written
		return written, err
	}

	lines := bytes.SplitAfter(p, []byte{0x0a})
	l := len(lines)
	if l > 0 && len(lines[l-1]) == 0 {
		lines = lines[:l-1]
	}

	written := 0

	l = len(lines)
	if l > tailLines {
		for _, line := range lines[:l-tailLines] {
			written += len(line)
		}
		lines = lines[l-tailLines:]
	}

	l = len(lines)
	if l+t.lines > tailLines {
		t.out.Write([]byte(fmt.Sprintf("\033[%dA\033[0J", t.lines)))
		t.lines = 0
	}

	var err error
	for _, line := range lines {
		w := 0
		w, err = t.out.Write(line)
		written += w
		if err != nil {
			break
		}
		// last line may not end with a newline
		l := len(line)
		if l > 0 && line[l-1] == '\n' {
			overflow := 0
			if t.cols > 0 {
				overflow = (w + t.bytes - 1) / t.cols // count no control sequences, nor unicode
			}
			t.lines += 1 + overflow
			t.bytes = 0
		} else {
			t.bytes = w
		}
	}

	return written, err
}