//go:build integration

package integration

import (
	"testing"

	"github.com/kong/deck/utils"
	"github.com/kong/go-kong/kong"
)

var (
	cert = kong.String(`-----BEGIN CERTIFICATE-----
MIICtjCCAZ6gAwIBAgIJAPQ/YU66PeQeMA0GCSqGSIb3DQEBCwUAMBAxDjAMBgNV
BAMMBUhlbGxvMB4XDTIyMDIyNDE5NDcxOFoXDTIyMDMyNjE5NDcxOFowEDEOMAwG
A1UEAwwFSGVsbG8wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDSs/dh
/DT2irLYZQAQFMxBK2Jd3Eak0L3TDxlubNEuwoNNi/t+M1u9l2tP2eA3hjrdW2UZ
XslaOnSQOsoHDDELwzoPEHOoSbwako5cxdHuDUEbwtf/bcjMQdxQM+fCuGKp6TyG
dFDXiFsauNBG4oAXR+X40gGUldSCIEupFm2t8VbqYK57nMBuJzkwy7zAEcO+kL2c
ajf/SZpMKVnrqqmiSNJnqt5m6sTsHZUkw9N63c9g2rCTvelmqqWCsqKYSYlxcjVc
4qQU2/zKKKNAQCWvi5tQBhJ12IN5BxFmV8Hj8TOG7o8/PPrN65ukuJ7WIn3OGUXs
FPhkG4JEOdLe5NJ1AgMBAAGjEzARMA8GA1UdEwQIMAYBAf8CAQAwDQYJKoZIhvcN
AQELBQADggEBACGyIg3m8Ih0DUSyKOBnpFjA/+Wf6kdn1uBgsiac3OMiz+YxnmeI
v2kK0txOiX86fzzPvifrA5kx6VzZn2PQnL112Y2kf2SfOdWfPk8+OuhNJgENMiVu
F04Tkj2rP0YwzzwUYwZBnfuOfEeHOANr2HrdqdwGCnNFIad/gRR+ORRT0mcRkMnK
J5BQ/1zaNukb9T9yyKWjgkH+eO/xXYI0kR5BpCh8Ok9RXBeXNeN1B7qwvjetrKfm
WVai1ncDPVwAHyCWa8th8Zl5QAd5scFT7N+om5niSNsz0CysJ1n4XJC5AqP8gpJk
tPtfJJclgHG7F1t76rsDstUMX8p+h0/t27U=
-----END CERTIFICATE-----`)
	certDigest = kong.String("bdcd2b905add88c20fa41f9844b36914c0dcd024a462240fc8face3901468149")
)

func Test_Reset_CACertificate(t *testing.T) {
	// setup stage
	client, err := getTestClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	tests := []struct {
		name                   string
		kongFile               string
		expectedStateAfterSync utils.KongRawState
	}{
		{
			name:     "creates a CA certificate and reset",
			kongFile: "testdata/reset/001-create-a-cacert-and-reset/kong.yaml",
			expectedStateAfterSync: utils.KongRawState{
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

			sync(tc.kongFile)
			testKongState(t, client, tc.expectedStateAfterSync, nil)

			reset(t)
			testKongState(t, client, utils.KongRawState{}, nil)
		})
	}
}
