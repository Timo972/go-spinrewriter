package spinrewriter_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/timo972/spinrewriter"
)

func TestMain(m *testing.M) {
	godotenv.Load()

	if _, ok := os.LookupEnv("SPINREWRITER_EMAIL"); !ok {
		log.Fatalln("SPINREWRITER_EMAIL is not set")
	}

	if _, ok := os.LookupEnv("SPINREWRITER_API_KEY"); !ok {
		log.Fatalln("SPINREWRITER_API_KEY is not set")
	}

	code := m.Run()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	client := spinrewriter.New("email", "api_key")

	if client == nil {
		t.Error("Expected client to be instance of spinrewriter.Client")
	}
}

func TestClient_Quota(t *testing.T) {
	client := spinrewriter.New(os.Getenv("SPINREWRITER_EMAIL"), os.Getenv("SPINREWRITER_API_KEY"))

	quota, err := client.Quota()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Quota:", quota)
}

func TestClient_Spintax(t *testing.T) {
	client := spinrewriter.New(os.Getenv("SPINREWRITER_EMAIL"), os.Getenv("SPINREWRITER_API_KEY"))

	spintax, err := client.Spintax("This is a seo sentence. Keep this tag.", spinrewriter.WithProtectedTerms("seo", "this"))
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Spintax: %+v\n", spintax)
}

func ExampleClient_Quota() {
	client := spinrewriter.New(os.Getenv("spinrewriter_EMAIL"), os.Getenv("spinrewriter_API_KEY"))

	quota, err := client.Quota()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Quota: %+v\n", quota)
}

func ExampleClient_Spintax() {
	client := spinrewriter.New(os.Getenv("spinrewriter_EMAIL"), os.Getenv("spinrewriter_API_KEY"))

	// Spin text with protected terms
	spintax, err := client.Spintax("This is a seo sentence. Keep this tag.", spinrewriter.WithProtectedTerms("seo", "this"))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Spintax: %+v\n", spintax)
}
