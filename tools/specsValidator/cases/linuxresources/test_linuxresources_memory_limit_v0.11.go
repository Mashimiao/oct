// +build v0.1.1

//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package linuxresources

import (
	"github.com/huawei-openlab/oct/tools/specsValidator/adaptor"
	"github.com/huawei-openlab/oct/tools/specsValidator/manager"
	"github.com/opencontainers/specs"
	"time"
)

func TestMemoryLimit() string {
	var testResourceseMemory specs.Resources = specs.Resources{
		Memory: specs.Memory{
			Limit:       204800,
			Reservation: 0,
			Swap:        0,
			Kernel:      0,
			Swappiness:  -1,
		},
	}
	linuxspec, linuxruntimespec := setResources(testResourceseMemory)
	failinfo := "Memory Limit"
	go testResources(&linuxspec, &linuxruntimespec)
	time.Sleep(time.Second * 1)
	result, err := checkConfigurationFromHost("memory", "memory.limit_in_bytes", "204800", failinfo)
	adaptor.DeleteRun()
	var testResult manager.TestResult
	testResult.Set("TestMemoryLimit", testResourceseMemory.Memory, err, result)
	return testResult.Marshal()
}
