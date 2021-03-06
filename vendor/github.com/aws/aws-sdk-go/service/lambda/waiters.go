// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilFunctionExists uses the AWS Lambda API operation
// GetFunction to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Lambda) WaitUntilFunctionExists(input *GetFunctionInput) error {
	return c.WaitUntilFunctionExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFunctionExistsWithContext is an extended version of WaitUntilFunctionExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Lambda) WaitUntilFunctionExistsWithContext(ctx aws.Context, input *GetFunctionInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFunctionExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(1 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *GetFunctionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetFunctionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
