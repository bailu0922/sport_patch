package sport_patch

import "context"

type Dealer interface {
	Deal(string) func(context.Context, any, func(context.Context, string, any)) error
	Clean() func(context.Context, string, string, func(context.Context, string, any)) error
	Check() func(context.Context, any, func(context.Context, string, any)) (bool, error)
}

type Helper interface {
	CreatedHandler(context.Context, any, func(context.Context, string, any)) error
	RunningHandler(context.Context, any, func(context.Context, string, any)) error
	StoppedHandler(context.Context, any, func(context.Context, string, any)) error
	CompletedHandler(context.Context, any, func(context.Context, string, any)) error
	ReadyHandler(context.Context, any, func(context.Context, string, any)) error
	ViolationHandler(context.Context, any, func(context.Context, string, any)) (bool, error)
	Teardown(context.Context, string, string, func(context.Context, string, any)) error
}

type Interface interface {
	GetHelper(string) Dealer
}
