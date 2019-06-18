package env

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	ENV_WHITESPACE     = " \t"
	ENV_COMMENT_PREFIX = "#"
)

var (
	// sample_sample__name=testteste=te121ss\sdasda
	// just ignore # sample_sample__name=testteste=te121ss\sdasda
	pattern = regexp.MustCompile("(\\w+)=(.*)")
)

func Pattern() *regexp.Regexp {
	return pattern
}

// []string{key, value}
func ParseKey(key string) string {
	lowered := strings.ToLower(key)
	return strings.ReplaceAll(lowered, "__", ".")
}

//
// nil, nil means comment
//
func ParseEnvLine(line string) (pair Pair, err error) {
	trimmed := strings.TrimLeft(line, " \t")
	if !strings.HasPrefix(trimmed, ENV_COMMENT_PREFIX) {
		groups := Pattern().FindStringSubmatch(trimmed)
		if len(groups) != 3 {
			return EmptyPair(), errors.New("")
		}
		pair = NewPair(ParseKey(groups[1]), groups[2])
		return pair, nil
	} else {
		return EmptyPair(), nil
	}
}

func ParseEnv(r io.Reader) (envs Pairs, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		log.Printf("line %#v\n", line)

		pair, err := ParseEnvLine(line)

		if err != nil {
			log.Printf("error can't read line: %v\n", line)
			continue
		}
		// if pair not empty
		if !IsEmptyPair(pair) {
			envs = append(envs, pair)
		}
	}

	return envs, nil
}
