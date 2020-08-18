package util

import "github.com/stephenchen/go-learning/gin/gin-example/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
