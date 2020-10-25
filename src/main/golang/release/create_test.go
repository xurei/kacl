package release

import (
	"bytes"
	"github.com/helstern/kacl/src/main/golang/changelog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var githubStyleUnreleased = `# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Fixed
- something broken
- some issue

### Removed
- some old stuff
- bad code

## [v0.3.0] - 2016-12-03
### Added
- This awesome feature
- More pewpew.

## [v0.2.0] - 2015-10-06
### Changed
- a thingy with some subpoints:
	- this one
	- that one
	- yay!

### Deprecated
- legacy stuff
- args of some function

## [0.1.0] - 2014-09-02
### Security
- hard coded passwords have been removed
- stack overflow issue solved!

[Unreleased]: https://github.com/myuser/myproject/compare/v0.3.0...HEAD
[v0.3.0]: https://github.com/myuser/myproject/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/myuser/myproject/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/myuser/myproject/compare/v0.0.8...v0.1.0
`

var githubStyleReleased = `# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [v0.3.1] - 2016-12-04
### Fixed
- something broken
- some issue

### Removed
- some old stuff
- bad code

## [v0.3.0] - 2016-12-03
### Added
- This awesome feature
- More pewpew.

## [v0.2.0] - 2015-10-06
### Changed
- a thingy with some subpoints:
	- this one
	- that one
	- yay!

### Deprecated
- legacy stuff
- args of some function

## [0.1.0] - 2014-09-02
### Security
- hard coded passwords have been removed
- stack overflow issue solved!

[Unreleased]: https://github.com/myuser/myproject/compare/v0.3.1...HEAD
[v0.3.1]: https://github.com/myuser/myproject/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/myuser/myproject/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/myuser/myproject/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/myuser/myproject/compare/v0.0.8...v0.1.0
`

var bitbucketUnreleased = `# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- add infrastructure resource catalog

[Unreleased]: https://bitbucket.org/helsternware/www.helstern.org/branches/compare/HEAD%0D8765b1b
`
var bitbucketReleased = `# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [v0.1.0] - 2016-12-04
### Added
- add infrastructure resource catalog

[Unreleased]: https://bitbucket.org/helsternware/www.helstern.org/branches/compare/HEAD%0Dv0.1.0
[v0.1.0]: https://bitbucket.org/helsternware/www.helstern.org/branches/compare/v0.1.0%0D8765b1b
`

func TestCreate_github(t *testing.T) {
	now, _ := time.Parse("2006-01-02", "2016-12-04")
	contents, _ := changelog.Parse(bytes.NewBufferString(githubStyleUnreleased))
	released, _ := Create(contents, "v0.3.1", now)

	out := bytes.NewBuffer([]byte{})
	_, err := released.WriteTo(out)
	assert.Nil(t, err)

	assert.Equal(t, githubStyleReleased, out.String())
}
func TestCreate_bitbucket(t *testing.T) {
	now, _ := time.Parse("2006-01-02", "2016-12-04")
	contents, _ := changelog.Parse(bytes.NewBufferString(bitbucketUnreleased))
	released, _ := Create(contents, "v0.1.0", now)

	out := bytes.NewBuffer([]byte{})
	_, err := released.WriteTo(out)
	assert.Nil(t, err)

	assert.Equal(t, bitbucketReleased, out.String())
}
