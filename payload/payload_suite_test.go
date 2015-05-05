package payload_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPayload(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payload Suite")
}
