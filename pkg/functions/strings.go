// Copyright © 2020 Joseph Wright <joseph@cloudboss.co>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package functions

import (
	"encoding/base64"
	"fmt"

	"github.com/cloudboss/unobin/pkg/types"
	"github.com/hashicorp/go-multierror"
)

func Format(_ctx *types.Context, format String, args ...Interface) String {
	if format.Error != nil {
		return String{Error: format.Error}
	}

	var err error
	var argsValues = make([]interface{}, len(args))
	for i, arg := range args {
		argsValues[i] = arg.Value
		if arg.Error != nil {
			err = multierror.Append(err, arg.Error)
		}
	}
	if err != nil {
		return String{Error: err}
	}

	return String{fmt.Sprintf(format.Value, argsValues...), nil}
}

func B64Decode(_ctx *types.Context, input String) String {
	if input.Error != nil {
		return String{Error: input.Error}
	}

	out, err := base64.StdEncoding.DecodeString(input.Value)
	if err != nil {
		return String{Error: err}
	}
	return String{string(out), nil}
}

func B64Encode(_ctx *types.Context, input String) String {
	if input.Error != nil {
		return String{Error: input.Error}
	}

	out := base64.StdEncoding.EncodeToString([]byte(input.Value))
	return String{string(out), nil}
}
