package logs

import (
	"../util"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Log(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start:= time.Now()
		inner.ServeHTTP(res, req)

		elaspsedTime:= time.Since(start)
		elaspsedTimeInNanoSeconds:= int64(elaspsedTime/time.Nanosecond)

		logMessage:= util.APILogMessage{
			Method:      req.Method,
			URI:         req.RequestURI,
			Name:        name,
			TimeElapsed: elaspsedTimeInNanoSeconds,
			Service:     "RecruitApp-Go",
			Type:        "logs",
		}

		message := fmt.Sprintf(
			"%s\t%s\t%s\t%s\t%s",
			time.Now(),
			logMessage.Method,
			logMessage.URI,
			logMessage.Name,
			elaspsedTime,
		)
		log.Println(message)
		err:= LogToFile(message, "logs.txt")
		if err != nil {
			fmt.Println(err)
		}
	})
}

func LogToFile(message, fileName string) error {
	message += "\n"
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 666)
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(message)
		return nil
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

//LogError
func LogError(err error) error {
	errorMessage := util.APIErrorMessage{
		Message: err.Error(),
		Service: "RecruitApp-Go API",
		Type:    "error",
	}
	log.Print(errorMessage)

	err = LogToFile(errorMessage.Message + errorMessage.Service, "api-errors.txt")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}