//go:build integration

package integration

import (
	"testing"

	"github.com/kong/deck/utils"
	"github.com/kong/go-kong/kong"
)

func Test_Dump_CACertificate(t *testing.T) {
	// setup stage
	client, err := getTestClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	tests := []struct {
		name          string
		kongFile      string
		expectedState utils.KongRawState
	}{
		{
			name:     "creates a CA certificate and dump",
			kongFile: "testdata/dump/001-create-a-cacert-and-dump/kong.yaml",
			expectedState: utils.KongRawState{
				CACertificates: []*kong.CACertificate{
					{
						Cert:       cert,
						CertDigest: certDigest,
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			kong.RunWhenKong(t, ">=2.1.4")
			teardown := setup(t)
			defer teardown(t)

			// sync configuration and test Kong entities
			sync(tc.kongFile)
			testKongState(t, client, tc.expectedState, nil)

			// dump configuration to the same state file
			dumpC(tc.kongFile)

			// sync configuration and test Kong entities are unchanged
			sync(tc.kongFile)
			testKongState(t, client, tc.expectedState, nil)
		})
	}
}
