/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PgUpgradeSpec defines the desired state of PgUpgrade
type PgUpgradeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PgUpgrade. Edit pgupgrade_types.go to remove/update
	// Foo string `json:"foo,omitempty"`

	// The container image of the PostgreSQL
	Image string `json:"image,omitempty"`

	// Namespace string `json:"namespace,omitempty"`

	// The subscription name of the PostgreSQL
	SubName string `json:"subname,omitempty"`

	// The target database name.
	DBName string `json:"dbname,omitempty"`

	// The host of the primary database.
	OldDBHost string `json:"olddbhost,omitempty"`

	// The port of the primary database.
	OldDBPort string `json:"olddbport,omitempty"`

	// The publication name in the tartget database.
	PubName string `json:"pubname,omitempty"`

	// Whether pg_dump is used to dump the data.
	PgDump bool `json:"pgdump,omitempty"`

	// // The service port of the PostgreSQL
	// ServicePort string `json:"serviceport"`
}

// PgUpgradeStatus defines the observed state of PgUpgrade
type PgUpgradeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// whether the upgrade is done
	Upgrade bool `json:"upgrade,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PgUpgrade is the Schema for the pgupgrades API
type PgUpgrade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PgUpgradeSpec   `json:"spec,omitempty"`
	Status PgUpgradeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PgUpgradeList contains a list of PgUpgrade
type PgUpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PgUpgrade `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PgUpgrade{}, &PgUpgradeList{})
}
