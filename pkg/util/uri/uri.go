/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package uri

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/apache/camel-k/pkg/util/log"
)

var queryExtractorRegexp = `^[^?]+\?(?:|.*[&])%s=([^&]+)(?:[&].*|$)`

// GetComponent returns the Camel component used in the URI
func GetComponent(uri string) string {
	parts := strings.Split(uri, ":")
	if len(parts) <= 1 {
		return ""
	}
	return parts[0]
}

// GetQueryParameter returns the given parameter from the uri, if present
func GetQueryParameter(uri string, param string) string {
	paramRegexp := regexp.MustCompile(fmt.Sprintf(queryExtractorRegexp, regexp.QuoteMeta(param)))
	val := matchOrEmpty(paramRegexp, uri)
	res, err := url.QueryUnescape(val)
	if err != nil {
		log.Error(err, fmt.Sprintf("Invalid character sequence in parameter %q", param))
		return ""
	}
	return res
}

func matchOrEmpty(reg *regexp.Regexp, str string) string {
	match := reg.FindStringSubmatch(str)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
