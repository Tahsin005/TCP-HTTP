package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	username := ""
	password := ""

	if parsedUrl.User != nil {
		username = parsedUrl.User.Username()
		if pw, ok := parsedUrl.User.Password(); ok {
			password = pw
		}
	}

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: username,
		password: password,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}
