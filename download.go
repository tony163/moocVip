package main

import (
	"fmt"
	"moocVip/util"
	"moocVip/video"
	"os"
)

func Download(InfoStruct util.MyMocTermDto, cookieStr string, token string) {
	basePath, _ := os.Getwd()
	courseName := InfoStruct.CourseName
	courseNamePath := fmt.Sprintf("%s\\download\\%s", basePath, courseName)
	util.PathExists(courseNamePath)

	for _, chapter := range InfoStruct.Chapters {
		chapterName := chapter.ChapterName
		chapterNamePath := fmt.Sprintf("%s\\%s", courseNamePath, chapterName)
		util.PathExists(chapterNamePath)
		for _, unit := range chapter.MyUnits {
			contentType := unit.ContentType
			UnitId := unit.UnitId
			unitName := unit.UnitName

			switch contentType {
			case 1:
				path := fmt.Sprintf("%s\\%s.mp4", chapterNamePath, unitName)
				_, err := os.Stat(path)
				if err != nil {
					video.Download(UnitId, token, cookieStr, unitName, chapterNamePath)
				}
			}
		}
	}

}
