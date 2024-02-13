// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.temp.checkSignupQueue

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// TempCheckSignupQueue_Output is the output of a com.atproto.temp.checkSignupQueue call.
type TempCheckSignupQueue_Output struct {
	Activated       bool   `json:"activated" cborgen:"activated"`
	EstimatedTimeMs *int64 `json:"estimatedTimeMs,omitempty" cborgen:"estimatedTimeMs,omitempty"`
	PlaceInQueue    *int64 `json:"placeInQueue,omitempty" cborgen:"placeInQueue,omitempty"`
}

// TempCheckSignupQueue calls the XRPC method "com.atproto.temp.checkSignupQueue".
func TempCheckSignupQueue(ctx context.Context, c *xrpc.Client) (*TempCheckSignupQueue_Output, error) {
	var out TempCheckSignupQueue_Output
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.temp.checkSignupQueue", nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}