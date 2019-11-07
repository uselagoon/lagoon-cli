module github.com/amazeeio/lagoon-cli

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/helloyi/go-sshclient v0.0.0-20190617151846-7e5120a78b77
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/logrusorgru/aurora v0.0.0-20191017060258-dc85c304c434
	github.com/machinebox/graphql v0.2.2
	github.com/manifoldco/promptui v0.3.2
	github.com/mattn/go-runewidth v0.0.5 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.1
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.2.2
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20180810215634-df19058c872c // indirect
	gopkg.in/yaml.v2 v2.2.4
)

// use shreddedbacon fork which has better table formatting for kubectl style tables
// https://github.com/olekukonko/tablewriter/pull/144
replace github.com/olekukonko/tablewriter => github.com/shreddedbacon/tablewriter v0.0.2-0.20191104214435-fac6022f4869
