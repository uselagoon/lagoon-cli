module github.com/uselagoon/lagoon-cli

go 1.21

require (
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/mock v1.6.0
	github.com/google/go-github v0.0.0-20180716180158-c0b63e2f9bb1
	github.com/google/uuid v1.5.0
	github.com/guregu/null v4.0.0+incompatible
	github.com/hashicorp/go-version v1.6.0
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/jedib0t/go-pretty/v6 v6.5.0
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/machinebox/graphql v0.2.3-0.20181106130121-3a9253180225
	github.com/manifoldco/promptui v0.9.0
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.2
	github.com/uselagoon/machinery v0.0.18
	golang.org/x/crypto v0.21.0
	gopkg.in/yaml.v3 v3.0.1
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/term v0.18.0 // indirect
)

// replace github.com/uselagoon/machinery => ../machinery
