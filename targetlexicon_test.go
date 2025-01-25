package targetlexicon

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerialize(t *testing.T) {
	pt := func(s string) Triple {
		t2, ok := ParseTriple(s)
		require.True(t, ok, "could not parse triple")
		return t2
	}
	triples := []Triple{
		pt("x86_64-unknown-linux-gnu"),
		pt("i686-pc-windows-gnu"),
	}

	json, err := json.Marshal(triples)
	assert.NoError(t, err)
	assert.Equal(t, `["x86_64-unknown-linux-gnu","i686-pc-windows-gnu"]`, string(json))
}

func TestDeserialize(t *testing.T) {
	pt := func(s string) Triple {
		t2, ok := ParseTriple(s)
		require.True(t, ok, "could not parse triple")
		return t2
	}
	triples := []Triple{
		pt("x86_64-unknown-linux-gnu"),
		pt("i686-pc-windows-gnu"),
	}

	var vals []Triple
	err := json.Unmarshal([]byte(`["x86_64-unknown-linux-gnu","i686-pc-windows-gnu"]`), &vals)
	assert.NoError(t, err)
	assert.True(t, cmp.Equal(triples, vals))
}
