package config

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name         string
		environment  map[string]string
		want         *Config
		wantErr      assert.ErrorAssertionFunc
		errorMessage string
	}{
		{
			name: "valid config",
			environment: map[string]string{
				"GODVILLE_GODNAME": "godville-godname",
				"GODVILLE_TOKEN":   "godville-token",
				"MASTODON_SERVER":  "mastodon-server",
				"MASTODON_ID":      "mastodon-id",
				"MASTODON_SECRET":  "mastodon-secret",
			},
			want: &Config{
				GodvilleGodname: "godville-godname",
				GodvilleToken:   "godville-token",
				MastodonServer:  "mastodon-server",
				MastodonID:      "mastodon-id",
				MastodonSecret:  "mastodon-secret",
			},
			wantErr: assert.NoError,
		},
		{
			name: "missing godville godname",
			environment: map[string]string{
				"GODVILLE_TOKEN": "godville-token",
			},
			want:         nil,
			wantErr:      assert.Error,
			errorMessage: "GODVILLE_GODNAME is not set",
		},
		{
			name: "missing godville token",
			environment: map[string]string{
				"GODVILLE_GODNAME": "godville-godname",
			},
			want:         nil,
			wantErr:      assert.Error,
			errorMessage: "GODVILLE_TOKEN is not set",
		},
		{
			name: "no publishers are configured",
			environment: map[string]string{
				"GODVILLE_GODNAME": "godville-godname",
				"GODVILLE_TOKEN":   "godville-token",
			},
			want:         nil,
			wantErr:      assert.Error,
			errorMessage: "no publishers are configured",
		},
		{
			name: "missing mastodon id",
			environment: map[string]string{
				"GODVILLE_GODNAME": "godville-godname",
				"GODVILLE_TOKEN":   "godville-token",
				"MASTODON_SERVER":  "mastodon-server",
			},
			want:         nil,
			wantErr:      assert.Error,
			errorMessage: "MASTODON_ID and MASTODON_SECRET must be set if MASTODON_SERVER is set",
		},
		{
			name: "missing mastodon secret",
			environment: map[string]string{
				"GODVILLE_GODNAME": "godville-godname",
				"GODVILLE_TOKEN":   "godville-token",
				"MASTODON_SERVER":  "mastodon-server",
				"MASTODON_ID":      "mastodon-id",
			},
			want:         nil,
			wantErr:      assert.Error,
			errorMessage: "MASTODON_ID and MASTODON_SECRET must be set if MASTODON_SERVER is set",
		},
	}

	for _, tt := range tests {
		localtt := tt
		t.Run(tt.name, func(t *testing.T) {

			// set environment variables for this test
			for k, v := range localtt.environment {
				t.Setenv(k, v)
			}

			got, err := Load()
			if !localtt.wantErr(t, err, "Load()") {
				return
			}
			if err != nil {
				assert.Equalf(t, localtt.errorMessage, err.Error(), "Load()", localtt.name)
			}
			require.Equalf(t, localtt.want, got, "Load()", localtt.name)
		})
	}
}
