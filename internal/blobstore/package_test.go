// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package blobstore_test

import (
	"testing"

	jujutesting "github.com/juju/testing"
)

func TestPackage(t *testing.T) {
	jujutesting.MgoTestPackage(t, nil)
}
