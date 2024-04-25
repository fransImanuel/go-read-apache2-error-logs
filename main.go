package main

import (
	"fmt"
	"go-read-apache2-error-logs/util/env"
	"go-read-apache2-error-logs/util/mail"
	"strings"

	"github.com/nxadm/tail"
)

/*
	How to Run it in vm

git clone https://github.com/fransImanuel/go-read-apache2-error-logs.git
cd go-read-apache2-error-logs
go build -o watch_logs ./main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o watch_logs ./main.go
chmod +x go.mod
chmod +x go.sum
chmod +x main.go
chmod -R a+rwx go-read-apache2-error-logs/
chmod +x crm-ticket-scheduler

./watch_logs > program_log.txt 2>&1 &
*/
func main() {
	// logs := `103.106.82.174 - - [24/Apr/2024:17:36:40 +0700] "POST /crm-ticket-copy-1/api/v1/user/login HTTP/1.1" 503 1146 "https://innodev.vnetcloud.com/metacrm-internal/login" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0"
	// `
	// fmt.Println("|" + extractHTTPStatusCode(logs) + "|")
	// panic(1)
	fmt.Println("---Starting---")

	smtpConfig := env.GetSMTPConfig()
	smtpClient := mail.InitEmail(smtpConfig)
	Email := []string{"frans.imanuel@visionet.co.id", "lishera.prihatni@visionet.co.id", "ari.darmawan@visionet.co.id", "azky.muhtarom@visionet.co.id" /*, "fransimanuel99@gmail.com" */}

	// panic(1)
	// fileLocation :=
	t, err := tail.TailFile("/var/log/apache2/access.log", tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
		// selain 200 dan 404
		if strings.Contains(line.Text, "nobucall-api-v2") || strings.Contains(line.Text, "nobucall-api-report") || strings.Contains(line.Text, "crm-ticket-copy-1") {

			if extractHTTPStatusCode(line.Text) == "400" || extractHTTPStatusCode(line.Text) == "500" || extractHTTPStatusCode(line.Text) == "503" {
				// fmt.Println("***************************GOTTEM***************************")
				fmt.Println(line.Text)
				// fmt.Println("***************************GOTTEM***************************")
				if err := smtpClient.Send(Email, nil, nil, "Apache Logs", "text/html", line.Text, nil); err != nil {
					fmt.Println("***************************ERROR***************************")
					fmt.Println(err)
					fmt.Println("***************************ERROR***************************")
				}
			}
		}

	}
	fmt.Println("---Finished---")
}

// 400, 500, 503
//
//	if err := smtpClient.Send(Email, nil, nil, "heading", "text/html", "test email123", nil); err != nil {
//		panic(err)
//	}
func extractHTTPStatusCode(logEntry string) string {
	// Split the log entry by space
	parts := strings.Split(logEntry, " ")
	// The HTTP status code is the 9th element in the split parts
	return parts[8]
}
