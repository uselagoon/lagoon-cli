package ssh

import (
	"os"
	"path"
	"testing"
)

func TestInteractiveKnownHosts_CreatesFileIfMissing(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Ensure .ssh directory doesn't exist
	sshDir := path.Join(tmpDir, ".ssh")
	if _, err := os.Stat(sshDir); err == nil {
		t.Fatal("Expected .ssh directory to not exist initially")
	}

	// Call InteractiveKnownHosts with a non-existent known_hosts file
	// This should create both the directory and file
	_, _, err := InteractiveKnownHosts(tmpDir, "test.example.com:22", false, true)
	if err != nil {
		t.Fatalf("InteractiveKnownHosts failed: %v", err)
	}

	// Verify .ssh directory was created with correct permissions
	sshInfo, err := os.Stat(sshDir)
	if err != nil {
		t.Fatalf("Expected .ssh directory to be created: %v", err)
	}
	if !sshInfo.IsDir() {
		t.Fatal("Expected .ssh to be a directory")
	}
	if sshInfo.Mode().Perm() != 0700 {
		t.Errorf("Expected .ssh directory permissions to be 0700, got %o", sshInfo.Mode().Perm())
	}

	// Verify known_hosts file was created with correct permissions
	knownHostsPath := path.Join(sshDir, "known_hosts")
	khInfo, err := os.Stat(knownHostsPath)
	if err != nil {
		t.Fatalf("Expected known_hosts file to be created: %v", err)
	}
	if khInfo.IsDir() {
		t.Fatal("Expected known_hosts to be a file, not a directory")
	}
	if khInfo.Mode().Perm() != 0600 {
		t.Errorf("Expected known_hosts file permissions to be 0600, got %o", khInfo.Mode().Perm())
	}
}

func TestInteractiveKnownHosts_HandlesExistingFile(t *testing.T) {
	// Create a temporary directory with existing .ssh/known_hosts
	tmpDir := t.TempDir()
	sshDir := path.Join(tmpDir, ".ssh")

	// Create .ssh directory
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		t.Fatalf("Failed to create .ssh directory: %v", err)
	}

	// Create known_hosts file
	knownHostsPath := path.Join(sshDir, "known_hosts")
	f, err := os.Create(knownHostsPath)
	if err != nil {
		t.Fatalf("Failed to create known_hosts file: %v", err)
	}
	f.Close()

	// Call InteractiveKnownHosts with existing file
	_, _, err = InteractiveKnownHosts(tmpDir, "test.example.com:22", false, true)
	if err != nil {
		t.Fatalf("InteractiveKnownHosts failed with existing file: %v", err)
	}

	// Verify file still exists
	if _, err := os.Stat(knownHostsPath); err != nil {
		t.Fatalf("known_hosts file should still exist: %v", err)
	}
}

func TestInteractiveKnownHosts_IgnoreHostKey(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()

	// Call with ignorehost=true, which should skip all host key checks
	// and not create any files
	_, _, err := InteractiveKnownHosts(tmpDir, "test.example.com:22", true, false)
	if err != nil {
		t.Fatalf("InteractiveKnownHosts failed with ignorehost=true: %v", err)
	}

	// Verify no .ssh directory was created when ignoring host keys
	sshDir := path.Join(tmpDir, ".ssh")
	if _, err := os.Stat(sshDir); err == nil {
		t.Error("Expected .ssh directory to not be created when ignoring host keys")
	}
}
