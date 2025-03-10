package test

import (
	"fmt"
	"github.com/ksctl/ksctl/pkg/helpers"
	"os"
	"testing"

	"github.com/ksctl/ksctl/pkg/helpers/consts"
)

func BenchmarkCivoTestingManaged(b *testing.B) {
	if err := os.Setenv(string(consts.KsctlFakeFlag), "1"); err != nil {
		b.Fatalf("Failed to set fake env %v", err)
	}
	StartCloud()

	for i := 0; i < b.N; i++ {
		if err := CivoTestingManaged(); err != nil {
			b.Fatalf("failed, err: %v", err)
		}
	}

	fmt.Println("Cleanup..")
	if err := os.RemoveAll(os.TempDir() + helpers.PathSeparator + "ksctl-black-box-test"); err != nil {
		panic(err)
	}
}

func BenchmarkCivoTestingHA(b *testing.B) {
	if err := os.Setenv(string(consts.KsctlFakeFlag), "1"); err != nil {
		b.Fatalf("Failed to set fake env %v", err)
	}
	StartCloud()

	for i := 0; i < b.N; i++ {
		if err := CivoTestingHA(); err != nil {
			b.Fatalf("failed, err: %v", err)
		}
	}

	fmt.Println("Cleanup..")
	if err := os.RemoveAll(os.TempDir() + helpers.PathSeparator + "ksctl-black-box-test"); err != nil {
		panic(err)
	}
}

func BenchmarkAzureTestingHA(b *testing.B) {
	if err := os.Setenv(string(consts.KsctlFakeFlag), "1"); err != nil {
		b.Fatalf("Failed to set fake env %v", err)
	}
	StartCloud()

	for i := 0; i < b.N; i++ {
		if err := AzureTestingHA(); err != nil {
			b.Fatalf("failed, err: %v", err)
		}
	}

	fmt.Println("Cleanup..")
	if err := os.RemoveAll(os.TempDir() + helpers.PathSeparator + "ksctl-black-box-test"); err != nil {
		panic(err)
	}
}

func BenchmarkAzureTestingManaged(b *testing.B) {
	if err := os.Setenv(string(consts.KsctlFakeFlag), "1"); err != nil {
		b.Fatalf("Failed to set fake env %v", err)
	}
	StartCloud()

	for i := 0; i < b.N; i++ {
		if err := AzureTestingManaged(); err != nil {
			b.Fatalf("failed, err: %v", err)
		}
	}

	fmt.Println("Cleanup..")
	if err := os.RemoveAll(os.TempDir() + helpers.PathSeparator + "ksctl-black-box-test"); err != nil {
		panic(err)
	}
}

func BenchmarkLocalTestingManaged(b *testing.B) {
	if err := os.Setenv(string(consts.KsctlFakeFlag), "1"); err != nil {
		b.Fatalf("Failed to set fake env %v", err)
	}
	StartCloud()

	for i := 0; i < b.N; i++ {
		if err := LocalTestingManaged(); err != nil {
			b.Fatalf("failed, err: %v", err)
		}
	}

	fmt.Println("Cleanup..")
	if err := os.RemoveAll(os.TempDir() + helpers.PathSeparator + "ksctl-black-box-test"); err != nil {
		panic(err)
	}
}
