/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package scale

import (
	"github.com/loopholelabs/scale/go/context"
	"net"
)

func Scale(ctx *context.Context) *context.Context {
	_, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		ctx.Response().SetBody("error opening connection")
	} else {
		ctx.Response().SetBody("success")
	}
	return ctx
}