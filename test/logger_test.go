package test

import (
	"testing"

	"github.com/rogersovich/go-portofolio-v4/utils"
	"github.com/sirupsen/logrus"
)

func TestInitLogger(t *testing.T) {
	utils.InitLogger()

	utils.Log.Debug("Test log")
	utils.Log.Info("Test log")
}

func TestInitLoggerJSON(t *testing.T) {
	utils.InitLogger()

	// Log as JSON instead of the default ASCII formatter.
	utils.Log.SetFormatter(&logrus.JSONFormatter{})

	utils.Log.Debug("Test log")
	utils.Log.WithField("field", "value").Info("Test log with field")
	utils.Log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
