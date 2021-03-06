/*
 * CS API
 *
 * This is a simple API
 *
 * API version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Game struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Slug string `json:"slug,omitempty"`

	Link string `json:"link,omitempty"`

	Active bool `json:"active,omitempty"`
}
