package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewDepartmentGetAll() DepartmentGetAll {
	r := DepartmentGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DepartmentGetAll struct {
	client      *Client
	queryParams *DepartmentGetAllQueryParams
	pathParams  *DepartmentGetAllPathParams
	method      string
	headers     http.Header
	requestBody DepartmentGetAllBody
}

func (r DepartmentGetAll) NewQueryParams() *DepartmentGetAllQueryParams {
	return &DepartmentGetAllQueryParams{}
}

type DepartmentGetAllQueryParams struct {
	Name              string `schema:"name,omitempty"`
	VATRegistrationID string `schema:"vatRegistrationId,omitempty"`
	Email             string `schema:"email,omitempty"`
	Phone             string `schema:"phone,omitempty"`
	Status            string `schema:"status,omitempty"`
}

func (p DepartmentGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DepartmentGetAll) QueryParams() *DepartmentGetAllQueryParams {
	return r.queryParams
}

func (r DepartmentGetAll) NewPathParams() *DepartmentGetAllPathParams {
	return &DepartmentGetAllPathParams{}
}

type DepartmentGetAllPathParams struct {
}

func (p *DepartmentGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DepartmentGetAll) PathParams() *DepartmentGetAllPathParams {
	return r.pathParams
}

func (r *DepartmentGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DepartmentGetAll) SetMethod(method string) {
	r.method = method
}

func (r *DepartmentGetAll) Method() string {
	return r.method
}

func (r DepartmentGetAll) NewRequestBody() DepartmentGetAllBody {
	return DepartmentGetAllBody{}
}

type DepartmentGetAllBody struct {
}

func (r *DepartmentGetAll) RequestBody() *DepartmentGetAllBody {
	return nil
}

func (r *DepartmentGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *DepartmentGetAll) SetRequestBody(body DepartmentGetAllBody) {
	r.requestBody = body
}

func (r *DepartmentGetAll) NewResponseBody() *DepartmentGetAllResponseBody {
	return &DepartmentGetAllResponseBody{}
}

type DepartmentGetAllResponseBody Departments

func (r *DepartmentGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/department", r.PathParams())
	return &u
}

func (r *DepartmentGetAll) Do() (DepartmentGetAllResponseBody, error) {
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
