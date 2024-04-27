package validator

import (
	"gin-one/model/common/request"
	"testing"
)

type PageInfoTest struct {
	PageInfo request.PageInfo
	Name     string
}

func TestNotEmpty(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 0
	testInfo.PageInfo.PageSize = 0
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉0值", err)
	}
	if innerErr != nil {
		t.Error(innerErr)
	}

	testInfo.Name = ""
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能正常检测name为空", err)
	}
	if innerErr != nil {
		t.Error(innerErr)
	}
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通过检测", err)
	}
	if innerErr != nil {
		t.Error(innerErr)
	}
}

func TestLt(t *testing.T) {
	PageInfoVerify := Rules{"Page": {Lt("10")}, "PageSize": {Lt("10")}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 100
	testInfo.PageInfo.PageSize = 100
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 LT 值", innerErr)
	}

	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 9
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通 LT 过检测", innerErr)
	}
}

func TestLe(t *testing.T) {
	PageInfoVerify := Rules{"Page": {Le("10")}, "PageSize": {Le("10")}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 100
	testInfo.PageInfo.PageSize = 100
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 LE 值", innerErr)
	}

	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通 LE 过检测", innerErr)
	}
}

func TestGt(t *testing.T) {
	PageInfoVerify := Rules{"Page": {Gt("10")}, "PageSize": {Gt("10")}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 100
	testInfo.PageInfo.PageSize = 100
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通 GT 过检测", innerErr)
	}

	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 9
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 GT 值", innerErr)
	}
}

func TestGe(t *testing.T) {
	PageInfoVerify := Rules{"Page": {Gt("10")}, "PageSize": {Gt("10")}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 100
	testInfo.PageInfo.PageSize = 100
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通 GE 过检测", innerErr)
	}

	testInfo.PageInfo.Page = 10
	testInfo.PageInfo.PageSize = 10
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 GE 值", innerErr)
	}
}

func TestEq(t *testing.T) {
	PageInfoVerify := Rules{"Page": {Eq("10")}, "PageSize": {Eq("10")}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 10
	testInfo.PageInfo.PageSize = 10
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通 EQ 过检测", innerErr)
	}

	testInfo.PageInfo.Page = 100
	testInfo.PageInfo.PageSize = 10
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 EQ 值", innerErr)
	}
}

func TestNe(t *testing.T) {
	PageInfoVerify := Rules{"Page": {Ne("10")}, "PageSize": {Ne("10")}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "111"
	testInfo.PageInfo.Page = 101
	testInfo.PageInfo.PageSize = 101
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通 NE 过检测", innerErr)
	}

	testInfo.PageInfo.Page = 100
	testInfo.PageInfo.PageSize = 10
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 NE 值", innerErr)
	}
}

func TestIn(t *testing.T) {
	PageInfoVerify := Rules{"Page": {In("10", "20", "30")}, "PageSize": {In("10", "20", "30")}, "Name": {In("111", "222")}}
	var testInfo PageInfoTest
	testInfo.Name = "111"
	testInfo.PageInfo.Page = 10
	testInfo.PageInfo.PageSize = 101
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 In 值", innerErr)
	}

	testInfo.Name = "000"
	testInfo.PageInfo.Page = 10
	testInfo.PageInfo.PageSize = 30
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 In 值", innerErr)
	}

	testInfo.Name = "111"
	testInfo.PageInfo.Page = 20
	testInfo.PageInfo.PageSize = 10
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能通过 In 值", innerErr)
	}
}

func TestNotIn(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotIn("10", "20", "30")}, "PageSize": {NotIn("10", "20", "30")}, "Name": {NotIn("111", "222")}}
	var testInfo PageInfoTest
	testInfo.Name = "000"
	testInfo.PageInfo.Page = 10
	testInfo.PageInfo.PageSize = 30
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 NotIn 值", innerErr)
	}

	testInfo.PageInfo.Page = 300
	testInfo.PageInfo.PageSize = 100
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通 NotIn 过检测", innerErr)
	}

	testInfo.Name = "111"
	testInfo.PageInfo.Page = 10
	testInfo.PageInfo.PageSize = 20
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉 NotIn 值", innerErr)
	}
}

func TestRegexp(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {RegexpMatch("^[AB]")}}
	var testInfo PageInfoTest
	testInfo.Name = "111"
	testInfo.PageInfo.Page = -10
	testInfo.PageInfo.PageSize = 101
	err, innerErr := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能正常通 Regexp 过检测", innerErr)
	}

	testInfo.Name = "A123"
	testInfo.PageInfo.Page = 300
	testInfo.PageInfo.PageSize = 100
	err, innerErr = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能捕捉 Regexp 值", innerErr)
	}
}
