/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// PrincipalApiService PrincipalApi service
type PrincipalApiService service

type ApiListPrincipalsRequest struct {
	ctx _context.Context
	ApiService *PrincipalApiService
	limit *int32
	offset *int32
	matchCriteria *string
	usernames *string
	sortOrder *string
	email *string
	status *string
	adminOnly *string
	orderBy *string
}

func (r ApiListPrincipalsRequest) Limit(limit int32) ApiListPrincipalsRequest {
	r.limit = &limit
	return r
}
func (r ApiListPrincipalsRequest) Offset(offset int32) ApiListPrincipalsRequest {
	r.offset = &offset
	return r
}
func (r ApiListPrincipalsRequest) MatchCriteria(matchCriteria string) ApiListPrincipalsRequest {
	r.matchCriteria = &matchCriteria
	return r
}
func (r ApiListPrincipalsRequest) Usernames(usernames string) ApiListPrincipalsRequest {
	r.usernames = &usernames
	return r
}
func (r ApiListPrincipalsRequest) SortOrder(sortOrder string) ApiListPrincipalsRequest {
	r.sortOrder = &sortOrder
	return r
}
func (r ApiListPrincipalsRequest) Email(email string) ApiListPrincipalsRequest {
	r.email = &email
	return r
}
func (r ApiListPrincipalsRequest) Status(status string) ApiListPrincipalsRequest {
	r.status = &status
	return r
}
func (r ApiListPrincipalsRequest) AdminOnly(adminOnly string) ApiListPrincipalsRequest {
	r.adminOnly = &adminOnly
	return r
}
func (r ApiListPrincipalsRequest) OrderBy(orderBy string) ApiListPrincipalsRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiListPrincipalsRequest) Execute() (PrincipalPagination, *_nethttp.Response, error) {
	return r.ApiService.ListPrincipalsExecute(r)
}

/*
 * ListPrincipals List the principals for a tenant
 * By default, responses are sorted in ascending order by username
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListPrincipalsRequest
 */
func (a *PrincipalApiService) ListPrincipals(ctx _context.Context) ApiListPrincipalsRequest {
	return ApiListPrincipalsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PrincipalPagination
 */
func (a *PrincipalApiService) ListPrincipalsExecute(r ApiListPrincipalsRequest) (PrincipalPagination, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrincipalPagination
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrincipalApiService.ListPrincipals")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/principals/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.matchCriteria != nil {
		localVarQueryParams.Add("match_criteria", parameterToString(*r.matchCriteria, ""))
	}
	if r.usernames != nil {
		localVarQueryParams.Add("usernames", parameterToString(*r.usernames, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sort_order", parameterToString(*r.sortOrder, ""))
	}
	if r.email != nil {
		localVarQueryParams.Add("email", parameterToString(*r.email, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.adminOnly != nil {
		localVarQueryParams.Add("admin_only", parameterToString(*r.adminOnly, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}