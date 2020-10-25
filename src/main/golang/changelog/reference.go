package changelog

import (
	"fmt"
	"io"
	"strings"
)

type ReferenceType string

const (
	GITHUB_REFERENCE    ReferenceType = "github"
	BITBUCKET_REFERENCE               = "bitbucket"
)

func FormatGithubReference(ref Reference, w io.Writer) (int64, error) {
	if ref.To != "" && ref.From != "" {
		n, err := fmt.Fprintf(w, "[%s]: %s/compare/%s%s%s\n", ref.Tag, ref.BaseURL, ref.From, ref.Separator, ref.To)
		return int64(n), err
	}
	n, err := fmt.Fprintf(w, "%s\n", ref.Raw)
	return int64(n), err
}

func FormatBitbucketReference(ref Reference, w io.Writer) (int64, error) {
	if ref.To != "" && ref.From != "" {
		n, err := fmt.Fprintf(w, "[%s]: %s/compare/%s%s%s\n", ref.Tag, ref.BaseURL, ref.To, ref.Separator, ref.From)
		return int64(n), err
	}
	n, err := fmt.Fprintf(w, "%s\n", ref.Raw)
	return int64(n), err
}

type Reference struct {
	Type      ReferenceType
	Tag       string
	Raw       string
	From      string
	To        string
	BaseURL   string
	Separator string
}

func (ref *Reference) WriteTo(w io.Writer) (int64, error) {

	if ref.Type == BITBUCKET_REFERENCE {
		return FormatBitbucketReference(*ref, w)
	}

	return FormatGithubReference(*ref, w)
}

func NewReferenceFromRegexp(matches []string) Reference {
	var from string
	var to string
	var refType ReferenceType

	if strings.HasPrefix(matches[2], "https://bitbucket.org/") || strings.HasPrefix(matches[2], "http://bitbucket.com/") {
		from = matches[5]
		to = matches[3]
		refType = BITBUCKET_REFERENCE
	} else {
		from = matches[3]
		to = matches[5]
		refType = GITHUB_REFERENCE
	}

	return Reference{
		Type:      refType,
		Tag:       matches[1],
		Raw:       matches[0],
		BaseURL:   matches[2],
		From:      from,
		Separator: matches[4],
		To:        to,
	}
}
