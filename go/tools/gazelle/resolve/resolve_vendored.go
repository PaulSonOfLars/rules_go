/* Copyright 2016 The Bazel Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resolve

import "github.com/bazelbuild/rules_go/go/tools/gazelle/config"

// vendoredResolver resolves external packages as packages in vendor/.
type vendoredResolver struct{
	prefixRoot string
}

func (v vendoredResolver) Resolve(importpath, dir string) (Label, error) {
	prefixRoot := ""
	if v.prefixRoot != "" {
		prefixRoot = v.prefixRoot + "/"
	}
	return Label{
		Pkg:  prefixRoot + "vendor/" + importpath,
		Name: config.DefaultLibName,
	}, nil
}
