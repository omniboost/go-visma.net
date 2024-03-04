package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewContextGet() ContextGet {
	r := ContextGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ContextGet struct {
	client      *Client
	queryParams *ContextGetQueryParams
	pathParams  *ContextGetPathParams
	method      string
	headers     http.Header
	requestBody ContextGetBody
}

func (r ContextGet) NewQueryParams() *ContextGetQueryParams {
	return &ContextGetQueryParams{}
}

type ContextGetQueryParams struct {
}

func (p ContextGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ContextGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r ContextGet) NewPathParams() *ContextGetPathParams {
	return &ContextGetPathParams{}
}

type ContextGetPathParams struct {
}

func (p *ContextGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ContextGet) PathParams() *ContextGetPathParams {
	return r.pathParams
}

func (r *ContextGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ContextGet) SetMethod(method string) {
	r.method = method
}

func (r *ContextGet) Method() string {
	return r.method
}

func (r ContextGet) NewRequestBody() ContextGetBody {
	return ContextGetBody{}
}

type ContextGetBody struct {
}

func (r *ContextGet) RequestBody() *ContextGetBody {
	return nil
}

func (r *ContextGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ContextGet) SetRequestBody(body ContextGetBody) {
	r.requestBody = body
}

func (r *ContextGet) NewResponseBody() *ContextGetResponseBody {
	return &ContextGetResponseBody{}
}

type ContextGetResponseBody Context

func (r *ContextGet) URL() *url.URL {
	u := r.client.GetEndpointURL("resources/v1/context", r.PathParams())
	return &u
}

func (r *ContextGet) Do() (ContextGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
