module github.com/amazeeio/lagoon-cli

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/logrusorgru/aurora v0.0.0-20191017060258-dc85c304c434
	// use latest version of graphql
	github.com/machinebox/graphql v0.0.0-20180711000000-3a9253180225
	github.com/manifoldco/promptui v0.3.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.2
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.5.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/yaml.v2 v2.2.4
)

// use shreddedbacon fork which has better table formatting for kubectl style tables
// https://github.com/olekukonko/tablewriter/pull/144
// replace github.com/olekukonko/tablewriter => github.com/shreddedbacon/tablewriter v0.0.2-0.20191104214435-fac6022f4869
// use this version for fixes to formatting of end header
replace github.com/olekukonko/tablewriter => github.com/shreddedbacon/tablewriter v0.0.2-0.20200114082035-d810c4a558bf

// replace github.com/machinebox/graphql => ../../shreddedbacon/graphql

// replace github.com/olekukonko/tablewriter => ../../shreddedbacon/tablewriter
