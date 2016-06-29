package providers

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/camptocamp/conplicity/lib"
	"github.com/docker/engine-api/types"
)

func TestGetProvider(t *testing.T) {
	var dir, expected, got string
	var p Provider

	// Test PostgreSQL detection
	expected = "PostgreSQL"
	dir, _ = ioutil.TempDir("", "test_get_provider_postgresql")
	defer os.RemoveAll(dir)
	ioutil.WriteFile(dir+"/PG_VERSION", []byte{}, 0644)

	p = GetProvider(&conplicity.Conplicity{}, &types.Volume{
		Mountpoint: dir,
	})
	got = p.GetName()
	if got != expected {
		t.Fatalf("Expected provider %s, got %s", expected, got)
	}

	// Test MySQL detection
	expected = "MySQL"
	dir, _ = ioutil.TempDir("", "test_get_provider_mysql")
	defer os.RemoveAll(dir)
	os.Mkdir(dir+"/mysql", 0755)

	p = GetProvider(&conplicity.Conplicity{}, &types.Volume{
		Mountpoint: dir,
	})
	got = p.GetName()
	if got != expected {
		t.Fatalf("Expected provider %s, got %s", expected, got)
	}

	// Test OpenLDAP detection
	expected = "OpenLDAP"
	dir, _ = ioutil.TempDir("", "test_get_provider_openldap")
	defer os.RemoveAll(dir)
	ioutil.WriteFile(dir+"/DB_CONFIG", []byte{}, 0644)

	p = GetProvider(&conplicity.Conplicity{}, &types.Volume{
		Mountpoint: dir,
	})
	got = p.GetName()
	if got != expected {
		t.Fatalf("Expected provider %s, got %s", expected, got)
	}

	// Test Default detection
	expected = "Default"
	dir, _ = ioutil.TempDir("", "test_get_provider_default")
	defer os.RemoveAll(dir)

	p = GetProvider(&conplicity.Conplicity{}, &types.Volume{
		Mountpoint: dir,
	})
	got = p.GetName()
	if got != expected {
		t.Fatalf("Expected provider %s, got %s", expected, got)
	}
}

func TestBaseGetHandler(t *testing.T) {
	expected := ""

	p := &BaseProvider{
		handler: &conplicity.Conplicity{},
	}
	got := p.GetHandler().Hostname
	if expected != got {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}

func TestBaseGetVolume(t *testing.T) {
	got := (&BaseProvider{}).GetVolume()
	if got != nil {
		t.Fatalf("Expected to get nil, got %s", got)
	}
}

func TestBaseGetBackupDir(t *testing.T) {
	got := (&BaseProvider{}).GetBackupDir()
	if got != "" {
		t.Fatalf("Expected to get nil, got %s", got)
	}
}
