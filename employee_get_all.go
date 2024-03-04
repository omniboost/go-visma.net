package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewEmployeeGetAll() EmployeeGetAll {
	r := EmployeeGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeeGetAll struct {
	client      *Client
	queryParams *EmployeeGetAllQueryParams
	pathParams  *EmployeeGetAllPathParams
	method      string
	headers     http.Header
	requestBody EmployeeGetAllBody
}

func (r EmployeeGetAll) NewQueryParams() *EmployeeGetAllQueryParams {
	return &EmployeeGetAllQueryParams{}
}

type EmployeeGetAllQueryParams struct {
	Name              string `schema:"name,omitempty"`
	VATRegistrationID string `schema:"vatRegistrationId,omitempty"`
	Email             string `schema:"email,omitempty"`
	Phone             string `schema:"phone,omitempty"`
	Status            string `schema:"status,omitempty"`
}

func (p EmployeeGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeeGetAll) QueryParams() *EmployeeGetAllQueryParams {
	return r.queryParams
}

func (r EmployeeGetAll) NewPathParams() *EmployeeGetAllPathParams {
	return &EmployeeGetAllPathParams{}
}

type EmployeeGetAllPathParams struct {
}

func (p *EmployeeGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeeGetAll) PathParams() *EmployeeGetAllPathParams {
	return r.pathParams
}

func (r *EmployeeGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeeGetAll) SetMethod(method string) {
	r.method = method
}

func (r *EmployeeGetAll) Method() string {
	return r.method
}

func (r EmployeeGetAll) NewRequestBody() EmployeeGetAllBody {
	return EmployeeGetAllBody{}
}

type EmployeeGetAllBody struct {
}

func (r *EmployeeGetAll) RequestBody() *EmployeeGetAllBody {
	return nil
}

func (r *EmployeeGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *EmployeeGetAll) SetRequestBody(body EmployeeGetAllBody) {
	r.requestBody = body
}

func (r *EmployeeGetAll) NewResponseBody() *EmployeeGetAllResponseBody {
	return &EmployeeGetAllResponseBody{}
}

type EmployeeGetAllResponseBody Employees

func (r *EmployeeGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/employee", r.PathParams())
	return &u
}

func (r *EmployeeGetAll) Do() (EmployeeGetAllResponseBody, error) {
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
