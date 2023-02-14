package cmd

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/guregu/null"
	"github.com/spf13/pflag"
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

func splitInvokeTaskArguments(invokedTaskArguments []string) (map[string]string, error) {
	parsedArgs := map[string]string{}

	for _, v := range invokedTaskArguments {
		fmt.Println(v)
		split := strings.Split(v, "=")
		if len(split) != 2 {
			return map[string]string{}, errors.New(fmt.Sprintf("Unable to parse `%v`, the form of arguments should be `KEY=VALUE`", v))
		}
		parsedArgs[split[0]] = split[1]
	}
	return parsedArgs, nil
}
