/*
 * foulefactoryapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 09/14/2016
 */
package projects_pkg


import(
	"foulefactoryapi_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"foulefactoryapi_lib"
	"foulefactoryapi_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type PROJECTS_IMPL struct { }

/**
 * Create new project
 * @param    *models_pkg.ProjectWriterServiceModel        project             parameter: Required
 * @param    *string                                      acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *PROJECTS_IMPL) CreateProjectsCreateProject (
            project *models_pkg.ProjectWriterServiceModel,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/projects"

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
    _request := unirest.PostWithAuth(_queryBuilder, headers, project, foulefactoryapi_lib.Config.BasicAuthUserName, foulefactoryapi_lib.Config.BasicAuthPassword)
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
 * Get All projects
 * @return	Returns the interface{} response from the API call
 */
func (me *PROJECTS_IMPL) GetProjectsGetUserProjects () (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/projects"

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
 * Get csv files by project id
 * @param    int64          id                  parameter: Required
 * @param    *string        acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *PROJECTS_IMPL) GetProjectsGetProjectFiles (
            id int64,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/projects/{id}/urlCsvFiles"

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
 * Get task lines by project id
 * @param    int64          id                  parameter: Required
 * @param    *string        acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *PROJECTS_IMPL) GetProjectsGetProjectTaskLines (
            id int64,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/projects/{id}/taskLines"

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
 * Get tasks by project id
 * @param    int64          id                  parameter: Required
 * @param    *string        acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *PROJECTS_IMPL) GetProjectsGetProjectTasks (
            id int64,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/projects/{id}/tasks"

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
 * Get project by id
 * @param    int64          id                  parameter: Required
 * @param    *string        acceptLanguage      parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *PROJECTS_IMPL) GetProjectsGet (
            id int64,
            acceptLanguage *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := foulefactoryapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1.1/projects/{id}"

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

