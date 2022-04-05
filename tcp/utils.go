package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func readLen(r *bufio.Reader) (int, error) {
	temp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	length, err := strconv.Atoi(strings.TrimSpace(temp))
	if err != nil {
		return 0, err
	}
	return length, nil
}

func (s *Server) readKey(r *bufio.Reader) (string, error) {
	keyLen, err := readLen(r)
	if err != nil {
		return "", err
	}
	key := make([]byte, keyLen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", err
	}
	return string(key), nil
}
func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	keyLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	valLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	key := make([]byte, keyLen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", nil, err
	}
	val := make([]byte, valLen)
	_, err = io.ReadFull(r, val)
	if err != nil {
		return "", nil, err
	}
	return string(key), val, nil
}
func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		response := fmt.Sprintf("-%d ", len(errString)) + errString
		_, err = conn.Write([]byte(response))
		return err
	}
	response := fmt.Sprintf("%d ", len(value)) + string(value)
	_, err = conn.Write([]byte(response))
	return err
}
