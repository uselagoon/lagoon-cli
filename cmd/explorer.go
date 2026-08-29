package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/explorer"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

var explorerCmd = &cobra.Command{
	Use:   "explorer",
	Short: "Open GraphQL Explorer",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}
		return server(port)
	},
}

func init() {
	explorerCmd.Flags().String("port", "8091", "Branch name to deploy")
}

func server(port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", serveExplorer)
	r.HandleFunc("/query", runQuery).Methods("POST")

	fmt.Printf("Open http://localhost:%s in your browser to start using the explorer", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		return err
	}
	return nil
}

func runQuery(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	result := ""
	if err != nil {
		result = err.Error()
	}
	raw := r.FormValue("raw")
	if raw != "" {
		validateToken(lagoonCLIConfig.Current)
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			"",
			"",
			&token,
			false)
		rawResp, err := lc.ProcessRaw(context.TODO(), raw, nil)
		if err != nil {
			result = err.Error()
		} else {
			resp, err := json.MarshalIndent(rawResp, "", "\t")
			if err != nil {
				result = err.Error()
			} else {
				result = string(resp)
			}
		}
	}
	_, _ = w.Write([]byte(result))
}

func serveExplorer(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("").ParseFS(explorer.Explorer, "templates/base.html")
	_ = tmpl.ExecuteTemplate(w, "base", nil)
}
