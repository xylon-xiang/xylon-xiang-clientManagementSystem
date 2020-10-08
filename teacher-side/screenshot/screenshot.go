package screenshot

import (
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/util"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func SaveScreenshot(studentName string, file *multipart.FileHeader) error {

	// save the screenshot at path "/screenshot"
	// time now
	dst, err := os.Create("screenshot/" + studentName + "-" +  file.Filename + time.Now().String())
	if err != nil{
		return err
	}
	src, err := file.Open()
	if err != nil{
		return err
	}
	if _, err = io.Copy(dst,src); err != nil{
		return err
	}
	return nil
}

func UpdateScreenFrozenDuration(studentStatus module.StudentStatus) error {

	startDate := strconv.FormatInt(studentStatus.ClassStartDate, 10)
	filter := map[string]string{
		"StudentId": studentStatus.StudentId,
		"ClassName": studentStatus.ClassName,
		"ClassStartDate": startDate,
	}

	result, err := util.FindOne(util.STUDENTSTATUS, filter)
	if err != nil{
		return err
	}

	status := result.(module.StudentStatus)

	status.ScreenFrozenDuration += config.Config.APIConfig.ScreenshotAPI.FrozenDuration

	err = util.UpdateOne(util.STUDENTSTATUS, status)
	if err != nil{
		return err
	}

	return nil
}