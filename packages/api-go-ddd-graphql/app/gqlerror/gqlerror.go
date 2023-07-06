package gqlerror

import (
	"context"
	"errors"

	"golang.org/x/text/message"
)

// Error GraphQLのエラー
type Error struct {
	s          string
	extensions map[string]interface{}
}

// Handler GraphQLのエラーハンドラ
type Handler struct {
	Presenters map[error]Presenter
}

func NewHandler(pre map[error]Presenter) (Handler, error) {
	for err, p := range pre {
		for tag, msg := range p.Lang2Msg {
			if err := message.SetString(tag, err.Error(), msg); err != nil {
				return Handler{}, err
			}
		}
	}

	return Handler{
		Presenters: pre,
	}, nil
}

// New エラーを生成する
func (h Handler) New(ctx context.Context, err error) Error {
	for perr, pre := range h.Presenters {
		if !errors.Is(err, perr) {
			continue
		}

		v := Error{
			s: err.Error(),
			extensions: map[string]interface{}{
				"code": pre.Code,
			},
		}

		tag := TagFromContext(ctx)
		msg := message.NewPrinter(tag).Sprintf(err.Error())

		v.extensions["message"] = msg
		return v
	}

	return Error{s: err.Error()}
}

// Error エラーを文字列に変換する
func (e Error) Error() string {
	return e.s
}

// Extensions エラーの拡張情報を取得する
func (e Error) Extensions() map[string]interface{} {
	return e.extensions
}
