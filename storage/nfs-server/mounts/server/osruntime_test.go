// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadExportsIncludesDropIns(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "exports")
	if err := os.WriteFile(path, []byte("/srv *(rw)\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(path+".d", 0o700); err != nil {
		t.Fatal(err)
	}
	for name, contents := range map[string]string{
		"nfs.exports":     "/var/mnt/nfs *(rw,mountpoint)\n",
		"ignored.txt":     "/ignored *(rw,mountpoint)\n",
		".hidden.exports": "/hidden *(rw,mountpoint)\n",
	} {
		if err := os.WriteFile(filepath.Join(path+".d", name), []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}

	contents, err := readExports(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(contents), "/srv *(rw)") || !strings.Contains(string(contents), "/var/mnt/nfs *(rw,mountpoint)") {
		t.Fatalf("missing exports: %q", contents)
	}
	if strings.Contains(string(contents), "/ignored") || strings.Contains(string(contents), "/hidden") {
		t.Fatalf("included ignored export files: %q", contents)
	}

	if err := os.WriteFile(filepath.Join(path+".d", "nfs.exports"), []byte("/var/mnt/new *(rw,mountpoint)\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	updated, err := readExports(path)
	if err != nil || string(updated) == string(contents) {
		t.Fatalf("drop-in update not detected: %q, %v", updated, err)
	}
}
