package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	// Output to stdout instead of the default stderr
	Log.SetOutput(os.Stdout)

	// Set log format (pretty & colored)
	// Log.SetFormatter(&logrus.TextFormatter{
	// 	FullTimestamp: true,
	// })

	// Log as JSON instead of the default ASCII formatter.
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Set default level (you can change to Warn/Error in production)
	Log.SetLevel(logrus.DebugLevel)
}

// 🔁 Reusable logging function for query errors
func LogError(message string, fields string) {
	if fields == "" {
		Log.Warn(message)
	} else {
		fieldsQuery := map[string]interface{}{"query": fields}

		Log.WithFields(fieldsQuery).Warn(message)
	}
}

func LogWarning(message string, fields string) {

	if fields == "" {
		Log.Warn(message)
	} else {
		fieldsQuery := map[string]interface{}{"query": fields}

		Log.WithFields(fieldsQuery).Warn(message)
	}

}
