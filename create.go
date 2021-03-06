package useragent

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var createInputRegex = regexp.MustCompile(`^\S+$`)

// ErrNameInvalid is returned if the appName contains any whitespace characters
var ErrNameInvalid = errors.New("invalid appName")

// ErrVersionInvalid is returned if the version contains any whitespace characters
var ErrVersionInvalid = errors.New("invalid version")

// ErrEnvironmentInvalid is returned if the environment contains any whitespace characters
var ErrEnvironmentInvalid = errors.New("invalid environment")

// Format will generate a useragent string which is compliant with RFC 22.
func (ua *UserAgent) Format() (string, error) {
	appNameMatch := createInputRegex.MatchString(ua.Name)
	if !appNameMatch {
		return "", fmt.Errorf("error validating %q: %w", ua.Name, ErrNameInvalid)
	}
	versionMatch := createInputRegex.MatchString(ua.Version)
	if !versionMatch {
		return "", fmt.Errorf("error validating %q: %w", ua.Version, ErrVersionInvalid)
	}
	environmentMatch := createInputRegex.MatchString(ua.Environment)
	if !environmentMatch {
		return "", fmt.Errorf("error validating %q: %w", ua.Environment, ErrEnvironmentInvalid)
	}

	var useragent strings.Builder

	compulsory := fmt.Sprintf("%s/%s (Kiwi.com %s)", ua.Name, ua.Version, ua.Environment)
	useragent.WriteString(compulsory)
	if ua.SystemInfo != "" {
		useragent.Grow(1 + len(ua.SystemInfo))
		useragent.WriteString(" ")
		useragent.WriteString(ua.SystemInfo)
	}

	return useragent.String(), nil
}
