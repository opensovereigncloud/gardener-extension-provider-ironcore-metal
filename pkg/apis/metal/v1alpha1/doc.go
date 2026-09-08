// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

//go:generate crd-ref-docs --source-path=. --config=../../../../hack/api-reference/api.yaml --renderer=markdown --templates-dir=$GARDENER_HACK_DIR/api-reference/template --log-level=ERROR --output-path=../../../../hack/api-reference/api.md

// +k8s:deepcopy-gen=package
// +k8s:conversion-gen=github.com/ironcore-dev/gardener-extension-provider-ironcore-metal/pkg/apis/metal
// +k8s:openapi-gen=true
// +k8s:defaulter-gen=TypeMeta

// Package v1alpha1 contains the metal provider API resources.
// +groupName=ironcore-metal.provider.extensions.gardener.cloud
package v1alpha1 // import "github.com/ironcore-dev/gardener-extension-provider-ironcore-metal/pkg/apis/metal/v1alpha1"
