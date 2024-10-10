package logger

import (
	"io"
	"os"
	"path"

	"bicomsystems.com/network/remote-agent/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zerolog.Logger
}

func Configure(cfg config.Config) *Logger {
	var writers []io.Writer

	if cfg.LogConsoleEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if cfg.LogFileEnabled {
		writers = append(writers, newRollingFile(cfg))
	}
	mw := io.MultiWriter(writers...)

	mll := getMinumumLogLevel(cfg.LogLevelMin)
	zerolog.SetGlobalLevel(mll)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(mw).With().Timestamp().Logger()
	logger.Info().
		Bool("fileLogging", cfg.LogFileEnabled).
		Bool("jsonLogOutput", cfg.LogEncodeAsJson).
		Str("logDirectory", cfg.LogDirectory).
		Str("fileName", cfg.LogFilename).
		Int("maxSizeMB", cfg.LogMaxSize).
		Int("maxBackups", cfg.LogMaxBackups).
		Int("maxAgeInDays", cfg.LogMaxAge).
		Msg("logging configured")

	return &Logger{
		Logger: &logger,
	}
}

func newRollingFile(cfg config.Config) io.Writer {
	if err := os.MkdirAll(cfg.LogDirectory, 0744); err != nil {
		log.Error().Err(err).Str("path", cfg.LogDirectory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(cfg.LogDirectory, cfg.LogFilename),
		MaxBackups: cfg.LogMaxBackups,
		MaxSize:    cfg.LogMaxSize,
		MaxAge:     cfg.LogMaxAge,
	}
}

func getMinumumLogLevel(logLevelMin int) zerolog.Level {
	switch logLevelMin {
	case -1:
		return zerolog.TraceLevel
	case 0:
		return zerolog.DebugLevel
	case 1:
		return zerolog.InfoLevel
	case 2:
		return zerolog.WarnLevel
	case 3:
		return zerolog.ErrorLevel
	case 4:
		return zerolog.FatalLevel
	case 5:
		return zerolog.PanicLevel
	default:
		return zerolog.NoLevel
	}
}
