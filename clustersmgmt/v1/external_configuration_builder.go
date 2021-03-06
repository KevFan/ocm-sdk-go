/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ExternalConfigurationBuilder contains the data and logic needed to build 'external_configuration' objects.
//
// Representation of cluster external configuration.
type ExternalConfigurationBuilder struct {
	syncsets *SyncsetListBuilder
}

// NewExternalConfiguration creates a new builder of 'external_configuration' objects.
func NewExternalConfiguration() *ExternalConfigurationBuilder {
	return new(ExternalConfigurationBuilder)
}

// Syncsets sets the value of the 'syncsets' attribute to the given values.
//
//
func (b *ExternalConfigurationBuilder) Syncsets(value *SyncsetListBuilder) *ExternalConfigurationBuilder {
	b.syncsets = value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ExternalConfigurationBuilder) Copy(object *ExternalConfiguration) *ExternalConfigurationBuilder {
	if object == nil {
		return b
	}
	if object.syncsets != nil {
		b.syncsets = NewSyncsetList().Copy(object.syncsets)
	} else {
		b.syncsets = nil
	}
	return b
}

// Build creates a 'external_configuration' object using the configuration stored in the builder.
func (b *ExternalConfigurationBuilder) Build() (object *ExternalConfiguration, err error) {
	object = new(ExternalConfiguration)
	if b.syncsets != nil {
		object.syncsets, err = b.syncsets.Build()
		if err != nil {
			return
		}
	}
	return
}
