// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build api

package lifecycle

import "github.com/agilestacks/hub/cmd/hub/api"

func getParameterOrMaybeCreateSecret(environment, stackInstance, application,
	name, component string, create bool) (bool, string, []error) {
	return api.GetParameterOrMaybeCreateSecret(environment, stackInstance, application, name, component, create)
}
