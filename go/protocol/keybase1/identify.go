// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/identify.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type IdentifyProofBreak struct {
	RemoteProof RemoteProof     `codec:"remoteProof" json:"remoteProof"`
	Lcr         LinkCheckResult `codec:"lcr" json:"lcr"`
}

func (o IdentifyProofBreak) DeepCopy() IdentifyProofBreak {
	return IdentifyProofBreak{
		RemoteProof: o.RemoteProof.DeepCopy(),
		Lcr:         o.Lcr.DeepCopy(),
	}
}

type IdentifyTrackBreaks struct {
	Keys   []IdentifyKey        `codec:"keys" json:"keys"`
	Proofs []IdentifyProofBreak `codec:"proofs" json:"proofs"`
}

func (o IdentifyTrackBreaks) DeepCopy() IdentifyTrackBreaks {
	return IdentifyTrackBreaks{
		Keys: (func(x []IdentifyKey) []IdentifyKey {
			var ret []IdentifyKey
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Keys),
		Proofs: (func(x []IdentifyProofBreak) []IdentifyProofBreak {
			var ret []IdentifyProofBreak
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Proofs),
	}
}

type Identify2Res struct {
	Upk         UserPlusKeys         `codec:"upk" json:"upk"`
	TrackBreaks *IdentifyTrackBreaks `codec:"trackBreaks,omitempty" json:"trackBreaks,omitempty"`
}

func (o Identify2Res) DeepCopy() Identify2Res {
	return Identify2Res{
		Upk: o.Upk.DeepCopy(),
		TrackBreaks: (func(x *IdentifyTrackBreaks) *IdentifyTrackBreaks {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TrackBreaks),
	}
}

type IdentifyLiteRes struct {
	Ul          UserOrTeamLite       `codec:"ul" json:"ul"`
	TrackBreaks *IdentifyTrackBreaks `codec:"trackBreaks,omitempty" json:"trackBreaks,omitempty"`
}

func (o IdentifyLiteRes) DeepCopy() IdentifyLiteRes {
	return IdentifyLiteRes{
		Ul: o.Ul.DeepCopy(),
		TrackBreaks: (func(x *IdentifyTrackBreaks) *IdentifyTrackBreaks {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TrackBreaks),
	}
}

type ResolveArg struct {
	Assertion string `codec:"assertion" json:"assertion"`
}

func (o ResolveArg) DeepCopy() ResolveArg {
	return ResolveArg{
		Assertion: o.Assertion,
	}
}

type Resolve2Arg struct {
	Assertion string `codec:"assertion" json:"assertion"`
}

func (o Resolve2Arg) DeepCopy() Resolve2Arg {
	return Resolve2Arg{
		Assertion: o.Assertion,
	}
}

type IdentifyArg struct {
	SessionID        int            `codec:"sessionID" json:"sessionID"`
	UserAssertion    string         `codec:"userAssertion" json:"userAssertion"`
	ForceRemoteCheck bool           `codec:"forceRemoteCheck" json:"forceRemoteCheck"`
	UseDelegateUI    bool           `codec:"useDelegateUI" json:"useDelegateUI"`
	Reason           IdentifyReason `codec:"reason" json:"reason"`
	Source           ClientType     `codec:"source" json:"source"`
}

func (o IdentifyArg) DeepCopy() IdentifyArg {
	return IdentifyArg{
		SessionID:        o.SessionID,
		UserAssertion:    o.UserAssertion,
		ForceRemoteCheck: o.ForceRemoteCheck,
		UseDelegateUI:    o.UseDelegateUI,
		Reason:           o.Reason.DeepCopy(),
		Source:           o.Source.DeepCopy(),
	}
}

type Identify2Arg struct {
	SessionID             int                 `codec:"sessionID" json:"sessionID"`
	Uid                   UID                 `codec:"uid" json:"uid"`
	UserAssertion         string              `codec:"userAssertion" json:"userAssertion"`
	Reason                IdentifyReason      `codec:"reason" json:"reason"`
	UseDelegateUI         bool                `codec:"useDelegateUI" json:"useDelegateUI"`
	AlwaysBlock           bool                `codec:"alwaysBlock" json:"alwaysBlock"`
	NoErrorOnTrackFailure bool                `codec:"noErrorOnTrackFailure" json:"noErrorOnTrackFailure"`
	ForceRemoteCheck      bool                `codec:"forceRemoteCheck" json:"forceRemoteCheck"`
	NeedProofSet          bool                `codec:"needProofSet" json:"needProofSet"`
	AllowEmptySelfID      bool                `codec:"allowEmptySelfID" json:"allowEmptySelfID"`
	NoSkipSelf            bool                `codec:"noSkipSelf" json:"noSkipSelf"`
	CanSuppressUI         bool                `codec:"canSuppressUI" json:"canSuppressUI"`
	IdentifyBehavior      TLFIdentifyBehavior `codec:"identifyBehavior" json:"identifyBehavior"`
	ForceDisplay          bool                `codec:"forceDisplay" json:"forceDisplay"`
}

func (o Identify2Arg) DeepCopy() Identify2Arg {
	return Identify2Arg{
		SessionID:             o.SessionID,
		Uid:                   o.Uid.DeepCopy(),
		UserAssertion:         o.UserAssertion,
		Reason:                o.Reason.DeepCopy(),
		UseDelegateUI:         o.UseDelegateUI,
		AlwaysBlock:           o.AlwaysBlock,
		NoErrorOnTrackFailure: o.NoErrorOnTrackFailure,
		ForceRemoteCheck:      o.ForceRemoteCheck,
		NeedProofSet:          o.NeedProofSet,
		AllowEmptySelfID:      o.AllowEmptySelfID,
		NoSkipSelf:            o.NoSkipSelf,
		CanSuppressUI:         o.CanSuppressUI,
		IdentifyBehavior:      o.IdentifyBehavior.DeepCopy(),
		ForceDisplay:          o.ForceDisplay,
	}
}

type IdentifyLiteArg struct {
	SessionID             int                 `codec:"sessionID" json:"sessionID"`
	Id                    UserOrTeamID        `codec:"id" json:"id"`
	Assertion             string              `codec:"assertion" json:"assertion"`
	Reason                IdentifyReason      `codec:"reason" json:"reason"`
	UseDelegateUI         bool                `codec:"useDelegateUI" json:"useDelegateUI"`
	AlwaysBlock           bool                `codec:"alwaysBlock" json:"alwaysBlock"`
	NoErrorOnTrackFailure bool                `codec:"noErrorOnTrackFailure" json:"noErrorOnTrackFailure"`
	ForceRemoteCheck      bool                `codec:"forceRemoteCheck" json:"forceRemoteCheck"`
	NeedProofSet          bool                `codec:"needProofSet" json:"needProofSet"`
	AllowEmptySelfID      bool                `codec:"allowEmptySelfID" json:"allowEmptySelfID"`
	NoSkipSelf            bool                `codec:"noSkipSelf" json:"noSkipSelf"`
	CanSuppressUI         bool                `codec:"canSuppressUI" json:"canSuppressUI"`
	IdentifyBehavior      TLFIdentifyBehavior `codec:"identifyBehavior" json:"identifyBehavior"`
	ForceDisplay          bool                `codec:"forceDisplay" json:"forceDisplay"`
}

func (o IdentifyLiteArg) DeepCopy() IdentifyLiteArg {
	return IdentifyLiteArg{
		SessionID:             o.SessionID,
		Id:                    o.Id.DeepCopy(),
		Assertion:             o.Assertion,
		Reason:                o.Reason.DeepCopy(),
		UseDelegateUI:         o.UseDelegateUI,
		AlwaysBlock:           o.AlwaysBlock,
		NoErrorOnTrackFailure: o.NoErrorOnTrackFailure,
		ForceRemoteCheck:      o.ForceRemoteCheck,
		NeedProofSet:          o.NeedProofSet,
		AllowEmptySelfID:      o.AllowEmptySelfID,
		NoSkipSelf:            o.NoSkipSelf,
		CanSuppressUI:         o.CanSuppressUI,
		IdentifyBehavior:      o.IdentifyBehavior.DeepCopy(),
		ForceDisplay:          o.ForceDisplay,
	}
}

type IdentifyInterface interface {
	// Resolve an assertion to a UID. On failure, resolves to an empty UID and returns
	// an error.
	Resolve(context.Context, string) (UID, error)
	// Resolve an assertion to a (UID,username) or (TeamID,teamname). On failure, returns an error.
	Resolve2(context.Context, string) (UserOrTeamLite, error)
	// DEPRECATED:  use identify2
	//
	// Identify a user from a username or assertion (e.g. kbuser, twuser@twitter).
	// If forceRemoteCheck is true, we force all remote proofs to be checked (otherwise a cache is used).
	Identify(context.Context, IdentifyArg) (IdentifyRes, error)
	Identify2(context.Context, Identify2Arg) (Identify2Res, error)
	IdentifyLite(context.Context, IdentifyLiteArg) (IdentifyLiteRes, error)
}

func IdentifyProtocol(i IdentifyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.identify",
		Methods: map[string]rpc.ServeHandlerDescription{
			"Resolve": {
				MakeArg: func() interface{} {
					ret := make([]ResolveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ResolveArg)
					if !ok {
						err = rpc.NewTypeError((*[]ResolveArg)(nil), args)
						return
					}
					ret, err = i.Resolve(ctx, (*typedArgs)[0].Assertion)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"Resolve2": {
				MakeArg: func() interface{} {
					ret := make([]Resolve2Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]Resolve2Arg)
					if !ok {
						err = rpc.NewTypeError((*[]Resolve2Arg)(nil), args)
						return
					}
					ret, err = i.Resolve2(ctx, (*typedArgs)[0].Assertion)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"identify": {
				MakeArg: func() interface{} {
					ret := make([]IdentifyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]IdentifyArg)
					if !ok {
						err = rpc.NewTypeError((*[]IdentifyArg)(nil), args)
						return
					}
					ret, err = i.Identify(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"identify2": {
				MakeArg: func() interface{} {
					ret := make([]Identify2Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]Identify2Arg)
					if !ok {
						err = rpc.NewTypeError((*[]Identify2Arg)(nil), args)
						return
					}
					ret, err = i.Identify2(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"identifyLite": {
				MakeArg: func() interface{} {
					ret := make([]IdentifyLiteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]IdentifyLiteArg)
					if !ok {
						err = rpc.NewTypeError((*[]IdentifyLiteArg)(nil), args)
						return
					}
					ret, err = i.IdentifyLite(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type IdentifyClient struct {
	Cli rpc.GenericClient
}

// Resolve an assertion to a UID. On failure, resolves to an empty UID and returns
// an error.
func (c IdentifyClient) Resolve(ctx context.Context, assertion string) (res UID, err error) {
	__arg := ResolveArg{Assertion: assertion}
	err = c.Cli.Call(ctx, "keybase.1.identify.Resolve", []interface{}{__arg}, &res)
	return
}

// Resolve an assertion to a (UID,username) or (TeamID,teamname). On failure, returns an error.
func (c IdentifyClient) Resolve2(ctx context.Context, assertion string) (res UserOrTeamLite, err error) {
	__arg := Resolve2Arg{Assertion: assertion}
	err = c.Cli.Call(ctx, "keybase.1.identify.Resolve2", []interface{}{__arg}, &res)
	return
}

// DEPRECATED:  use identify2
//
// Identify a user from a username or assertion (e.g. kbuser, twuser@twitter).
// If forceRemoteCheck is true, we force all remote proofs to be checked (otherwise a cache is used).
func (c IdentifyClient) Identify(ctx context.Context, __arg IdentifyArg) (res IdentifyRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identify.identify", []interface{}{__arg}, &res)
	return
}

func (c IdentifyClient) Identify2(ctx context.Context, __arg Identify2Arg) (res Identify2Res, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identify.identify2", []interface{}{__arg}, &res)
	return
}

func (c IdentifyClient) IdentifyLite(ctx context.Context, __arg IdentifyLiteArg) (res IdentifyLiteRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.identify.identifyLite", []interface{}{__arg}, &res)
	return
}
