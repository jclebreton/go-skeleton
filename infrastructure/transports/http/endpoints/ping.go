//nolint:dupl
package endpoints

import (
	"net/http"
)

func (ep *Endpoints) Ping(w http.ResponseWriter, r *http.Request) {
	ep.logger.Debug("ping/pong")
	w.Write([]byte("pong"))
}
