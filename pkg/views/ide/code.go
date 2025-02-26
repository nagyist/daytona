// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package ide

import (
	"fmt"

	"github.com/daytonaio/daytona/cmd/daytona/config"
	"github.com/daytonaio/daytona/pkg/views"
)

func RenderIdeOpeningMessage(target, workspaceName, ideId string, ideList []config.Ide) {
	ideName := ""
	for _, ide := range ideList {
		if ide.Id == ideId {
			ideName = ide.Name
			break
		}
	}
	views.RenderInfoMessage(fmt.Sprintf("Opening workspace '%s' from target '%s' in %s", workspaceName, target, ideName))
}
