package screenshot

import (
	"clientManagementSystem/module"
	"io"
	"mime/multipart"
	"os"
)

func UpdateScreenFrozeDuration() {

}

func SaveScreenshot(studentName string, file *multipart.FileHeader) error {

	// save the screenshot at path "/screenshot"
	dst, err := os.Create("screenshot/" + studentName + "-" +  file.Filename)
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




}