package distribution

import (
	"github.com/irisnet/irishub/modules/distribution/keeper"
	"github.com/irisnet/irishub/modules/distribution/tags"
	"github.com/irisnet/irishub/modules/distribution/types"
	sdk "github.com/irisnet/irishub/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		// NOTE msg already has validate basic run
		switch msg := msg.(type) {
		case types.MsgSetWithdrawAddress:
			return handleMsgModifyWithdrawAddress(ctx, msg, k)
		case types.MsgWithdrawDelegatorRewardsAll:
			return handleMsgWithdrawDelegatorRewardsAll(ctx, msg, k)
		case types.MsgWithdrawDelegatorReward:
			return handleMsgWithdrawDelegatorReward(ctx, msg, k)
		case types.MsgWithdrawValidatorRewardsAll:
			return handleMsgWithdrawValidatorRewardsAll(ctx, msg, k)
		default:
			return sdk.ErrTxDecode("invalid message parse in distribution module").Result()
		}
	}
}

//_____________________________________________________________________

// These functions assume everything has been authenticated,
// now we just perform action and save

func handleMsgModifyWithdrawAddress(ctx sdk.Context, msg types.MsgSetWithdrawAddress, k keeper.Keeper) sdk.Result {

	k.SetDelegatorWithdrawAddr(ctx, msg.DelegatorAddr, msg.WithdrawAddr)

	tags := sdk.NewTags(
		tags.Action, tags.ActionModifyWithdrawAddress,
		tags.Delegator, []byte(msg.DelegatorAddr.String()),
	)
	return sdk.Result{
		Tags: tags,
	}
}

func handleMsgWithdrawDelegatorRewardsAll(ctx sdk.Context, msg types.MsgWithdrawDelegatorRewardsAll, k keeper.Keeper) sdk.Result {

	reward, withdrawTags := k.WithdrawDelegationRewardsAll(ctx, msg.DelegatorAddr)

	resultTags := sdk.NewTags(
		tags.Action, tags.ActionWithdrawDelegatorRewardsAll,
		tags.Delegator, []byte(msg.DelegatorAddr.String()),
		tags.Reward, []byte(reward.ToString()),
	)
	resultTags = resultTags.AppendTags(withdrawTags)
	return sdk.Result{
		Tags: resultTags,
	}
}

func handleMsgWithdrawDelegatorReward(ctx sdk.Context, msg types.MsgWithdrawDelegatorReward, k keeper.Keeper) sdk.Result {

	reward, err := k.WithdrawDelegationReward(ctx, msg.DelegatorAddr, msg.ValidatorAddr)
	if err != nil {
		return err.Result()
	}

	tags := sdk.NewTags(
		tags.Action, tags.ActionWithdrawDelegatorReward,
		tags.Delegator, []byte(msg.DelegatorAddr.String()),
		tags.Validator, []byte(msg.ValidatorAddr.String()),
		tags.Reward, []byte(reward.ToString()),
	)
	return sdk.Result{
		Tags: tags,
	}
}

func handleMsgWithdrawValidatorRewardsAll(ctx sdk.Context, msg types.MsgWithdrawValidatorRewardsAll, k keeper.Keeper) sdk.Result {

	reward, withdrawTags, err := k.WithdrawValidatorRewardsAll(ctx, msg.ValidatorAddr)
	if err != nil {
		return err.Result()
	}

	resultTags := sdk.NewTags(
		tags.Action, tags.ActionWithdrawValidatorRewardsAll,
		tags.Validator, []byte(msg.ValidatorAddr.String()),
		tags.Reward, []byte(reward.ToString()),
	)
	resultTags = resultTags.AppendTags(withdrawTags)
	return sdk.Result{
		Tags: resultTags,
	}
}
