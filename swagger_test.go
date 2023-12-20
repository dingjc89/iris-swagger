package irisSwagger

import (
	_ "github.com/dingjc89/iris-swagger/example/docs"
	"github.com/stretchr/testify/assert"
	"github.com/swaggo/swag"
	"net/http"
	"os"
	"testing"
)

type mockedSwag struct {
}

func (s *mockedSwag) ReadDoc() string {
	content, _ := os.ReadFile("swagger.json")
	return string(content)
}

type httpWriter struct {
}

func (h httpWriter) Header() http.Header {
	return http.Header{}
}

func (h httpWriter) Write(bytes []byte) (int, error) {
	return len(bytes), nil
}

func (h httpWriter) WriteHeader(statusCode int) {
}

func TestURL(t *testing.T) {
	var cfg Config
	expected := "https://github.com/swagger/http-swagger"
	URL(expected)(&cfg)
	assert.Equal(t, expected, cfg.URLs[0])
}

func TestDeepLinking(t *testing.T) {
	var cfg Config
	expected := true
	DeepLinking(expected)(&cfg)
	assert.Equal(t, expected, cfg.DeepLinking)
}

func TestDocExpansion(t *testing.T) {
	var cfg Config
	expected := "https://github.com/swaggo/docs"
	DocExpansion(expected)(&cfg)
	assert.Equal(t, expected, cfg.DocExpansion)
}

func TestDomID(t *testing.T) {
	var cfg Config
	expected := "swagger-ui"
	DomID(expected)(&cfg)
	assert.Equal(t, expected, cfg.DomID)
}

func TestInstanceName(t *testing.T) {
	var cfg Config

	expected := "custom-instance-name"
	InstanceName(expected)(&cfg)
	assert.Equal(t, expected, cfg.InstanceName)

	newCfg := newConfig(InstanceName(""))
	assert.Equal(t, swag.Name, newCfg.InstanceName)
}

func TestPersistAuthorization(t *testing.T) {
	var cfg Config
	expected := true
	PersistAuthorization(expected)(&cfg)
	assert.Equal(t, expected, cfg.PersistAuthorization)
}

func TestOAuth(t *testing.T) {
	var cfg Config
	expected := OAuthConfig{
		ClientId: "my-client-id",
		Realm:    "my-realm",
		AppName:  "My App Name",
	}
	OAuth(&expected)(&cfg)
	assert.Equal(t, expected.ClientId, cfg.OAuth.ClientId)
	assert.Equal(t, expected.Realm, cfg.OAuth.Realm)
	assert.Equal(t, expected.AppName, cfg.OAuth.AppName)
}
