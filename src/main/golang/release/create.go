package release

import (
	"bytes"
	"github.com/helstern/kacl/src/main/golang/changelog"
	"github.com/mitchellh/copystructure"
	"time"
)

func Create(document *changelog.Contents, tag string, date time.Time) (*changelog.Contents, error) {

	clone, err := copystructure.Copy(document)
	if err != nil {
		return nil, err
	}

	contents := clone.(*changelog.Contents)
	changes := contents.Unreleased
	changes.Tag = tag
	changes.Time = date

	rest := bytes.NewBufferString("")
	changes.WriteTo(rest)
	//rest.WriteString("\n")
	rest.WriteString(contents.Rest)
	contents.Rest = rest.String()

	if len(contents.Refs) > 0 {

		lastTag := "TAIL"
		if len(contents.Changes) > 1 {
			last := contents.Changes[1]
			lastTag = last.Tag
		} else if len(contents.Refs) == 1 {
			last := contents.Refs[0]
			lastTag = last.From
		}

		baseRef := contents.Refs[0]
		contents.Refs = append([]changelog.Reference{
			{
				Type:      baseRef.Type,
				Tag:       "Unreleased",
				From:      tag,
				To:        "HEAD",
				BaseURL:   baseRef.BaseURL,
				Separator: baseRef.Separator,
			},
			{
				Type:      baseRef.Type,
				Tag:       tag,
				From:      lastTag,
				To:        tag,
				BaseURL:   baseRef.BaseURL,
				Separator: baseRef.Separator,
			},
		}, contents.Refs[1:]...)
	}

	contents.Unreleased = changelog.NewChanges("Unreleased")
	return contents, nil
}
