package handlers

import (
	"net/http"
	"testing"

	"github.com/k8s-community/k8sapp/pkg/config"
	"github.com/k8s-community/k8sapp/pkg/logger"
	"github.com/k8s-community/k8sapp/pkg/logger/standard"
	"github.com/k8s-community/k8sapp/pkg/router/bitroute"
)

func TestReady(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Ready)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
