package config

import "regexp"

var UsernameRegex = regexp.MustCompile(`^[a-z](?:[a-z]|[._][a-z])*[a-z]$`)
