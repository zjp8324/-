package model

import (
	"errors"
	"regexp"

	json "github.com/json-iterator/go"

	"github.com/gin-gonic/gin"
)

var (
	ErrEmptyRoomId          = errors.New("empty room id")
	ErrRoomIdTooLong        = errors.New("room id too long")
	ErrRoomIdHasInvalidChar = errors.New("room id has invalid char")

	ErrEmptyPassword   = errors.New("empty password")
	ErrPasswordTooLong = errors.New("password too long")

	ErrEmptyUsername          = errors.New("empty username")
	ErrUsernameTooLong        = errors.New("username too long")
	ErrUsernameHasInvalidChar = errors.New("username has invalid char")
)

var (
	alphaNumReg = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

type CreateRoomReq struct {
	RoomId       string `json:"roomId"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	UserPassword string `json:"userPassword"`
	Hidden       bool   `json:"hidden"`
}

func (c *CreateRoomReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(c)
}

func (c *CreateRoomReq) Validate() error {
	if c.RoomId == "" {
		return ErrEmptyRoomId
	} else if len(c.RoomId) > 32 {
		return ErrRoomIdTooLong
	} else if !alphaNumReg.MatchString(c.RoomId) {
		return ErrRoomIdHasInvalidChar
	}

	if len(c.Password) > 32 {
		return ErrPasswordTooLong
	}

	if c.Username == "" {
		return ErrEmptyUsername
	} else if len(c.Username) > 32 {
		return ErrUsernameTooLong
	} else if !alphaNumReg.MatchString(c.Username) {
		return ErrUsernameHasInvalidChar
	}

	if c.UserPassword == "" {
		return ErrEmptyPassword
	} else if len(c.UserPassword) > 32 {
		return ErrPasswordTooLong
	}

	return nil
}

type RoomListResp struct {
	RoomId       string `json:"roomId"`
	PeopleNum    int64  `json:"peopleNum"`
	NeedPassword bool   `json:"needPassword"`
	Creator      string `json:"creator"`
	CreateAt     int64  `json:"createAt"`
}

type LoginRoomReq struct {
	RoomId       string `json:"roomId"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	UserPassword string `json:"userPassword"`
}

func (l *LoginRoomReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(l)
}

func (l *LoginRoomReq) Validate() error {
	if l.RoomId == "" {
		return ErrEmptyRoomId
	}

	if l.Username == "" {
		return ErrEmptyUsername
	}

	if l.UserPassword == "" {
		return ErrEmptyPassword
	}

	return nil
}

type SetRoomPasswordReq struct {
	Password string `json:"password"`
}

func (s *SetRoomPasswordReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(s)
}

func (s *SetRoomPasswordReq) Validate() error {
	if len(s.Password) > 32 {
		return ErrPasswordTooLong
	}
	return nil
}

type UsernameReq struct {
	Username string `json:"username"`
}

func (u *UsernameReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(u)
}

func (u *UsernameReq) Validate() error {
	if u.Username == "" {
		return ErrEmptyUsername
	} else if len(u.Username) > 32 {
		return ErrUsernameTooLong
	} else if !alphaNumReg.MatchString(u.Username) {
		return ErrUsernameHasInvalidChar
	}
	return nil
}