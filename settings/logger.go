/*
logger.go
Author: Naveenraj O M
Description: This file initializes the logger using the Zap library and Lumberjack for log rotation.
*/

package settings

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(fileName string, maxSize, maxBkp, maxAge int, compress bool, logLevel string) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,  // Maximum size of a log file in MB
		MaxBackups: maxBkp,   // Maximum number of backup log files to retain
		MaxAge:     maxAge,   // Maximum number of days to retain old log files
		Compress:   compress, // Compress the old log files
	}

	writeSyncer := zapcore.AddSync(lumberjackLogger)

	// Configure the encoder for the log entries.
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO8601 format for timestamps
	encoder := zapcore.NewJSONEncoder(encoderConfig)      // Use JSON format for log entries

	// Determine the log level based on the provided logLevel parameter.
	var level zapcore.Level
	switch logLevel {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	case "FATAL":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel // Default to INFO level if an unknown level is provided
	}

	// Create a core that writes logs to both the console (stdout) and the lumberjack logger.
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writeSyncer),
		level,
	)

	// Initialize the global logger with the core, and add caller info and stack traces for errors.
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// Log a message indicating that the logger has been initialized successfully.
	Log.Info("Logger initialized",
		zap.String("fileName", fileName),
		zap.Int("maxSize(MB)", maxSize),
		zap.Int("maxBackups", maxBkp),
		zap.Int("maxAge(days)", maxAge),
		zap.Bool("compress", compress),
		zap.String("logLevel", logLevel),
	)
}
