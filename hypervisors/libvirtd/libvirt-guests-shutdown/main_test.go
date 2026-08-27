// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	"errors"
	"slices"
	"strings"
	"testing"
)

type virshReply struct {
	after  func()
	output string
	err    error
	args   []string
}

func scriptedVirsh(t *testing.T, replies []virshReply) virshCommand {
	t.Helper()

	index := 0
	t.Cleanup(func() {
		if index != len(replies) {
			t.Errorf("called %d of %d expected virsh commands", index, len(replies))
		}
	})

	return func(_ context.Context, args ...string) (string, error) {
		if index >= len(replies) {
			t.Fatalf("unexpected virsh command: %v", args)
		}

		reply := replies[index]
		index++

		if !slices.Equal(reply.args, args) {
			t.Errorf("unexpected virsh arguments: got %v, expected %v", args, reply.args)
		}

		if reply.after != nil {
			reply.after()
		}

		return reply.output, reply.err
	}
}

func TestShutdownGuests(t *testing.T) {
	run := scriptedVirsh(t, []virshReply{
		{args: []string{"list", "--uuid"}, output: "domain-a\ndomain-b\n"},
		{args: []string{"managedsave", "--running", "domain-a"}},
		{args: []string{"managedsave", "--running", "domain-b"}, err: errors.New("transient domain")},
		{args: []string{"list", "--uuid"}, output: "domain-b\n"},
		{args: []string{"destroy", "domain-b"}},
		{args: []string{"list", "--uuid"}},
	})

	if err := shutdownGuestsWith(t.Context(), t.Context(), run); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}
}

func TestShutdownGuestsReportsRemainingDomains(t *testing.T) {
	run := scriptedVirsh(t, []virshReply{
		{args: []string{"list", "--uuid"}, output: "domain-a\n"},
		{args: []string{"managedsave", "--running", "domain-a"}, err: errors.New("save failed")},
		{args: []string{"list", "--uuid"}, output: "domain-a\n"},
		{args: []string{"destroy", "domain-a"}, err: errors.New("destroy failed")},
		{args: []string{"list", "--uuid"}, output: "domain-a\n"},
	})

	err := shutdownGuestsWith(t.Context(), t.Context(), run)
	if err == nil {
		t.Fatal("expected shutdown to fail")
	}

	for _, expected := range []string{
		"failed to stop domain domain-a: destroy failed",
		"domains still active after shutdown: domain-a",
	} {
		if !strings.Contains(err.Error(), expected) {
			t.Errorf("error %q does not contain %q", err, expected)
		}
	}
}

func TestShutdownGuestsDestroysDomainsAfterManagedSaveTimeout(t *testing.T) {
	managedSaveCtx, cancel := context.WithCancel(t.Context())
	run := scriptedVirsh(t, []virshReply{
		{args: []string{"list", "--uuid"}, output: "domain-a\n", after: cancel},
		{args: []string{"list", "--uuid"}, output: "domain-a\n"},
		{args: []string{"destroy", "domain-a"}},
		{args: []string{"list", "--uuid"}},
	})

	if err := shutdownGuestsWith(t.Context(), managedSaveCtx, run); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}
}

func TestShutdownGuestsStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	run := scriptedVirsh(t, nil)

	cancel()

	if err := shutdownGuestsWith(ctx, ctx, run); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
}
