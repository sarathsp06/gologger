package logger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGologger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gologger Suite")
}
