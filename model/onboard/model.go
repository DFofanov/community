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

// Package onboarding handles the setup of sample data for a new Documize instance.
package onboard

import (
	"github.com/DFofanov/community/domain"
	"github.com/DFofanov/community/model/attachment"
	"github.com/DFofanov/community/model/category"
	"github.com/DFofanov/community/model/doc"
	"github.com/DFofanov/community/model/label"
	"github.com/DFofanov/community/model/link"
	"github.com/DFofanov/community/model/page"
	"github.com/DFofanov/community/model/space"
)

// SampleData holds initial welcome data used during installation process.
type SampleData struct {
	LoadFailure        bool // signals any data load failure
	Context            domain.RequestContext
	Category           []category.Category
	CategoryMember     []category.Member
	Document           []doc.Document
	DocumentAttachment []attachment.Attachment
	DocumentLink       []link.Link
	Section            []page.Page
	SectionMeta        []page.Meta
	Space              []space.Space
	SpaceLabel         []label.Label
}
