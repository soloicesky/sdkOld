package BEA

import (
	"errors"
)

var (
	CONN_ERR error = errors.New("fail to connect to host")   //连接失败
	SEND_ERR error = errors.New("fail to send data to host") //发送失败
	RECV_ERR error = errors.New("fail to recv from host")    //接收失败
)

const (
	//bindo generated error code
	BINDO_CONN_ERR BEACode = "BINDO_CONN_ERR"
	BINDO_SEND_ERR BEACode = "BINDO_SEND_ERR"
	BINDO_RECV_ERR BEACode = "BINDO_RECV_ERR"
	BINDO_COMM_ERR BEACode = "BINDO_COMM_ERR"
)
