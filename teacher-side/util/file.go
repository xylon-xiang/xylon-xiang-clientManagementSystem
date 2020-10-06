package util

import "clientManagementSystem/config"

func GetFilePath(studentId string, studentName string, fileName string) string{

	return config.Config.APIConfig.HomeworkFileFolderPath +
		studentId + studentName + "-" + fileName

}
