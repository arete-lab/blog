package keeper

import (
	"blog/x/blog/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	post.Id = k.GetPostCount(ctx)
	store.Set(k.GetPostIDBytes(post.Id), k.cdc.MustMarshal(&post))
	k.SetPostCount(ctx, post.Id+1)
	return post.Id
}

func (k Keeper) UpdatePost(ctx sdk.Context, post types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	store.Set(k.GetPostIDBytes(post.Id), k.cdc.MustMarshal(&post))
}

func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	store.Set(k.GetPostIDBytes(post.Id), k.cdc.MustMarshal(&post))
}

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	post := store.Get(k.GetPostIDBytes(id))
	if post == nil {
		return types.Post{}, false
	}
	k.cdc.MustUnmarshal(post, &val)
	return val, true
}

func (k Keeper) GetPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) RemovePost(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	store.Delete(k.GetPostIDBytes(id))
}
