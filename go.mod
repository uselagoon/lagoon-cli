module github.com/uselagoon/lagoon-cli

go 1.23

require (
	github.com/Masterminds/semver/v3 v3.3.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/go-github/v66 v66.0.0
	github.com/google/uuid v1.6.0
	github.com/guregu/null v4.0.0+incompatible
	github.com/hashicorp/go-version v1.7.0
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/jedib0t/go-pretty/v6 v6.6.5
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/machinebox/graphql v0.2.3-0.20181106130121-3a9253180225
	github.com/manifoldco/promptui v0.9.0
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
	github.com/skeema/knownhosts v1.3.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.10.0
	github.com/uselagoon/machinery v0.0.31
	go.uber.org/mock v0.5.0
	golang.org/x/crypto v0.31.0
	golang.org/x/term v0.27.0
	gopkg.in/yaml.v3 v3.0.1
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

//replace github.com/uselagoon/machinery => ../machinery
