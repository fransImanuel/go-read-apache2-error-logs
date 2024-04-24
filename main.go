package main

import (
	"fmt"

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
	fmt.Println("---Starting---")

	// smtpConfig := env.GetSMTPConfig()
	// smtpClient := mail.InitEmail(smtpConfig)
	// Email := []string{"frans.imanuel@visionet.co.id", "lishera.prihatni@visionet.co.id", "ari.darmawan@visionet.co.id", "azky.muhtarom@visionet.co.id" /*, "fransimanuel99@gmail.com" */}

	// panic(1)
	fileLocation := "/var/log/apache2/access.log"
	t, err := tail.TailFile(fileLocation, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
		// selain 200 dan 404
		// if strings.Contains(line.Text, "nobucall-api-v2") || strings.Contains(line.Text, "nobucall-api-report") {
		// 	if !strings.Contains(line.Text, " 200 ") && !strings.Contains(line.Text, " 404 ") {
		// 		if err := smtpClient.Send(Email, nil, nil, "Apache Logs", "text/html", line.Text, nil); err != nil {
		// 			fmt.Println("=============================ERROR==================================")
		// 			fmt.Println(err)
		// 			fmt.Println("=============================ERROR==================================")
		// 		}
		// 	}
		// }

	}
	fmt.Println("---Finished---")
}

// if err := smtpClient.Send(Email, nil, nil, "heading", "text/html", "test email123", nil); err != nil {
// 	panic(err)
// }
