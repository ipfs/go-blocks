package bstest

// import (
// 	"bytes"
// 	"testing"
// 	"time"

// 	blocks "github.com/ipfs/go-blocks"
// 	. "github.com/ipfs/go-blocks/blockservice"
// 	offline "github.com/ipfs/go-blocks/blockservice/exchange/offline"
// 	blockstore "github.com/ipfs/go-blocks/blockstore"
// 	blocksutil "github.com/ipfs/go-blocks/blocksutil"
// 	hash "github.com/ipfs/go-blocks/hash"
// 	key "github.com/ipfs/go-blocks/key"

// 	ds "github.com/ipfs/go-blocks/Godeps/_workspace/src/github.com/jbenet/go-datastore"
// 	dssync "github.com/ipfs/go-blocks/Godeps/_workspace/src/github.com/jbenet/go-datastore/sync"
// 	"github.com/ipfs/go-blocks/Godeps/_workspace/src/golang.org/x/net/context"
// )

// func TestBlocks(t *testing.T) {
// 	bstore := blockstore.NewBlockstore(dssync.MutexWrap(ds.NewMapDatastore()))
// 	bs, err := New(bstore, offline.Exchange(bstore))
// 	if err != nil {
// 		t.Error("failed to construct block service", err)
// 		return
// 	}
// 	defer bs.Close()

// 	b := blocks.NewBlock([]byte("beep boop"))
// 	h := hash.Hash([]byte("beep boop"))
// 	if !bytes.Equal(b.Multihash, h) {
// 		t.Error("Block Multihash and data multihash not equal")
// 	}

// 	if b.Key() != key.Key(h) {
// 		t.Error("Block key and data multihash key not equal")
// 	}

// 	k, err := bs.AddBlock(b)
// 	if err != nil {
// 		t.Error("failed to add block to BlockService", err)
// 		return
// 	}

// 	if k != b.Key() {
// 		t.Error("returned key is not equal to block key", err)
// 	}

// 	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
// 	b2, err := bs.GetBlock(ctx, b.Key())
// 	if err != nil {
// 		t.Error("failed to retrieve block from BlockService", err)
// 		return
// 	}

// 	if b.Key() != b2.Key() {
// 		t.Error("Block keys not equal.")
// 	}

// 	if !bytes.Equal(b.Data, b2.Data) {
// 		t.Error("Block data is not equal.")
// 	}
// }

// func TestGetBlocksSequential(t *testing.T) {
// 	var servs = Mocks(t, 4)
// 	for _, s := range servs {
// 		defer s.Close()
// 	}
// 	bg := blocksutil.NewBlockGenerator()
// 	blks := bg.Blocks(50)

// 	var keys []key.Key
// 	for _, blk := range blks {
// 		keys = append(keys, blk.Key())
// 		servs[0].AddBlock(blk)
// 	}

// 	t.Log("one instance at a time, get blocks concurrently")

// 	for i := 1; i < len(servs); i++ {
// 		ctx, _ := context.WithTimeout(context.TODO(), time.Second*50)
// 		out := servs[i].GetBlocks(ctx, keys)
// 		gotten := make(map[key.Key]*blocks.Block)
// 		for blk := range out {
// 			if _, ok := gotten[blk.Key()]; ok {
// 				t.Fatal("Got duplicate block!")
// 			}
// 			gotten[blk.Key()] = blk
// 		}
// 		if len(gotten) != len(blks) {
// 			t.Fatalf("Didnt get enough blocks back: %d/%d", len(gotten), len(blks))
// 		}
// 	}
// }
