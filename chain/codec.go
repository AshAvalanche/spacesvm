// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

const (
	// codecVersion is the current default codec version
	codecVersion = 0

	// maxSize is 1MB to support large blocks (~9 large key settings)
	maxSize = 1 * units.MiB
)

var codecManager codec.Manager

func init() {
	c := linearcodec.NewDefault()
	codecManager = codec.NewManager(maxSize)
	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterType(&BaseTx{}),
		c.RegisterType(&ClaimTx{}),
		c.RegisterType(&LifelineTx{}),
		c.RegisterType(&SetTx{}),
		c.RegisterType(&Transaction{}),
		c.RegisterType(&StatefulBlock{}),
		c.RegisterType(&PrefixInfo{}),
		codecManager.RegisterCodec(codecVersion, c),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}

func Marshal(source interface{}) ([]byte, error) {
	return codecManager.Marshal(codecVersion, source)
}

func Unmarshal(source []byte, destination interface{}) (uint16, error) {
	return codecManager.Unmarshal(source, destination)
}
