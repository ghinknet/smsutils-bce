package bce

import (
	"github.com/baidubce/bce-sdk-go/services/sms"
	"github.com/ghinknet/smsutils/v3/errors"
	"github.com/ghinknet/smsutils/v3/model"
	"github.com/ghinknet/toolbox/expr"
)

type Client struct {
	Client *sms.Client
	// JSON
	Marshal   func(any) ([]byte, error)
	Unmarshal func([]byte, any) error
}

type Driver struct{}

func (d Driver) NewClient(params model.DriverClientParam) (model.Client, error) {
	// Check credential
	ak, sk := params.Credential[AccessKeyID], params.Credential[SecretAccessKey]
	if ak == "" || sk == "" {
		return Client{}, errors.ErrDriverCredentialInvalid.WithDriverName(Name)
	}

	// Create bce client
	client, err := sms.NewClient(
		ak, sk, expr.Ternary(params.Credential[Endpoint] != "", params.Credential[Endpoint], DefaultEndpoint),
	)
	if err != nil {
		return Client{}, nil
	}

	return Client{
		Client:    client,
		Marshal:   params.Marshal,
		Unmarshal: params.Unmarshal,
	}, nil
}
