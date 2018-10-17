/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"flag"
	"fmt"

	crdutils "github.com/ant31/crd-validation/pkg"

	extensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kubemetal.io/kubemetal/pkg/api/v1"
)

func generateBlankCrd() *extensionsv1.CustomResourceDefinition {
	return &extensionsv1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1beta1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"kubemetal.io": "",
			},
		},
	}
}

func generateBareMetalCrd() {
	crd := generateBlankCrd()

	crd.ObjectMeta.Name = "baremetals." + v1.BareMetalGroupVersionKind.Group
	crd.Spec = extensionsv1.CustomResourceDefinitionSpec{
		Group:   v1.BareMetalGroupVersionKind.Group,
		Version: v1.BareMetalGroupVersionKind.Version,
		Scope:   "Namespaced",

		Names: extensionsv1.CustomResourceDefinitionNames{
			Plural:     "baremetals",
			Singular:   "baremetal",
			Kind:       v1.BareMetalGroupVersionKind.Kind,
			ShortNames: []string{"bm", "bms"},
		},
	}

	crdutils.MarshallCrd(crd, "yaml")
}

func main() {
	crdType := flag.String("crd-type", "", "Type of crd to generate. vm")
	flag.Parse()

	switch *crdType {
	case "vm":
		generateBareMetalCrd()
	default:
		panic(fmt.Errorf("unknown crd type %s", *crdType))
	}
}
