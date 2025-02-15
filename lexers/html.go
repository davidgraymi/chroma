package lexers

import (
	"github.com/davidgraymi/chroma/v2"
)

// HTML lexer.
var HTML = chroma.MustNewXMLLexer(embedded, "embedded/html.xml")
