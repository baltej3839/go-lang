package main

import (
	"fmt"
	"io"
)

type AuditLogger interface {
	Log(int)
}

type FileAudit struct {
	fileLocation string
}

type DatabaseAudit struct {
	DatabaseLocation string
}

type RemoteAudit struct {
	RemoteLocation string
}

func (f *FileAudit) Log(amount int) {
	fmt.Printf("amount transferred was %s , is logged in fileAudit log", amount)
}

func NewFileAudit() *FileAudit {
	return &FileAudit{fileLocation: "asdsadsad"}
}

func (r *DatabaseAudit) Log(amount int) {
	fmt.Printf("amount transferred was %s , is logged in database log", amount)
}

func NewDatabaseAudit() *DatabaseAudit {
	return &DatabaseAudit{DatabaseLocation: "asdsadsad"}
}

func (r *RemoteAudit) Log(amount int) {
	fmt.Printf("amount transferred was %s , is logged in remote log", amount)
}


func NewRemoteAudit() *RemoteAudit {
	return &RemoteAudit{RemoteLocation: "asdsadsad"}
}


type TransferService struct {
	amountTransfered int
	auditLogger      []AuditLogger
}

func (t *TransferService) SendMoney(amount int) {
	fmt.Print("amount transferred and logged")
	for _,v:=range t.auditLogger {
		v.Log(amount)
	}
}

func NewTransferService(f *FileAudit, d *DatabaseAudit, r *RemoteAudit) *TransferService {
	return &TransferService{amountTransfered: 0, auditLogger:[]AuditLogger{f,d,r}}
}

func main() {
	t:=200
	r:=NewRemoteAudit()
	f:=NewFileAudit()
	d:=NewDatabaseAudit()
	tsvc:=NewTransferService(f,d,r)

 
	tsvc.SendMoney(t)
}