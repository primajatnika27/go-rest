package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetupLogger() {
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,                  // Menampilkan timestamp lengkap
		TimestampFormat: "02-01-2006 15:04:05", // Mengatur format waktu
		ForceColors:     true,                  // Memaksa output warna (berguna jika output ke file)
		ForceQuote:      true,                  // Memaksa penulisan field dalam quotes,
	}

	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)       // Mengatur output log, bisa juga ke file dengan os.OpenFile(...)
	logrus.SetLevel(logrus.InfoLevel) // Mengatur level log yang ingin ditampilkan
}
