package controlador_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestControlador(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controlador Suite")
}
