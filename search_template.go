// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/CharlotteFTD/elastic/v5/uritemplates"
)

// SearchTemplateService reads a search template.
// It is documented at https://www.elastic.co/guide/en/elasticsearch/reference/5.2/search-template.html.
type SearchTemplateService struct {
	client               *Client
//	searchTemplateSource *SearchTemplateSource
	pretty               bool
//	version              interface{}
//	versionType          string
	index                []string
	typ                  []string
	source               interface{}
}

// NewSearchTemplateService creates a new SearchTemplateService.
func NewSearchTemplateService(client *Client) *SearchTemplateService {
	return &SearchTemplateService{
		client:               client,
//		searchTemplateSource: NewSearchTemplateSource(),
	}
}

/*
// SearchSource sets the search source builder to use with this service.
func (s *SearchTemplateService) SearchTemplateSource(searchTemplateSource *SearchTemplateSource) *SearchTemplateService {
	s.searchTemplateSource = searchTemplateSource
	if s.searchTemplateSource == nil {
		s.searchTemplateSource = NewSearchTemplateSource()
	}
	return s
}

// Id is the template ID.
func (s *SearchTemplateService) Id(id string) *SearchTemplateService {
	s.searchTemplateSource = s.searchTemplateSource.Id(id)
	return s
}
*/

/*
// Version is an explicit version number for concurrency control.
func (s *SearchTemplateService) Version(version interface{}) *SearchTemplateService {
	s.version = version
	return s
}

// VersionType is a specific version type.
func (s *SearchTemplateService) VersionType(versionType string) *SearchTemplateService {
	s.versionType = versionType
	return s
}
*/

/*
// Param is an explicit Params for concurrency control.
func (s *SearchTemplateService) Phrase(phrase string) *SearchTemplateService {
	s.searchTemplateSource = s.searchTemplateSource.Phrase(phrase)
	return s
}


// Param is an explicit Params for concurrency control.
func (s *SearchTemplateService) From(from int) *SearchTemplateService {
	s.searchTemplateSource = s.searchTemplateSource.From(from)
	return s
}

// Param is an explicit Params for concurrency control.
func (s *SearchTemplateService) Size(size int) *SearchTemplateService {
	s.searchTemplateSource = s.searchTemplateSource.Size(size)
	return s
}

// Param is an explicit Params for concurrency control.
func (s *SearchTemplateService) Include(includes ...string) *SearchTemplateService {
	s.searchTemplateSource = s.searchTemplateSource.Include(includes...)
	return s
}

// Param is an explicit Params for concurrency control.
func (s *SearchTemplateService) Param(name string, value interface{}) *SearchTemplateService {
	s.searchTemplateSource = s.searchTemplateSource.Param(name, value)
	return s
}

// Params is an explicit Params for concurrency control.
func (s *SearchTemplateService) Params(params map[string]interface{}) *SearchTemplateService {
	s.searchTemplateSource = s.searchTemplateSource.Params(params)
	return s
}
*/

// Source allows the user to set the request body manually without using
// any of the structs and interfaces in Elastic.
func (s *SearchTemplateService) Source(source interface{}) *SearchTemplateService {
	s.source = source
	return s
}

// Index sets the names of the indices to use for search.
func (s *SearchTemplateService) Index(index ...string) *SearchTemplateService {
	s.index = append(s.index, index...)
	return s
}

// Types adds search restrictions for a list of types.
func (s *SearchTemplateService) Type(typ ...string) *SearchTemplateService {
	s.typ = append(s.typ, typ...)
	return s
}

// buildURL builds the URL for the operation.
func (s *SearchTemplateService) buildURL() (string, url.Values, error) {
	// Build URL
	var err error
	var path string

	if len(s.index) > 0 && len(s.typ) > 0 {
		path, err = uritemplates.Expand("/{index}/{type}/_search/template", map[string]string{
			"index": strings.Join(s.index, ","),
			"type":  strings.Join(s.typ, ","),
		})
	} else if len(s.index) > 0 {
		path, err = uritemplates.Expand("/{index}/_search/template", map[string]string{
			"index": strings.Join(s.index, ","),
		})
	} else if len(s.typ) > 0 {
		path, err = uritemplates.Expand("/_all/{type}/_search/template", map[string]string{
			"type": strings.Join(s.typ, ","),
		})
	} else {
		path = "/_search"
	}
	if err != nil {
		return "", url.Values{}, err
	}

	// Add query string parameters
	params := url.Values{}
	if s.pretty {
		params.Set("pretty", fmt.Sprintf("%v", s.pretty))
	}
	/*
	if s.version != nil {
		params.Set("version", fmt.Sprintf("%v", s.version))
	}
	if s.versionType != "" {
		params.Set("version_type", s.versionType)
	}
	*/
	return path, params, nil
}

// Validate checks if the operation is valid.
func (s *SearchTemplateService) Validate() error {
	var invalid []string
	if s.source == nil {
		invalid = append(invalid, "Source")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do executes the operation and returns the template.
func (s *SearchTemplateService) Do(ctx context.Context) (*SearchResult, error) {
	// Check pre-conditions
	if err := s.Validate(); err != nil {
		return nil, err
	}

	// Get URL for request
	path, params, err := s.buildURL()
	if err != nil {
		return nil, err
	}

	// Perform request
	var body interface{}
/*
	if s.source != nil {
		body = s.source
	} else {
		src, err := s.searchTemplateSource.Source()
		if err != nil {
			return nil, err
		}
		body = src
	}
*/
	body = s.source

	// Get HTTP response
	res, err := s.client.PerformRequest(ctx, "POST", path, params, body)
	if err != nil {
		return nil, err
	}

	// Return result
	ret := new(SearchResult)
	if err := s.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
/*
type SearchTemplateSource struct {
	id     string
	params map[string]interface{}
}

func NewSearchTemplateSource() *SearchTemplateSource {
	return &SearchTemplateSource{
		params: make(map[string]interface{}),
	}
}

func (s *SearchTemplateSource) Id(id string) *SearchTemplateSource {
	s.id = id
	return s
}

func (s *SearchTemplateSource) Phrase(phrase string) *SearchTemplateSource {
	s.Param("phrase", phrase)
	return s
}

func (s *SearchTemplateSource) From(from int) *SearchTemplateSource {
	s.Param("from", from)
	return s
}

func (s *SearchTemplateSource) Size(size int) *SearchTemplateSource {
	s.Param("size", size)
	return s
}

func (s *SearchTemplateSource) Include(includes ...string) *SearchTemplateSource {
	if len(includes) > 0 {
		s.Param("source", includes)
	}
	return s
}

func (s *SearchTemplateSource) Param(name string, value interface{}) *SearchTemplateSource {
	if s.params == nil {
		s.params = make(map[string]interface{})
	}
	s.params[name] = value
	return s
}

func (s *SearchTemplateSource) Params(params map[string]interface{}) *SearchTemplateSource {
	s.params = params
	return s
}

func (s *SearchTemplateSource) Source() (interface{}, error) {
	if s.id == "" {
		return nil, errors.New("SearchTemplateSource expects id")
	}
	source := make(map[string]interface{})
	source["id"] = s.id
	if len(s.params) > 0 {
		source["params"] = s.params
	}
	return source, nil
}
*/
