// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/76e25d34bff1060e300c95f4be468ef88e4f3465

// Creates part of a trained model definition
package puttrainedmodeldefinitionpart

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	modelidMask = iota + 1

	partMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTrainedModelDefinitionPart struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	modelid string
	part    string
}

// NewPutTrainedModelDefinitionPart type alias for index.
type NewPutTrainedModelDefinitionPart func(modelid, part string) *PutTrainedModelDefinitionPart

// NewPutTrainedModelDefinitionPartFunc returns a new instance of PutTrainedModelDefinitionPart with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutTrainedModelDefinitionPartFunc(tp elastictransport.Interface) NewPutTrainedModelDefinitionPart {
	return func(modelid, part string) *PutTrainedModelDefinitionPart {
		n := New(tp)

		n.ModelId(modelid)

		n.Part(part)

		return n
	}
}

// Creates part of a trained model definition
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-definition-part.html
func New(tp elastictransport.Interface) *PutTrainedModelDefinitionPart {
	r := &PutTrainedModelDefinitionPart{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutTrainedModelDefinitionPart) Raw(raw io.Reader) *PutTrainedModelDefinitionPart {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutTrainedModelDefinitionPart) Request(req *Request) *PutTrainedModelDefinitionPart {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutTrainedModelDefinitionPart) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutTrainedModelDefinitionPart: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask|partMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("definition")
		path.WriteString("/")

		path.WriteString(r.part)

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutTrainedModelDefinitionPart) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutTrainedModelDefinitionPart query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttrainedmodeldefinitionpart.Response
func (r PutTrainedModelDefinitionPart) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the PutTrainedModelDefinitionPart headers map.
func (r *PutTrainedModelDefinitionPart) Header(key, value string) *PutTrainedModelDefinitionPart {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model.
// API Name: modelid
func (r *PutTrainedModelDefinitionPart) ModelId(modelid string) *PutTrainedModelDefinitionPart {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// Part The definition part number. When the definition is loaded for inference the
// definition parts are streamed in the
// order of their part number. The first part must be `0` and the final part
// must be `total_parts - 1`.
// API Name: part
func (r *PutTrainedModelDefinitionPart) Part(part string) *PutTrainedModelDefinitionPart {
	r.paramSet |= partMask
	r.part = part

	return r
}

// Definition The definition part for the model. Must be a base64 encoded string.
// API name: definition
func (r *PutTrainedModelDefinitionPart) Definition(definition string) *PutTrainedModelDefinitionPart {

	r.req.Definition = definition

	return r
}

// TotalDefinitionLength The total uncompressed definition length in bytes. Not base64 encoded.
// API name: total_definition_length
func (r *PutTrainedModelDefinitionPart) TotalDefinitionLength(totaldefinitionlength int64) *PutTrainedModelDefinitionPart {

	r.req.TotalDefinitionLength = totaldefinitionlength

	return r
}

// TotalParts The total number of parts that will be uploaded. Must be greater than 0.
// API name: total_parts
func (r *PutTrainedModelDefinitionPart) TotalParts(totalparts int) *PutTrainedModelDefinitionPart {
	r.req.TotalParts = totalparts

	return r
}
