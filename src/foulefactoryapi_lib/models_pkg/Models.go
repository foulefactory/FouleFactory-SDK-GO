/*
 * foulefactoryapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 09/14/2016
 */

package models_pkg

import "time"

/*
 * Structure for the custom type AccountWriterServiceModel
 */
type AccountWriterServiceModel struct {
    IdGender        int64           `json:"IdGender" form:"IdGender"` //TODO: Write general description for this field
    FirstName       string          `json:"FirstName" form:"FirstName"` //TODO: Write general description for this field
    Name            string          `json:"Name" form:"Name"` //TODO: Write general description for this field
    Email           string          `json:"Email" form:"Email"` //TODO: Write general description for this field
    Phone           string          `json:"Phone" form:"Phone"` //TODO: Write general description for this field
    Birthday        *time.Time      `json:"Birthday" form:"Birthday"` //TODO: Write general description for this field
    Address1        string          `json:"Address1" form:"Address1"` //TODO: Write general description for this field
    City            string          `json:"City" form:"City"` //TODO: Write general description for this field
    PostalCode      string          `json:"PostalCode" form:"PostalCode"` //TODO: Write general description for this field
    CountryCode     string          `json:"CountryCode" form:"CountryCode"` //TODO: Write general description for this field
    Nationality     string          `json:"Nationality" form:"Nationality"` //TODO: Write general description for this field
    Optin           bool            `json:"Optin" form:"Optin"` //TODO: Write general description for this field
    Company         *string         `json:"Company,omitempty" form:"Company,omitempty"` //TODO: Write general description for this field
    Address2        *string         `json:"Address2,omitempty" form:"Address2,omitempty"` //TODO: Write general description for this field
    BillAddress1    *string         `json:"BillAddress1,omitempty" form:"BillAddress1,omitempty"` //TODO: Write general description for this field
    BillAddress2    *string         `json:"BillAddress2,omitempty" form:"BillAddress2,omitempty"` //TODO: Write general description for this field
    BillCity        *string         `json:"BillCity,omitempty" form:"BillCity,omitempty"` //TODO: Write general description for this field
    BillPostalCode  *string         `json:"BillPostalCode,omitempty" form:"BillPostalCode,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type PayinServiceModel
 */
type PayinServiceModel struct {
    Amount          int64           `json:"Amount" form:"Amount"` //TODO: Write general description for this field
    ReturnUrl       *string         `json:"ReturnUrl,omitempty" form:"ReturnUrl,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type CsvFileWriterServiceModel
 */
type CsvFileWriterServiceModel struct {
    IdProject       int64           `json:"IdProject" form:"IdProject"` //TODO: Write general description for this field
    File            string          `json:"File" form:"File"` //TODO: Write general description for this field
    Header          bool            `json:"Header" form:"Header"` //TODO: Write general description for this field
    Separator       string          `json:"Separator" form:"Separator"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ProjectWriterServiceModel
 */
type ProjectWriterServiceModel struct {
    IdTemplate              int64           `json:"IdTemplate" form:"IdTemplate"` //TODO: Write general description for this field
    Title                   string          `json:"Title" form:"Title"` //TODO: Write general description for this field
    EstimatedTimePerTask    string          `json:"EstimatedTimePerTask" form:"EstimatedTimePerTask"` //TODO: Write general description for this field
    NbSupplierPerTask       int64           `json:"NbSupplierPerTask" form:"NbSupplierPerTask"` //TODO: Write general description for this field
    AmountWithoutTaxPerTask int64           `json:"AmountWithoutTaxPerTask" form:"AmountWithoutTaxPerTask"` //TODO: Write general description for this field
    AutomaticValidation     bool            `json:"AutomaticValidation" form:"AutomaticValidation"` //TODO: Write general description for this field
    MaxEndDate              *time.Time      `json:"MaxEndDate,omitempty" form:"MaxEndDate,omitempty"` //TODO: Write general description for this field
    IdCertification         *int64          `json:"IdCertification,omitempty" form:"IdCertification,omitempty"` //TODO: Write general description for this field
    UrlNotification         *string         `json:"UrlNotification,omitempty" form:"UrlNotification,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TaskLinesWriterServiceModel
 */
type TaskLinesWriterServiceModel struct {
    IdProject       int64           `json:"IdProject" form:"IdProject"` //TODO: Write general description for this field
    TaskColumns     []string        `json:"TaskColumns" form:"TaskColumns"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TaskValidationWriterServiceModel
 */
type TaskValidationWriterServiceModel struct {
    IdTask          *int64          `json:"IdTask,omitempty" form:"IdTask,omitempty"` //TODO: Write general description for this field
    State           *string         `json:"State,omitempty" form:"State,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TemplateNewWriterServiceModel
 */
type TemplateNewWriterServiceModel struct {
    IdProjectType   int64           `json:"IdProjectType" form:"IdProjectType"` //TODO: Write general description for this field
    Title           string          `json:"Title" form:"Title"` //TODO: Write general description for this field
    Description     string          `json:"Description" form:"Description"` //TODO: Write general description for this field
    Instructions    []*TemplateInstructionWriterServiceModel `json:"Instructions" form:"Instructions"` //TODO: Write general description for this field
    Columns         []*TemplateColumnWriterServiceModel `json:"Columns" form:"Columns"` //TODO: Write general description for this field
    Questions       []*TemplateQuestionWriterServiceModel `json:"Questions" form:"Questions"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TemplateInstructionWriterServiceModel
 */
type TemplateInstructionWriterServiceModel struct {
    Instruction     string          `json:"Instruction" form:"Instruction"` //TODO: Write general description for this field
    Order           int64           `json:"Order" form:"Order"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TemplateColumnWriterServiceModel
 */
type TemplateColumnWriterServiceModel struct {
    IdTemplateColumnType int64           `json:"IdTemplateColumnType" form:"IdTemplateColumnType"` //TODO: Write general description for this field
    Column               int64           `json:"Column" form:"Column"` //TODO: Write general description for this field
    Order                int64           `json:"Order" form:"Order"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type TemplateQuestionWriterServiceModel
 */
type TemplateQuestionWriterServiceModel struct {
    Title                        string          `json:"Title" form:"Title"` //TODO: Write general description for this field
    Require                      bool            `json:"Require" form:"Require"` //TODO: Write general description for this field
    IdTemplateObjectQuestionType int64           `json:"IdTemplateObjectQuestionType" form:"IdTemplateObjectQuestionType"` //TODO: Write general description for this field
    Order                        int64           `json:"Order" form:"Order"` //TODO: Write general description for this field
    Option                       *string         `json:"Option,omitempty" form:"Option,omitempty"` //TODO: Write general description for this field
    Choices                      *[]string       `json:"Choices,omitempty" form:"Choices,omitempty"` //TODO: Write general description for this field
}
