package main

import (
	"fmt"
	"go-read-apache2-error-logs/util"
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
	util.CheckFileLength("asd.txt")
	util.CheckFileLength("fgh.txt")
	panic(1)
	// util.TestEmailSendAttachment()
	// panic(1)
	// logs := `103.106.82.174 - - [24/Apr/2024:17:36:40 +0700] "POST /crm-ticket-copy-1/api/v1/user/login HTTP/1.1" 503 1146 "https://innodev.vnetcloud.com/metacrm-internal/login" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0"
	// `
	// fmt.Println("|" + util.ExtractHTTPStatusCode(logs) + "|")
	// panic(1)
	fmt.Println("---Starting---")

	//activate the scheduler
	util.Scheduler()

	//create first file
	util.WriteFile()

	t, err := tail.TailFile("/var/log/apache2/access.log", tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		fmt.Println("***************************ERROR TAILING FILE LOG***************************")
		fmt.Println(err)
		fmt.Println("***************************ERROR TAILING FILE LOG***************************")
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
		if strings.Contains(line.Text, "news") || strings.Contains(line.Text, "1_6_5_prd_mf_") || strings.Contains(line.Text, "1_6_6_prd_mf_web") {
			// if strings.Contains(line.Text, "nobucall-api-v2") || strings.Contains(line.Text, "nobucall-api-report") || strings.Contains(line.Text, "crm-ticket-copy-1") {
			if util.ExtractHTTPStatusCode(line.Text) == "400" || util.ExtractHTTPStatusCode(line.Text) == "500" || util.ExtractHTTPStatusCode(line.Text) == "503" {
				fmt.Println(line.Text)

				_, fname := util.FindFile()
				util.WriteLog(fname, line.Text)
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

// API_PROD_NEWS_GO=https://apigw.vnetcloud.com/news/
// API_PROD_GO=https://apigw.vnetcloud.com/1_6_5_prd_mf_go/
// API_PROD_NODE=https://apigw.vnetcloud.com/1_6_5_prd_mf_mobile/
// API_PROD_WB=https://apigw.vnetcloud.com/1_6_6_prd_mf_web/
