// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package validator_test

import (
	"context"

	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	"github.com/gardener/gardener/pkg/apis/core"
	testutils "github.com/gardener/gardener/pkg/utils/test"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/ironcore-dev/gardener-extension-provider-ironcore-metal/pkg/admission/validator"
)

var _ = Describe("SecretBinding validator", func() {

	Describe("#Validate", func() {
		const (
			namespace = "garden-dev"
			name      = "my-provider-account"
		)

		var (
			ctx           = context.TODO()
			secretBinding = &core.SecretBinding{
				SecretRef: corev1.SecretReference{
					Name:      name,
					Namespace: namespace,
				},
			}

			scheme = func() *runtime.Scheme {
				s := runtime.NewScheme()
				Expect(corev1.AddToScheme(s)).To(Succeed())
				return s
			}()
		)

		newValidator := func(objects ...client.Object) extensionswebhook.Validator {
			b := fakeclient.NewClientBuilder().WithScheme(scheme)
			if len(objects) > 0 {
				b = b.WithObjects(objects...)
			}
			mgr := &testutils.FakeManager{APIReader: b.Build()}
			return validator.NewSecretBindingValidator(mgr)
		}

		It("should return err when obj is not a SecretBinding", func() {
			secretBindingValidator := newValidator()
			err := secretBindingValidator.Validate(ctx, &corev1.Secret{}, nil)
			Expect(err).To(MatchError("wrong object type *v1.Secret"))
		})

		It("should return err when oldObj is not a SecretBinding", func() {
			secretBindingValidator := newValidator()
			err := secretBindingValidator.Validate(ctx, &core.SecretBinding{}, &corev1.Secret{})
			Expect(err).To(MatchError("wrong object type *v1.Secret for old object"))
		})

		It("should return err if it fails to get the corresponding Secret", func() {
			// no secret pre-loaded → Get returns NotFound
			secretBindingValidator := newValidator()
			err := secretBindingValidator.Validate(ctx, secretBinding, nil)
			Expect(err).To(HaveOccurred())
		})

		It("should return err when the corresponding Secret is not valid", func() {
			secret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
				Data: map[string][]byte{
					"namespace": []byte("foo"),
				},
			}
			secretBindingValidator := newValidator(secret)
			err := secretBindingValidator.Validate(ctx, secretBinding, nil)
			Expect(err).To(MatchError("missing field: token in cloud provider secret"))
		})

		It("should return nil when the corresponding Secret is valid", func() {
			secret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
				Data: map[string][]byte{
					"namespace": []byte("default"),
					"token":     []byte("abcd"),
					"username":  []byte("admin"),
				},
			}
			secretBindingValidator := newValidator(secret)
			err := secretBindingValidator.Validate(ctx, secretBinding, nil)
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
