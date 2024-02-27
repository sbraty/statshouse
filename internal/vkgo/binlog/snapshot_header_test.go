// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package binlog

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleWriteRead(t *testing.T) {
	hdr := SnapshotHeader{
		ClusterId:    "cluster",
		ShardId:      "shard",
		SnapshotMeta: "binarydata",
		Dependencies: []SnapshotDependency{{
			ClusterId:     "cluster",
			ShardId:       "shard",
			PayloadOffset: 100500,
		}},
		PayloadOffset:    100600,
		EngineVersion:    "version",
		CreationTimeNano: 10001234,
	}
	hdr.SetControlMeta("binarydata2")

	w := WriteSnapshotHeader(nil, &hdr)
	_, hdr2, _, err := ReadSnapshotHeader(w)
	require.NoError(t, err)
	require.Equal(t, hdr, hdr2)
}

func TestCorrupt(t *testing.T) {
	hdr := SnapshotHeader{
		ClusterId:    "cluster",
		ShardId:      "shard",
		SnapshotMeta: "binarydata",
	}

	w := WriteSnapshotHeader(nil, &hdr)
	w[2]++ // data corruption

	_, _, _, err := ReadSnapshotHeader(w)
	require.Error(t, err)
}

func TestReadHeader(t *testing.T) {
	// hexdump of snapshot written by pmemcache-disk
	hdrDataHex := `741b0d1de00000000000000001000000053134373934000003305f323804
2c793201c00000368e0b0f00000000ffc4765b54d19c1731a16a3cceaea1
991052c334badea4373a0a330f0000000027fc0000000000000000000100
000000000000053134373934000003305f32bdabce0e00000000a39d0b0f
000000002831643233613036343933643838373161656638303530303462
616339656430633834643465376639000000c8172fb8e4649d1738042c79
3201c00000368e0b0f00000000ffc4765b54d19c1731a16a3cceaea19910
52c334badea4373a0a330f0000000027fc000000000000000000a213c951
ec35124bab940ca6bdaab3e114b6f84b1000000000000000000000001fc5
6c65000000007b0a93276eb47c656bcc08d32fcabd2da2bfa00bb7961050`
	hdrDataHex = strings.ReplaceAll(hdrDataHex, "\n", "")
	expectSnapMeta, _ := hex.DecodeString("042c793201c00000368e0b0f00000000ffc4765b54d19c1731a16a3cceaea1991052c334badea4373a0a330f0000000027fc000000000000")

	hdrExpect := SnapshotHeader{
		FieldsMask:   0,
		ClusterId:    "14794",
		ShardId:      "0_2",
		SnapshotMeta: string(expectSnapMeta),
		Dependencies: []SnapshotDependency{{
			ClusterId:     "14794",
			ShardId:       "0_2",
			PayloadOffset: 248425405,
		}},
		PayloadOffset:    252419491,
		EngineVersion:    "1d23a06493d8871aef805004bac9ed0c84d4e7f9",
		CreationTimeNano: 1701627167721330632,
	}
	hdrExpect.SetControlMeta(string(expectSnapMeta))

	hdrData, err := hex.DecodeString(hdrDataHex)
	require.NoError(t, err)

	_, hdr, _, err := ReadSnapshotHeader(hdrData)
	require.NoError(t, err)

	require.Equal(t, hdrExpect, hdr)
}
