// Copyright © 2016 Abcum Ltd
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

package db

import (
	"testing"

	"github.com/abcum/surreal/util/data"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInfo(t *testing.T) {

	Convey("Info for namespace", t, func() {

		setupDB()

		txt := `
		USE NS test DB test;
		DEFINE LOGIN test ON NAMESPACE PASSWORD "test";
		DEFINE TOKEN test ON NAMESPACE TYPE HS512 VALUE "test";
		DEFINE DATABASE test;
		INFO FOR NAMESPACE;
		REMOVE LOGIN test ON NAMESPACE;
		REMOVE TOKEN test ON NAMESPACE;
		REMOVE DATABASE test;
		INFO FOR NAMESPACE;
		`

		res, err := Execute(setupKV(), txt, nil)
		So(err, ShouldBeNil)
		So(res, ShouldHaveLength, 9)
		So(res[1].Status, ShouldEqual, "OK")
		So(res[2].Status, ShouldEqual, "OK")
		So(res[3].Status, ShouldEqual, "OK")
		So(res[4].Status, ShouldEqual, "OK")
		So(data.Consume(res[4].Result[0]).Get("login").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[4].Result[0]).Get("login.test").Data(), ShouldEqual, "DEFINE LOGIN test ON NAMESPACE PASSWORD ********")
		So(data.Consume(res[4].Result[0]).Get("token").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[4].Result[0]).Get("token.test").Data(), ShouldEqual, "DEFINE TOKEN test ON NAMESPACE TYPE HS512 VALUE ********")
		So(data.Consume(res[4].Result[0]).Get("database").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[4].Result[0]).Get("database.test").Data(), ShouldEqual, "DEFINE DATABASE test")
		So(res[5].Status, ShouldEqual, "OK")
		So(res[6].Status, ShouldEqual, "OK")
		So(res[7].Status, ShouldEqual, "OK")
		So(res[8].Status, ShouldEqual, "OK")
		So(data.Consume(res[8].Result[0]).Get("login").Data(), ShouldHaveLength, 0)
		So(data.Consume(res[8].Result[0]).Get("token").Data(), ShouldHaveLength, 0)
		So(data.Consume(res[8].Result[0]).Get("database").Data(), ShouldHaveLength, 0)

	})

	Convey("Info for database", t, func() {

		setupDB()

		txt := `
		USE NS test DB test;
		DEFINE LOGIN test ON DATABASE PASSWORD "test";
		DEFINE TOKEN test ON DATABASE TYPE HS512 VALUE "test";
		DEFINE SCOPE test;
		DEFINE TABLE test;
		INFO FOR DATABASE;
		REMOVE LOGIN test ON DATABASE;
		REMOVE TOKEN test ON DATABASE;
		REMOVE SCOPE test;
		REMOVE TABLE test;
		INFO FOR DATABASE;
		`

		res, err := Execute(setupKV(), txt, nil)
		So(err, ShouldBeNil)
		So(res, ShouldHaveLength, 11)
		So(res[1].Status, ShouldEqual, "OK")
		So(res[2].Status, ShouldEqual, "OK")
		So(res[3].Status, ShouldEqual, "OK")
		So(res[4].Status, ShouldEqual, "OK")
		So(res[5].Status, ShouldEqual, "OK")
		So(data.Consume(res[5].Result[0]).Get("login").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[5].Result[0]).Get("login.test").Data(), ShouldEqual, "DEFINE LOGIN test ON DATABASE PASSWORD ********")
		So(data.Consume(res[5].Result[0]).Get("token").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[5].Result[0]).Get("token.test").Data(), ShouldEqual, "DEFINE TOKEN test ON DATABASE TYPE HS512 VALUE ********")
		So(data.Consume(res[5].Result[0]).Get("scope").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[5].Result[0]).Get("scope.test").Data(), ShouldEqual, "DEFINE SCOPE test")
		So(data.Consume(res[5].Result[0]).Get("table").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[5].Result[0]).Get("table.test").Data(), ShouldEqual, "DEFINE TABLE test")
		So(res[6].Status, ShouldEqual, "OK")
		So(res[7].Status, ShouldEqual, "OK")
		So(res[8].Status, ShouldEqual, "OK")
		So(res[9].Status, ShouldEqual, "OK")
		So(res[10].Status, ShouldEqual, "OK")
		So(data.Consume(res[10].Result[0]).Get("login").Data(), ShouldHaveLength, 0)
		So(data.Consume(res[10].Result[0]).Get("token").Data(), ShouldHaveLength, 0)
		So(data.Consume(res[10].Result[0]).Get("scope").Data(), ShouldHaveLength, 0)
		So(data.Consume(res[10].Result[0]).Get("table").Data(), ShouldHaveLength, 0)

	})

	Convey("Info for table", t, func() {

		setupDB()

		txt := `
		USE NS test DB test;
		DEFINE EVENT test ON test WHEN true THEN null;
		DEFINE FIELD test ON test;
		DEFINE INDEX test ON test COLUMNS id;
		INFO FOR TABLE test;
		REMOVE EVENT test ON test;
		REMOVE FIELD test ON test;
		REMOVE INDEX test ON test;
		INFO FOR TABLE test;
		`

		res, err := Execute(setupKV(), txt, nil)
		So(err, ShouldBeNil)
		So(res, ShouldHaveLength, 9)
		So(res[1].Status, ShouldEqual, "OK")
		So(res[2].Status, ShouldEqual, "OK")
		So(res[3].Status, ShouldEqual, "OK")
		So(res[4].Status, ShouldEqual, "OK")
		So(data.Consume(res[4].Result[0]).Get("event").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[4].Result[0]).Get("event.test").Data(), ShouldEqual, "DEFINE EVENT test ON test WHEN true THEN NULL")
		So(data.Consume(res[4].Result[0]).Get("field").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[4].Result[0]).Get("field.test").Data(), ShouldEqual, "DEFINE FIELD test ON test")
		So(data.Consume(res[4].Result[0]).Get("index").Data(), ShouldHaveLength, 1)
		So(data.Consume(res[4].Result[0]).Get("index.test").Data(), ShouldEqual, "DEFINE INDEX test ON test COLUMNS id")
		So(res[5].Status, ShouldEqual, "OK")
		So(res[6].Status, ShouldEqual, "OK")
		So(res[7].Status, ShouldEqual, "OK")
		So(res[8].Status, ShouldEqual, "OK")
		So(data.Consume(res[8].Result[0]).Get("event").Data(), ShouldHaveLength, 0)
		So(data.Consume(res[8].Result[0]).Get("field").Data(), ShouldHaveLength, 0)
		So(data.Consume(res[8].Result[0]).Get("index").Data(), ShouldHaveLength, 0)

	})

}