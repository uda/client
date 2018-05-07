// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/stellar1/remote.avdl

package stellar1

import (
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PaymentDirectPost struct {
	FromDeviceID      keybase1.DeviceID     `codec:"fromDeviceID" json:"fromDeviceID"`
	To                *keybase1.UserVersion `codec:"to,omitempty" json:"to,omitempty"`
	DisplayAmount     string                `codec:"displayAmount" json:"displayAmount"`
	DisplayCurrency   string                `codec:"displayCurrency" json:"displayCurrency"`
	NoteB64           string                `codec:"noteB64" json:"noteB64"`
	SignedTransaction string                `codec:"signedTransaction" json:"signedTransaction"`
}

func (o PaymentDirectPost) DeepCopy() PaymentDirectPost {
	return PaymentDirectPost{
		FromDeviceID: o.FromDeviceID.DeepCopy(),
		To: (func(x *keybase1.UserVersion) *keybase1.UserVersion {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.To),
		DisplayAmount:     o.DisplayAmount,
		DisplayCurrency:   o.DisplayCurrency,
		NoteB64:           o.NoteB64,
		SignedTransaction: o.SignedTransaction,
	}
}

type PaymentRelayPost struct {
	FromDeviceID      keybase1.DeviceID     `codec:"fromDeviceID" json:"fromDeviceID"`
	To                *keybase1.UserVersion `codec:"to,omitempty" json:"to,omitempty"`
	ToAssertion       string                `codec:"toAssertion" json:"toAssertion"`
	RelayAccount      AccountID             `codec:"relayAccount" json:"relayAccount"`
	TeamID            keybase1.TeamID       `codec:"teamID" json:"teamID"`
	DisplayAmount     string                `codec:"displayAmount" json:"displayAmount"`
	DisplayCurrency   string                `codec:"displayCurrency" json:"displayCurrency"`
	BoxB64            string                `codec:"boxB64" json:"boxB64"`
	SignedTransaction string                `codec:"signedTransaction" json:"signedTransaction"`
}

func (o PaymentRelayPost) DeepCopy() PaymentRelayPost {
	return PaymentRelayPost{
		FromDeviceID: o.FromDeviceID.DeepCopy(),
		To: (func(x *keybase1.UserVersion) *keybase1.UserVersion {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.To),
		ToAssertion:       o.ToAssertion,
		RelayAccount:      o.RelayAccount.DeepCopy(),
		TeamID:            o.TeamID.DeepCopy(),
		DisplayAmount:     o.DisplayAmount,
		DisplayCurrency:   o.DisplayCurrency,
		BoxB64:            o.BoxB64,
		SignedTransaction: o.SignedTransaction,
	}
}

type PaymentSummary struct {
	Stellar     *PaymentSummaryStellar `codec:"stellar,omitempty" json:"stellar,omitempty"`
	Keybase     *PaymentSummaryKeybase `codec:"keybase,omitempty" json:"keybase,omitempty"`
	StellarTxID TransactionID          `codec:"stellarTxID" json:"stellarTxID"`
	From        AccountID              `codec:"from" json:"from"`
	To          AccountID              `codec:"to" json:"to"`
	Amount      string                 `codec:"amount" json:"amount"`
	Asset       Asset                  `codec:"asset" json:"asset"`
}

func (o PaymentSummary) DeepCopy() PaymentSummary {
	return PaymentSummary{
		Stellar: (func(x *PaymentSummaryStellar) *PaymentSummaryStellar {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Stellar),
		Keybase: (func(x *PaymentSummaryKeybase) *PaymentSummaryKeybase {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Keybase),
		StellarTxID: o.StellarTxID.DeepCopy(),
		From:        o.From.DeepCopy(),
		To:          o.To.DeepCopy(),
		Amount:      o.Amount,
		Asset:       o.Asset.DeepCopy(),
	}
}

type PaymentSummaryKeybase struct {
	KbTxID          KeybaseTransactionID  `codec:"kbTxID" json:"kbTxID"`
	Status          TransactionStatus     `codec:"status" json:"status"`
	SubmitErrMsg    string                `codec:"submitErrMsg" json:"submitErrMsg"`
	Ctime           TimeMs                `codec:"ctime" json:"ctime"`
	Rtime           TimeMs                `codec:"rtime" json:"rtime"`
	From            keybase1.UserVersion  `codec:"from" json:"from"`
	FromDeviceID    keybase1.DeviceID     `codec:"fromDeviceID" json:"fromDeviceID"`
	To              *keybase1.UserVersion `codec:"to,omitempty" json:"to,omitempty"`
	DisplayAmount   *string               `codec:"displayAmount,omitempty" json:"displayAmount,omitempty"`
	DisplayCurrency *string               `codec:"displayCurrency,omitempty" json:"displayCurrency,omitempty"`
	NoteB64         string                `codec:"noteB64" json:"noteB64"`
}

func (o PaymentSummaryKeybase) DeepCopy() PaymentSummaryKeybase {
	return PaymentSummaryKeybase{
		KbTxID:       o.KbTxID.DeepCopy(),
		Status:       o.Status.DeepCopy(),
		SubmitErrMsg: o.SubmitErrMsg,
		Ctime:        o.Ctime.DeepCopy(),
		Rtime:        o.Rtime.DeepCopy(),
		From:         o.From.DeepCopy(),
		FromDeviceID: o.FromDeviceID.DeepCopy(),
		To: (func(x *keybase1.UserVersion) *keybase1.UserVersion {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.To),
		DisplayAmount: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayAmount),
		DisplayCurrency: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.DisplayCurrency),
		NoteB64: o.NoteB64,
	}
}

type PaymentSummaryStellar struct {
	OperationID uint64 `codec:"operationID" json:"operationID"`
	Ctime       TimeMs `codec:"ctime" json:"ctime"`
}

func (o PaymentSummaryStellar) DeepCopy() PaymentSummaryStellar {
	return PaymentSummaryStellar{
		OperationID: o.OperationID,
		Ctime:       o.Ctime.DeepCopy(),
	}
}

type BalancesArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
}

type RecentPaymentsArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
	Limit     int                  `codec:"limit" json:"limit"`
}

type PaymentDetailArg struct {
	Caller keybase1.UserVersion `codec:"caller" json:"caller"`
	TxID   string               `codec:"txID" json:"txID"`
}

type AccountSeqnoArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
}

type SubmitPaymentArg struct {
	Caller  keybase1.UserVersion `codec:"caller" json:"caller"`
	Payment PaymentDirectPost    `codec:"payment" json:"payment"`
}

type SubmitRelayPaymentArg struct {
	Caller  keybase1.UserVersion `codec:"caller" json:"caller"`
	Payment PaymentRelayPost     `codec:"payment" json:"payment"`
}

type IsMasterKeyActiveArg struct {
	Caller    keybase1.UserVersion `codec:"caller" json:"caller"`
	AccountID AccountID            `codec:"accountID" json:"accountID"`
}

type PingArg struct {
}

type RemoteInterface interface {
	Balances(context.Context, BalancesArg) ([]Balance, error)
	RecentPayments(context.Context, RecentPaymentsArg) ([]PaymentSummary, error)
	PaymentDetail(context.Context, PaymentDetailArg) (PaymentSummary, error)
	AccountSeqno(context.Context, AccountSeqnoArg) (string, error)
	SubmitPayment(context.Context, SubmitPaymentArg) (PaymentResult, error)
	SubmitRelayPayment(context.Context, SubmitRelayPaymentArg) (PaymentResult, error)
	IsMasterKeyActive(context.Context, IsMasterKeyActiveArg) (bool, error)
	Ping(context.Context) (string, error)
}

func RemoteProtocol(i RemoteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "stellar.1.remote",
		Methods: map[string]rpc.ServeHandlerDescription{
			"balances": {
				MakeArg: func() interface{} {
					ret := make([]BalancesArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]BalancesArg)
					if !ok {
						err = rpc.NewTypeError((*[]BalancesArg)(nil), args)
						return
					}
					ret, err = i.Balances(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"recentPayments": {
				MakeArg: func() interface{} {
					ret := make([]RecentPaymentsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RecentPaymentsArg)
					if !ok {
						err = rpc.NewTypeError((*[]RecentPaymentsArg)(nil), args)
						return
					}
					ret, err = i.RecentPayments(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"paymentDetail": {
				MakeArg: func() interface{} {
					ret := make([]PaymentDetailArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PaymentDetailArg)
					if !ok {
						err = rpc.NewTypeError((*[]PaymentDetailArg)(nil), args)
						return
					}
					ret, err = i.PaymentDetail(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"accountSeqno": {
				MakeArg: func() interface{} {
					ret := make([]AccountSeqnoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AccountSeqnoArg)
					if !ok {
						err = rpc.NewTypeError((*[]AccountSeqnoArg)(nil), args)
						return
					}
					ret, err = i.AccountSeqno(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"submitPayment": {
				MakeArg: func() interface{} {
					ret := make([]SubmitPaymentArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SubmitPaymentArg)
					if !ok {
						err = rpc.NewTypeError((*[]SubmitPaymentArg)(nil), args)
						return
					}
					ret, err = i.SubmitPayment(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"submitRelayPayment": {
				MakeArg: func() interface{} {
					ret := make([]SubmitRelayPaymentArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SubmitRelayPaymentArg)
					if !ok {
						err = rpc.NewTypeError((*[]SubmitRelayPaymentArg)(nil), args)
						return
					}
					ret, err = i.SubmitRelayPayment(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"isMasterKeyActive": {
				MakeArg: func() interface{} {
					ret := make([]IsMasterKeyActiveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]IsMasterKeyActiveArg)
					if !ok {
						err = rpc.NewTypeError((*[]IsMasterKeyActiveArg)(nil), args)
						return
					}
					ret, err = i.IsMasterKeyActive(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"ping": {
				MakeArg: func() interface{} {
					ret := make([]PingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.Ping(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RemoteClient struct {
	Cli rpc.GenericClient
}

func (c RemoteClient) Balances(ctx context.Context, __arg BalancesArg) (res []Balance, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.balances", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) RecentPayments(ctx context.Context, __arg RecentPaymentsArg) (res []PaymentSummary, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.recentPayments", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) PaymentDetail(ctx context.Context, __arg PaymentDetailArg) (res PaymentSummary, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.paymentDetail", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) AccountSeqno(ctx context.Context, __arg AccountSeqnoArg) (res string, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.accountSeqno", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) SubmitPayment(ctx context.Context, __arg SubmitPaymentArg) (res PaymentResult, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.submitPayment", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) SubmitRelayPayment(ctx context.Context, __arg SubmitRelayPaymentArg) (res PaymentResult, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.submitRelayPayment", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) IsMasterKeyActive(ctx context.Context, __arg IsMasterKeyActiveArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.isMasterKeyActive", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) Ping(ctx context.Context) (res string, err error) {
	err = c.Cli.Call(ctx, "stellar.1.remote.ping", []interface{}{PingArg{}}, &res)
	return
}
