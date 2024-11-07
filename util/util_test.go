/*
Copyright 2016 The Kubernetes Authors.

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

package util

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestShouldEnqueuePersistentVolumeChange(t *testing.T) {
	oldValue := "old"
	newValue := "new"

	testcases := []struct {
		name           string
		old            *v1.PersistentVolume
		new            *v1.PersistentVolume
		expectedResult bool
	}{
		{
			name:           "basic no change",
			old:            &v1.PersistentVolume{},
			new:            &v1.PersistentVolume{},
			expectedResult: false,
		},
		{
			name: "basic change",
			old: &v1.PersistentVolume{
				Spec: v1.PersistentVolumeSpec{
					PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimDelete,
				},
			},
			new: &v1.PersistentVolume{
				Spec: v1.PersistentVolumeSpec{
					PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
				},
			},
			expectedResult: true,
		},
		{
			name: "finalizers change",
			old: &v1.PersistentVolume{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{
						oldValue,
					},
				},
			},
			new: &v1.PersistentVolume{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{
						newValue,
					},
				},
			},
			expectedResult: false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := ShouldEnqueuePersistentVolumeChange(tc.old, tc.new)
			if result != tc.expectedResult {
				t.Fatalf("Incorrect result: Expected %v received %v", tc.expectedResult, result)
			}
		})
	}
}
