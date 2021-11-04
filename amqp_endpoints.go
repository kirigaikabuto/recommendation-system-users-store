package recommendation_system_users_store

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type UserAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewUserAmqpEndpoints(ch setdata_common.CommandHandler) UserAmqpEndpoints {
	return UserAmqpEndpoints{ch: ch}
}

func (u *UserAmqpEndpoints) MakeCreateUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (u *UserAmqpEndpoints) MakeUpdateUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &UpdateUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (u *UserAmqpEndpoints) MakeGetUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (u *UserAmqpEndpoints) MakeGetUserByUsernameAndPasswordAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetUserByUsernameAndPassword{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (u *UserAmqpEndpoints) MakeDeleteUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}

func (u *UserAmqpEndpoints) MakeListUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListUserCommand{}
		err := json.Unmarshal(message.Body, cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := u.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonData}
	}
}
