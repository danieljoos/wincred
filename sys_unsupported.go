// +build !windows

package wincred

import "errors"

func sysCredRead(...interface{}) (*Credential, error) {
	return nil, errors.New("Operation not supported")
}

func sysCredWrite(...interface{}) error {
	return errors.New("Operation not supported")
}

func sysCredDelete(...interface{}) error {
	return errors.New("Operation not supported")
}

func sysCredEnumerate(...interface{}) ([]*Credential, error) {
	return nil, errors.New("Operation not supported")
}
