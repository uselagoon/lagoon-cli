package cmd

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/guregu/null"
	"github.com/spf13/pflag"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// makeSafe ensures that any string is dns safe
func makeSafe(in string) string {
	out := regexp.MustCompile(`[^0-9a-z-]`).ReplaceAllString(
		strings.ToLower(in),
		"$1-$2",
	)
	return out
}

// shortenEnvironment shortens the environment name down the same way that Lagoon does
func shortenEnvironment(project, environment string) string {
	overlength := 58 - len(project)
	if len(environment) > overlength {
		environment = fmt.Sprintf("%s-%s", environment[0:overlength-5], hashString(environment)[0:4])
	}
	return environment
}

// hashString get the hash of a given string.
func hashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func flagStringNullValueOrNil(flags *pflag.FlagSet, flag string) (*null.String, error) {
	flagValue, err := flags.GetString(flag)
	if err != nil {
		return nil, err
	}
	changed := flags.Changed(flag)
	if changed && flagValue == "" {
		// if the flag is defined, and is empty value, return a `null` string
		return &null.String{}, nil
	} else if changed {
		// otherwise set the flag to be the value from the flag
		value := null.StringFrom(flagValue)
		return &value, nil
	}
	// if not defined, return nil
	return nil, nil
}

// buildVarsToMap is used to convert incoming build variable arguments into a structure consumable by the graphQL call
func buildVarsToMap(slice []string) ([]schema.EnvKeyValueInput, error) {
	result := []schema.EnvKeyValueInput{}

	for _, entry := range slice {
		// Split the entry by "="
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) != 2 {
			return []schema.EnvKeyValueInput{}, errors.New("Malformed build variable entry (expects `KEY=VALUE`) got: " + entry)
		}

		// Trim spaces from key and value
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		result = append(result, schema.EnvKeyValueInput{Name: key, Value: value})
	}

	return result, nil
}
