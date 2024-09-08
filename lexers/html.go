package lexers

import (
	"github.com/cakekindel/chroma/v2"
)

// HTML lexer.
var HTML = chroma.MustNewXMLLexer(embedded, "embedded/html.xml")
