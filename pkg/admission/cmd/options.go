// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"

	"github.com/ironcore-dev/gardener-extension-provider-ironcore-metal/pkg/admission/validator"
)

// GardenWebhookSwitchOptions are the webhookcmd.SwitchOptions for the admission webhooks.
func GardenWebhookSwitchOptions() *webhookcmd.SwitchOptions {
	return webhookcmd.NewSwitchOptions(
		webhookcmd.Switch(validator.Name, validator.New),
		webhookcmd.Switch(validator.SecretsValidatorName, validator.NewSecretsWebhook),
	)
}
