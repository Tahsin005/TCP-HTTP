package request

import (
	"errors"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	// convert to string and split by CRLF (carriage return + line feed = \r\n)
	text := string(data)
	lines := strings.SplitN(text, "\r\n", 2)

	if len(lines) == 0 || strings.TrimSpace(lines[0]) == "" {
		return nil, errors.New("empty request")
	}

	readLine, err := parseRequestLine(lines[0])
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: readLine,
	}, nil
}

func parseRequestLine(line string) (RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return RequestLine{}, errors.New("invalid request line")
	}

	method, target, version := parts[0], parts[1], parts[2]

	// method must be in uppercase
	if method != strings.ToUpper(method) {
		return RequestLine{}, errors.New("invalid method")
	}

	// version must start with HTTP/
	if !strings.HasPrefix(version, "HTTP/") {
		return RequestLine{}, errors.New("invalid HTTP version")
	}

	return RequestLine{
		Method:        method,
		RequestTarget: target,
		HttpVersion:   strings.TrimPrefix(version, "HTTP/"),
	}, nil
}