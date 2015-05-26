package http

import (
	cutils "github.com/my-open-falcon/common/utils"
	"github.com/my-open-falcon/transfer/proc"
	"net/http"
	"strings"
)

func configProcHttpRoutes() {
	// statistics
	http.HandleFunc("/statistics/all", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, proc.GetAll())
	})

	// trace
	http.HandleFunc("/trace/", func(w http.ResponseWriter, r *http.Request) {
		urlParam := r.URL.Path[len("/trace/"):]
		args := strings.Split(urlParam, "/")

		argsLen := len(args)
		endpoint := args[0]
		metric := args[1]
		tags := make(map[string]string)
		if argsLen > 2 {
			tagVals := strings.Split(args[2], ",")
			for _, tag := range tagVals {
				tagPairs := strings.Split(tag, "=")
				if len(tagPairs) == 2 {
					tags[tagPairs[0]] = tagPairs[1]
				}
			}
		}
		proc.RecvDataTrace.SetPK(cutils.PK(endpoint, metric, tags))
		RenderDataJson(w, proc.RecvDataTrace.GetAllTraced())
	})
}
