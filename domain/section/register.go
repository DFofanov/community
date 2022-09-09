// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package section

import (
	"fmt"

	"github.com/DFofanov/community/core/env"
	"github.com/DFofanov/community/domain/section/airtable"
	"github.com/DFofanov/community/domain/section/code"
	"github.com/DFofanov/community/domain/section/flowchart"
	"github.com/DFofanov/community/domain/section/frame"
	"github.com/DFofanov/community/domain/section/gemini"
	"github.com/DFofanov/community/domain/section/jira"
	"github.com/DFofanov/community/domain/section/markdown"
	"github.com/DFofanov/community/domain/section/papertrail"
	"github.com/DFofanov/community/domain/section/pdfjs"
	"github.com/DFofanov/community/domain/section/plantuml"
	"github.com/DFofanov/community/domain/section/provider"
	"github.com/DFofanov/community/domain/section/table"
	"github.com/DFofanov/community/domain/section/tabular"
	"github.com/DFofanov/community/domain/section/trello"
	"github.com/DFofanov/community/domain/section/wysiwyg"
	"github.com/DFofanov/community/domain/store"
)

// Register sections
func Register(rt *env.Runtime, s *store.Store) {
	provider.Register("code", &code.Provider{Runtime: rt, Store: s})
	provider.Register("jira", &jira.Provider{Runtime: rt, Store: s})
	provider.Register("gemini", &gemini.Provider{Runtime: rt, Store: s})
	provider.Register("markdown", &markdown.Provider{Runtime: rt, Store: s})
	provider.Register("papertrail", &papertrail.Provider{Runtime: rt, Store: s})
	provider.Register("tabular", &tabular.Provider{Runtime: rt, Store: s})
	provider.Register("table", &table.Provider{Runtime: rt, Store: s})
	provider.Register("code", &code.Provider{Runtime: rt, Store: s})
	provider.Register("trello", &trello.Provider{Runtime: rt, Store: s})
	provider.Register("wysiwyg", &wysiwyg.Provider{Runtime: rt, Store: s})
	provider.Register("airtable", &airtable.Provider{Runtime: rt, Store: s})
	provider.Register("plantuml", &plantuml.Provider{Runtime: rt, Store: s})
	provider.Register("flowchart", &flowchart.Provider{Runtime: rt, Store: s})
	provider.Register("pdf", &pdfjs.Provider{Runtime: rt, Store: s})
	provider.Register("frame", &frame.Provider{Runtime: rt, Store: s})

	p := provider.List()
	rt.Log.Info(fmt.Sprintf("Extensions: registered %d section types", len(p)))
}
