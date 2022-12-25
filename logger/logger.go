package logger

import (
	"AppCMDBService/setting"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var Logger *zap.Logger

// Init 初始化日志
func Init() error {

	// 1. encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"                          // 时间的key,默认是ts
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 日期格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // level 大写
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	// encoder := zapcore.NewConsoleEncoder(encoderCfg)

	// 2. writeSyncer
	lumberJackLogger := &lumberjack.Logger{
		Filename:   setting.Conf.LogConfig.Filename,
		MaxSize:    setting.Conf.LogConfig.MaxSize,
		MaxBackups: setting.Conf.LogConfig.MaxBackups,
		MaxAge:     setting.Conf.LogConfig.MaxAge,
		Compress:   false,
	}
	lumberJackLoggerError := &lumberjack.Logger{
		Filename:   strings.Replace(setting.Conf.LogConfig.Filename, ".log", "_err.log", 1),
		MaxSize:    setting.Conf.LogConfig.MaxSize,
		MaxBackups: setting.Conf.LogConfig.MaxBackups,
		MaxAge:     setting.Conf.LogConfig.MaxAge,
		Compress:   false,
	}
	level, err := zapcore.ParseLevel(setting.Conf.LogConfig.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	// 3. 创建zapCore
	coreInfo := zapcore.NewCore(encoder, zapcore.AddSync(lumberJackLogger), level)
	coreError := zapcore.NewCore(encoder, zapcore.AddSync(lumberJackLoggerError), zapcore.ErrorLevel)
	core := zapcore.NewTee(coreInfo, coreError)

	// 4. 利用core生成logger
	lg := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(lg) // 替换zap全局的logger
	zap.L().Info("logger init success")

	Logger = lg
	//Logger.Info("logger init success2")

	return nil

}
