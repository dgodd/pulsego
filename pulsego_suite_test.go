package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPulsego(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pulsego Suite")
}
