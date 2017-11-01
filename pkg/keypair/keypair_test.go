// Copyright 2017 The nem-toolchain project authors. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package keypair

import (
	"encoding/hex"
	"testing"

	"bytes"

	"github.com/nem-toolchain/nem-toolchain/pkg/core"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ed25519"
)

func TestKeyPair_Gen(t *testing.T) {
	pr, _ := hex.DecodeString("2c52aee96f0e30f21c86b3fab7a18e927f579618818e8148e7ded1e01875ef0b")
	exp, _ := hex.DecodeString("9d1e9d01ab916dbdde0e76ba43df2246575d637db0bca090f46c1abce19a43e3")
	act, _, _ := ed25519.GenerateKey(bytes.NewReader(pr))
	assert.Equal(t, exp, []byte(act))
}

func TestKeyPair_Address_mainnet(t *testing.T) {
	pub, _ := hex.DecodeString("9d1e9d01ab916dbdde0e76ba43df2246575d637db0bca090f46c1abce19a43e3")
	addr, _ := ParseAddress("NAKTWAOYSE5F3J2FJWOXR56UTQLIOUXRJLBJ7CBF")
	assert.Equal(t, addr, KeyPair{Public: pub}.Address(core.Mainnet))
}

func TestKeyPair_Address_testnet(t *testing.T) {
	pub, _ := hex.DecodeString("4fe5efd97360bc8a32ec105d419222eeb714e6d06fd8b895a5eedda2b0edf931")
	addr, _ := ParseAddress("TA6XFSJYZYAIYP7FL7X2RL63647FRMB65YC6CO3G")
	assert.Equal(t, addr, KeyPair{Public: pub}.Address(core.Testnet))
}
