/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudevents

import (
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/binding/spec"
	"github.com/cloudevents/sdk-go/v2/types"
)

// IDExtractorTransformer implements binding.Transformer. Upon the execution of the transformer, the underlying value
// is updated to be event ID.
type IDExtractorTransformer string
type SourceExtractorTransformer string
type SubjectExtractorTransformer string
type TypeExtractorTransformer string
type SpecVersionExtractorTransformer string

func (a *IDExtractorTransformer) Transform(reader binding.MessageMetadataReader, _ binding.MessageMetadataWriter) error {
	_, ty := reader.GetAttribute(spec.ID)
	if ty != nil {
		tyParsed, err := types.ToString(ty)
		if err != nil {
			return err
		}
		*a = IDExtractorTransformer(tyParsed)
	}

	return nil
}

func (a *SourceExtractorTransformer) Transform(reader binding.MessageMetadataReader, _ binding.MessageMetadataWriter) error {
	_, ty := reader.GetAttribute(spec.Source)
	if ty != nil {
		tyParsed, err := types.ToString(ty)
		if err != nil {
			return err
		}
		*a = SourceExtractorTransformer(tyParsed)
	}

	return nil
}

func (a *SubjectExtractorTransformer) Transform(reader binding.MessageMetadataReader, _ binding.MessageMetadataWriter) error {
	_, ty := reader.GetAttribute(spec.Subject)
	if ty != nil {
		tyParsed, err := types.ToString(ty)
		if err != nil {
			return err
		}
		*a = SubjectExtractorTransformer(tyParsed)
	}

	return nil
}

func (a *TypeExtractorTransformer) Transform(reader binding.MessageMetadataReader, _ binding.MessageMetadataWriter) error {
	_, ty := reader.GetAttribute(spec.Type)
	if ty != nil {
		tyParsed, err := types.ToString(ty)
		if err != nil {
			return err
		}
		*a = TypeExtractorTransformer(tyParsed)
	}

	return nil
}

func (a *SpecVersionExtractorTransformer) Transform(reader binding.MessageMetadataReader, _ binding.MessageMetadataWriter) error {
	_, ty := reader.GetAttribute(spec.SpecVersion)
	if ty != nil {
		tyParsed, err := types.ToString(ty)
		if err != nil {
			return err
		}
		*a = SpecVersionExtractorTransformer(tyParsed)
	}

	return nil
}


