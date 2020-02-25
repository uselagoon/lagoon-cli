module github.com/amazeeio/lagoon-cli

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/logrusorgru/aurora v0.0.0-20191017060258-dc85c304c434
	github.com/machinebox/graphql v0.2.3-0.20181106130121-3a9253180225
	github.com/manifoldco/promptui v0.3.2
	github.com/matryer/is v1.2.0 // indirect
	github.com/olekukonko/tablewriter v0.0.4
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.5.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20191105091915-95d230a53780 // indirect
	gopkg.in/yaml.v2 v2.2.4
)

require (
	github.com/Masterminds/semver v1.4.2
	github.com/google/go-github v0.0.0-20180716180158-c0b63e2f9bb1
	github.com/google/go-querystring v1.0.0 // indirect
	// workaround for https://github.com/manifoldco/promptui/issues/98
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/stretchr/testify v1.2.2
)

// use this version for fixes to formatting of end header
replace github.com/olekukonko/tablewriter => github.com/shreddedbacon/tablewriter v0.0.2-0.20200114082015-d810c4a558bf

// replace github.com/machinebox/graphql => ../../shreddedbacon/graphql

// replace github.com/olekukonko/tablewriter => ../../shreddedbacon/tablewriter
