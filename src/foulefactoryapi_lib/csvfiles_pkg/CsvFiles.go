/*
 * foulefactoryapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 09/14/2016
 */

package csvfiles_pkg

import "foulefactoryapi_lib/models_pkg"

/*
 * Interface for the CSVFILES_IMPL
 */
type CSVFILES interface {
    CreateCsvFilesCreateCsvFile (*models_pkg.CsvFileWriterServiceModel, *string) (interface{}, error)

    GetCsvFilesGet (int64, *string) (interface{}, error)
}

/*
 * Factory for the CSVFILES interaface returning CSVFILES_IMPL
 */
func NewCSVFILES() CSVFILES {
    return &CSVFILES_IMPL{}
}