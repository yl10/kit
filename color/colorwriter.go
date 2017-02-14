// +build !windows

package color

import "io"

type colorWriter struct {
	w    io.Writer
	mode outputMode
}

func (cw *colorWriter) Write(p []byte) (int, error) {
	return cw.w.Write(p)
}
