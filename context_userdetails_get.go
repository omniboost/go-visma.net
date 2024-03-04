package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewContextUserdetailsGet() ContextUserdetailsGet {
	r := ContextUserdetailsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ContextUserdetailsGet struct {
	client      *Client
	queryParams *ContextUserdetailsGetQueryParams
	pathParams  *ContextUserdetailsGetPathParams
	method      string
	headers     http.Header
	requestBody ContextUserdetailsGetBody
}

func (r ContextUserdetailsGet) NewQueryParams() *ContextUserdetailsGetQueryParams {
	return &ContextUserdetailsGetQueryParams{}
}

type ContextUserdetailsGetQueryParams struct {
}

func (p ContextUserdetailsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ContextUserdetailsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r ContextUserdetailsGet) NewPathParams() *ContextUserdetailsGetPathParams {
	return &ContextUserdetailsGetPathParams{}
}

type ContextUserdetailsGetPathParams struct {
}

func (p *ContextUserdetailsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ContextUserdetailsGet) PathParams() *ContextUserdetailsGetPathParams {
	return r.pathParams
}

func (r *ContextUserdetailsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ContextUserdetailsGet) SetMethod(method string) {
	r.method = method
}

func (r *ContextUserdetailsGet) Method() string {
	return r.method
}

func (r ContextUserdetailsGet) NewRequestBody() ContextUserdetailsGetBody {
	return ContextUserdetailsGetBody{}
}

type ContextUserdetailsGetBody struct {
}

func (r *ContextUserdetailsGet) RequestBody() *ContextUserdetailsGetBody {
	return nil
}

func (r *ContextUserdetailsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ContextUserdetailsGet) SetRequestBody(body ContextUserdetailsGetBody) {
	r.requestBody = body
}

func (r *ContextUserdetailsGet) NewResponseBody() *ContextUserdetailsGetResponseBody {
	return &ContextUserdetailsGetResponseBody{}
}

type ContextUserdetailsGetResponseBody ContextUserdetails

func (r *ContextUserdetailsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("resources/v1/context/userdetails", r.PathParams())
	return &u
}

func (r *ContextUserdetailsGet) Do() (ContextUserdetailsGetResponseBody, error) {
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
