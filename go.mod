module github.com/uselagoon/lagoon-cli

go 1.16

require (
	github.com/Masterminds/semver v1.4.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/mock v1.6.0
	github.com/google/go-github v0.0.0-20180716180158-c0b63e2f9bb1
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-version v1.6.0
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/logrusorgru/aurora v0.0.0-20191017060258-dc85c304c434
	github.com/machinebox/graphql v0.2.3-0.20181106130121-3a9253180225
	github.com/manifoldco/promptui v0.3.2
	github.com/olekukonko/tablewriter v0.0.4
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.2.2
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b
	gopkg.in/yaml.v2 v2.2.8
	sigs.k8s.io/yaml v1.2.0
)

require (
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/guregu/null v4.0.0+incompatible
	// workaround for https://github.com/manifoldco/promptui/issues/98
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/uselagoon/machinery v0.0.13-0.20231116024123-c712ade42522
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20191105091915-95d230a53780 // indirect
)

// use this version for fixes to formatting of end header
replace github.com/olekukonko/tablewriter => github.com/shreddedbacon/tablewriter v0.0.2-0.20200114082015-d810c4a558bf

// replace github.com/machinebox/graphql => ../../shreddedbacon/graphql

// replace github.com/olekukonko/tablewriter => ../../shreddedbacon/tablewriter
