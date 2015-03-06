/*
 * Mini Object Storage, (C) 2015 Minio, Inc.
 *
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
 */

package x509_test

import (
	"testing"
	"time"

	"github.com/minio-io/minio/pkg/utils/crypto/x509"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) Testing(c *C) {
	certObj := x509.Certificates{}
	params := x509.Params{
		Hostname:   "example.com",
		IsCA:       false,
		EcdsaCurve: "P224",
		ValidFrom:  "Jan 1 15:04:05 2015",
		ValidFor:   time.Duration(3600),
	}
	err := certObj.GenerateCertificates(params)
	c.Assert(err, IsNil)
}
