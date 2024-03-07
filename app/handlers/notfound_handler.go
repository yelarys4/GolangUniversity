package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/yelarys4/GolangUniversity/app/utils"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
	}).Error("404 Not found")
	utils.RespondWithError(w, 404, "Page not found!")
}
