package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	age := 10
	if age < 18 {
		fmt.Println("You are not old enough to vote.")
	} else {
		fmt.Println("You are old enough to vote.")
	}

	for {
		logFile, err := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed to open audit log file: %v", err)
		}

		auditLogger := log.New(logFile, "AUDIT: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)

		userID := "user123"
		ipAddress := "192.168.1.100"
		resource := "/api/v1/users/456"

		auditLogger.Printf("event_type=user_login, user_id=%s, ip_address=%s, status=success\n", userID, ipAddress)

		action := "GET"
		status := "200 OK"
		auditLogger.Printf("event_type=resource_access, user_id=%s, ip_address=%s, action=%s, resource=%s, status=%s\n", userID, ipAddress, action, resource, status)

		failedUserID := "unknown_user"
		auditLogger.Printf("event_type=login_attempt, user_id=%s, ip_address=%s, status=failed, reason=invalid_credentials\n", failedUserID, ipAddress)

		fmt.Println("Audit logs written to audit.log")
	}
}
