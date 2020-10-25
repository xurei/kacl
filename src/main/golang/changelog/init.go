package changelog

import (
	"html/template"
	"io"
	"strings"
)

const v1TemplateBitbucket string = `# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added 
- This CHANGELOG file!


[Unreleased]: {{ .ProjectURL }}/branches/compare/HEAD%0D{{ .InitialTag }} `

const v1Template string = `# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added 
- This CHANGELOG file!


[Unreleased]: {{ .ProjectURL }}/compare/{{ .InitialTag }}...HEAD `

type InitTemplateData struct {
	ProjectURL string
	InitialTag string
}

func Init(wr io.Writer, data InitTemplateData) error {

	var t *template.Template

	if strings.HasPrefix(data.ProjectURL, "https://bitbucket.org/") || strings.HasPrefix(data.ProjectURL, "http://bitbucket.com/") {
		t = template.Must(template.New("version").Parse(v1TemplateBitbucket))
	} else {
		t = template.Must(template.New("version").Parse(v1Template))
	}

	return t.Execute(wr, data)
}
