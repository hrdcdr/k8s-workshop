package handlers

import (
	"net/http"
	"testing"

	"github.com/hrdcdr/k8s-workshop/pkg/config"
	"github.com/hrdcdr/k8s-workshop/pkg/logger"
	"github.com/hrdcdr/k8s-workshop/pkg/logger/standard"
	"github.com/hrdcdr/k8s-workshop/pkg/router/bitroute"
)

func TestHealth(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Health)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
