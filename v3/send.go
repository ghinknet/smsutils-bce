package bce

import (
	"github.com/baidubce/bce-sdk-go/services/sms/api"
	"go.gh.ink/smsutils/v3/errors"
	"go.gh.ink/smsutils/v3/model"
	"go.gh.ink/smsutils/v3/utils"
)

func (c Client) SendMessage(dest string, sender string, template string, vars model.Vars) error {
	// Try to parse number
	dest, _, _, _, err := utils.ProcessNumberForChinese(dest)
	if err != nil {
		return err
	}

	// Preprocess vars
	params := make(map[string]any)
	for _, v := range vars {
		params[v.Key] = v.Value
	}

	// Construct args
	sendSmsArgs := &api.SendSmsArgs{
		Mobile:      dest,
		Template:    template,
		SignatureId: sender,
		ContentVar:  params,
	}

	// Send request
	result, err := c.Client.SendSms(sendSmsArgs)
	if err != nil {
		return err
	}
	if result != nil {
		// Success
		if result.Code == "1000" {
			return nil
		}

		return errors.ErrDriverSendFailed.
			WithDriverName(Name).
			WithDriverCode(result.Code).
			WithDriverMessage(result.Message).
			WithDriverRequestID(result.RequestId).
			WithDriverResponse(result)
	}
	return errors.ErrDriverSendFailed.WithDriverName(Name)
}
