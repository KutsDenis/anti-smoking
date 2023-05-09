package config

import (
	"github.com/KutsDenis/logpac"
	"testing"
)

func TestGetConfig(t *testing.T) {
	l := logpac.GetLogger()

	GetConfig(l)

	if Get == nil {
		t.Error("Конфигурация не выгружена либо не файл не содержит параметров")
	}
}
