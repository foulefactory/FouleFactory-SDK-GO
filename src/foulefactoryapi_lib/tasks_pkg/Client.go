/*
 * foulefactoryapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 09/14/2016
 */
package tasks_pkg


import(
	"foulefactoryapi_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"foulefactoryapi_lib"
	"foulefactoryapi_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type TASKS_IMPL struct { }

/**
 * Validate task
 * @param    *models_pkg.TaskValidationWriterServiceModel        task                parameter: Required
 * @param    *string                                             acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *TASKS_IMPL) UpdateTasksValidate (
            task *models_pkg.TaskValidationWriterServiceModel,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/tasks"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Accept-Language" : apihelper_pkg.ToString(*acceptLanguage, "fr"),
    }

    //prepare API request
    _request := unirest.PutWithAuth(_queryBuilder, headers, task, foulefactoryapi_lib.Config.BasicAuthUserName, foulefactoryapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Get task answer choices by task id
 * @param    int64          id                  parameter: Required
 * @param    *string        acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *TASKS_IMPL) GetTasksGetAnswerChoices (
            id int64,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/tasks/{id}/taskAnswerChoices"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "Accept-Language" : apihelper_pkg.ToString(*acceptLanguage, "fr"),
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, foulefactoryapi_lib.Config.BasicAuthUserName, foulefactoryapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Get task answer texts by task id
 * @param    int64          id                  parameter: Required
 * @param    *string        acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *TASKS_IMPL) GetTasksGetAnswerTexts (
            id int64,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/tasks/{id}/taskAnswerTexts"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "Accept-Language" : apihelper_pkg.ToString(*acceptLanguage, "fr"),
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, foulefactoryapi_lib.Config.BasicAuthUserName, foulefactoryapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Get task by id
 * @param    int64          id                  parameter: Required
 * @param    *string        acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *TASKS_IMPL) GetTasksGet (
            id int64,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/tasks/{id}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "Accept-Language" : apihelper_pkg.ToString(*acceptLanguage, "fr"),
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, foulefactoryapi_lib.Config.BasicAuthUserName, foulefactoryapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

