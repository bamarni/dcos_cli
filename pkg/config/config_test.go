package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var cfgStr = `
[core]
ssl_verify = "/home/mesosphere/.dcos/clusters/456934a6-fc1f-4d52-9419-a4c722abe2da/dcos_ca.crt"
dcos_url = "https://dcos.example.com"
dcos_acs_token = "eyJ0eXAiOFJHFADCJhbGciOiJSUzI1NiJ9"

[marathon]
url = "https://marathon.dcos.example.com"

[cluster]
name = "mr-cluster-1835928138"
`

func TestStringAt(t *testing.T) {
	cfg, err := FromString(cfgStr)
	require.NoError(t, err)

	expectedOutputs := []struct {
		key []string
		out  string
	}{
		{[]string{"core", "dcos_url"}, "https://dcos.example.com\n"},
		{[]string{"marathon", "url"}, "https://marathon.dcos.example.com\n"},
		{[]string{"marathon"}, "marathon.url https://marathon.dcos.example.com\n"},
		{[]string{"cluster", "name"}, "mr-cluster-1835928138\n"},
		{[]string{"cluster"}, "cluster.name mr-cluster-1835928138\n"},
	}

	for _, expectedOutput := range expectedOutputs {
		out, err := cfg.StringAt(expectedOutput.key)
		require.NoError(t, err)
		require.Equal(t, expectedOutput.out, out)
	}
}

func TestString(t *testing.T) {
	cfg, err := FromString(cfgStr)
	require.NoError(t, err)

	expectedOutput := `core.dcos_acs_token ********
core.dcos_url https://dcos.example.com
core.ssl_verify /home/mesosphere/.dcos/clusters/456934a6-fc1f-4d52-9419-a4c722abe2da/dcos_ca.crt
cluster.name mr-cluster-1835928138
marathon.url https://marathon.dcos.example.com
`
	require.Equal(t, expectedOutput, cfg.String())
}

func TestSSL(t *testing.T) {
	cfg, err := FromString("[core]\nssl_verify = \"True\"")
	require.NoError(t, err)
	require.Equal(t, false, cfg.Core.SSL.Insecure())

	cfg, err = FromString("[core]\nssl_verify = \"true\"")
	require.NoError(t, err)
	require.Equal(t, false, cfg.Core.SSL.Insecure())

	cfg, err = FromString("[core]\nssl_verify = \"False\"")
	require.NoError(t, err)
	require.Equal(t, true, cfg.Core.SSL.Insecure())

	cfg, err = FromString("[core]\nssl_verify = \"false\"")
	require.NoError(t, err)
	require.Equal(t, true, cfg.Core.SSL.Insecure())

	cfg, err = FromString("[core]\nssl_verify = \"/path/to/ca\"")
	require.NoError(t, err)
	require.Equal(t, false, cfg.Core.SSL.Insecure())
	require.Equal(t, "/path/to/ca", cfg.Core.SSL.CAPath())
}
