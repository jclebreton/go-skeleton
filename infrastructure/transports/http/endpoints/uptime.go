//nolint:dupl
package endpoints

import (
	"net/http"
)

func (ep *Endpoints) Uptime(w http.ResponseWriter, r *http.Request) {
	uptime := ep.usecases.GetAppUptime()
	w.Write([]byte(uptime.String()))
}
