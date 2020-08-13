package logging

import (
	"fmt"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/setting"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s",
		setting.AppSetting.RuntimeRootPath,
		setting.AppSetting.LogSavePath,
	)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}
