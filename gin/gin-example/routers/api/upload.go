package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/e"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/logging"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/upload"
	"net/http"
)

/*
// Summary Import Image
// @Produce json
// @Param image fromData file true "Image file"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
*/

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else if err := upload.CheckImage(fullPath); err != nil {
			logging.Warn(err)
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
		} else if err := c.SaveUploadedFile(image, src); err != nil {
			logging.Warn(err)
			code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
		} else {
			data["image_url"] = upload.GetImageFullUrl(imageName)
			data["image_save_url"] = savePath + imageName
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
