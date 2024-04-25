package util

import (
	"fmt"
	"go-read-apache2-error-logs/util/env"
	"go-read-apache2-error-logs/util/mail"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func ExtractHTTPStatusCode(logEntry string) string {
	// Split the log entry by space
	parts := strings.Split(logEntry, " ")
	// The HTTP status code is the 9th element in the split parts
	return parts[8]
}

func WriteFile() string /*os.File*/ {
	t := time.Now().Format("2006-01-02_15-04-05")
	// f, err := os.Create("apache_logs_" + t + ".txt")
	fname := "apache_logs_" + t + ".txt"
	if err := os.WriteFile(fname, []byte{}, 0644); err != nil {
		fmt.Println("********************************CREATE FILE ERROR********************************")
		fmt.Println(err)
		fmt.Println("********************************CREATE FILE ERROR********************************")
	}

	return fname
}

func WriteLog( /*f *os.File*/ filename string, text string) {
	// _, err := f.WriteString(text)
	// if err != nil {
	// 	fmt.Println("********************************WRITE LOG ERROR********************************")
	// 	fmt.Println(err)
	// 	fmt.Println("********************************WRITE LOG ERROR********************************")
	// }
	// f.Sync()

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("********************************OPEN FILE ERROR********************************")
		fmt.Println(err)
		fmt.Println("********************************OPEN FILE ERROR********************************")
	}

	defer f.Close()

	if _, err = f.WriteString(text + "\n"); err != nil {
		fmt.Println("********************************WRITE LOG ERROR********************************")
		fmt.Println(err)
		fmt.Println("********************************WRITE LOG ERROR********************************")
	}
}

func Scheduler() {
	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	// gunakan crontab string untuk mengatur jadwal
	scheduler.AddFunc("0 * * * *", func() { LogsFileMaintainer() }) // start scheduler
	go scheduler.Start()
}

func LogsFileMaintainer() {
	var filePath []string

	smtpConfig := env.GetSMTPConfig()
	smtpClient := mail.InitEmail(smtpConfig)
	Email := []string{"frans.imanuel@visionet.co.id", "lishera.prihatni@visionet.co.id", "ari.darmawan@visionet.co.id", "azky.muhtarom@visionet.co.id"}

	exist, fname := FindFile()
	if !exist {
		filePath = append(filePath, WriteFile())
	} else {
		filePath = append(filePath, fname)
	}

	if CheckFileLength(fname) > 0 {
		if err := smtpClient.Send(Email, nil, nil, "Apache Logs", "text/html", "Apache Logs with txt", filePath); err != nil {
			fmt.Println("***************************ERROR SEND EMAIL***************************")
			fmt.Println(err)
			fmt.Println("***************************ERROR SEND EMAIL***************************")
		}
	}

	if err := os.Remove(fname); err != nil {
		fmt.Println("***************************ERROR DELETE LOG FILE***************************")
		fmt.Println(err)
		fmt.Println("***************************ERROR DELETE LOG FILE***************************")
	}

	//WRITE NEW FILE LOG AFTER SENT THE OLD THROUGH EMAIL
	WriteFile()

}

func FindFile() (bool, string) {
	dirname := "./"
	files, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Println("********************************FIND FILE ERROR********************************")
		fmt.Println("Error reading directory:", err)
		fmt.Println("********************************FIND FILE ERROR********************************")
		return false, ""
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "apache_logs_") {
			return true, file.Name()
		}
	}
	return false, ""
}

func TestEmailSendAttachment() {

	smtpConfig := env.GetSMTPConfig()
	smtpClient := mail.InitEmail(smtpConfig)
	// Email := []string{"frans.imanuel@visionet.co.id"}
	Email := []string{"frans.imanuel@visionet.co.id", "lishera.prihatni@visionet.co.id", "ari.darmawan@visionet.co.id", "azky.muhtarom@visionet.co.id"}
	if err := smtpClient.Send(Email, nil, nil, "Apache Logs", "text/html", "Apache Logs with txt", []string{"test.txt"}); err != nil {
		fmt.Println("***************************ERROR SEND EMAIL***************************")
		fmt.Println(err)
		fmt.Println("***************************ERROR SEND EMAIL***************************")
	}

}

func CheckFileLength(filename string) int64 {
	f, err := os.Open(filename)
	fi, err := f.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		fmt.Println(err)
	}
	return fi.Size()
}
