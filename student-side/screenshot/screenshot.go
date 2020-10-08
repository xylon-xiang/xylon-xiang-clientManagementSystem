package screenshot

import (
	"bytes"
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/util"
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getScreenShot() (filePaths []string, err error) {

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return nil, err
		}

		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())

		filePath := "screenshot/" + fileName

		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}

		err = png.Encode(file, img)
		if err != nil {
			return nil, err
		}

		file.Close()

		filePaths = append(filePaths, filePath)

	}

	return filePaths, nil
}

func SendScreenshot(inputValues *module.InputValue) error {

	filePaths, err := getScreenShot()
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	value := *inputValues

	startDate := strconv.FormatInt(value.ClassStartDate, 10)
	_ = writer.WriteField("studentName", value.StudentName)
	_ = writer.WriteField("studentId", value.StudentId)
	_ = writer.WriteField("className", value.ClassName)
	_ = writer.WriteField("classStartDate", startDate)

	for _, filePath := range filePaths {

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		part, err := writer.CreateFormFile("screenshot", file.Name())
		if err != nil {
			return err
		}

		_, err = io.Copy(part, file)

	}

	// send the multipart/form-data POST Http request
	hostUrl := config.Config.APIConfig.TeacherHost +
		config.Config.APIConfig.ScreenshotAPI.Path
	uri := util.GetRealUrl(hostUrl, value.StudentId)
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func MonitorDesktop(studentId string) error {

	var (
		sample int64
		diff   int64
	)

	files := make([]*os.File, 2)
	i := 0

	// get the init screenshot
	filePaths, err := getScreenShot()
	if err != nil {
		return err
	}

	files[i&1], err = os.Open(filePaths[0])
	if err != nil {
		return err
	}

	i++

	time.Sleep(time.Duration(config.Config.APIConfig.ScreenshotAPI.FrozenDuration))

	for true {
		filePaths, err := getScreenShot()
		if err != nil {
			return err
		}

		files[i&1], err = os.Open(filePaths[0])
		if err != nil {
			return err
		}

		oldPic, _, err := image.Decode(files[i&1])
		if err != nil {
			return err
		}

		newPic, _, err := image.Decode(files[(i+1)&1])
		if err != nil {
			return err
		}

		oldPicBounds := oldPic.Bounds()
		newPicBounds := newPic.Bounds()

		if !boundsMatch(oldPicBounds, newPicBounds) {
			log.Printf("")
			continue
		}

		for y := oldPicBounds.Min.Y; y < newPicBounds.Max.Y; y++ {
			for x := oldPicBounds.Min.X; x < newPicBounds.Max.X; x++ {
				sample += getColor(newPic.At(x, y))
				diff += compareColor(oldPic.At(x, y), newPic.At(x, y))
			}
		}

		// if the change between two pictures is less than ChangeThreshold, send report to teacher
		if float32(diff)/float32(sample) < config.Config.APIConfig.ScreenshotAPI.ChangeThreshold {

			// send put request to teacher
			err := ReportFrozenToTeacher(studentId)
			if err != nil {
				return err
			}

		}

		i = (i + 1) % 2

		time.Sleep(time.Duration(config.Config.APIConfig.ScreenshotAPI.FrozenDuration))

	}

	return nil
}

func compareColor(a, b color.Color) (diff int64) {
	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()

	diff += int64(math.Abs(float64(r1 - r2)))
	diff += int64(math.Abs(float64(g1 - g2)))
	diff += int64(math.Abs(float64(b1 - b2)))
	diff += int64(math.Abs(float64(a1 - a2)))
	return diff
}

func getColor(picture color.Color) (sum int64) {

	r, g, b, a := picture.RGBA()

	return int64(r + g + b + a)
}

func boundsMatch(a, b image.Rectangle) bool {
	return a.Min.X == b.Min.X && a.Min.Y == b.Min.Y && a.Max.X == b.Max.X && a.Max.Y == b.Max.Y
}
