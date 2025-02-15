package formatters

import (
	"fmt"
	"io"

	"github.com/davidgraymi/chroma/v2"
)

// Tokens formatter outputs the raw token structures.
var Tokens = Register("tokens", chroma.FormatterFunc(func(w io.Writer, s *chroma.Style, it chroma.Iterator, escape bool) error {
	for t := it(); t != chroma.EOF; t = it() {
		if _, err := fmt.Fprintln(w, t.GoString()); err != nil {
			return err
		}
	}
	return nil
}))
