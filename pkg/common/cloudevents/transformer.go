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
    "fmt"
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/binding/spec"
	"github.com/cloudevents/sdk-go/v2/types"
)

// IDExtractorTransformer implements binding.Transformer. Upon the execution of the transformer, the underlying value
// is updated to be event ID.
type IDExtractorTransformer string
type CloudeventExtractorTransformer map[string]string


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


func (a *CloudeventExtractorTransformer) Transform(reader binding.MessageMetadataReader, _ binding.MessageMetadataWriter) error {
    spec_fields := []spec.Kind{
        spec.ID,
        spec.Source,
        spec.SpecVersion,
        spec.Type,
        spec.DataContentType,
        spec.DataSchema,
        spec.Subject,
        spec.Time,
    }
    // fetch extensions
    // TODO: (@omgbebebe) dunno how to get list of event extensions...

    // fetch spec
    for _, f := range spec_fields {
	    _, ty := reader.GetAttribute(f)
	    if ty != nil {
		    tyParsed, err := types.ToString(ty)
		    if err != nil {
			    return err
		    }
            (*a)[fmt.Sprintf("ce-%s", f.String())] = tyParsed
        }
    }


	return nil
}
