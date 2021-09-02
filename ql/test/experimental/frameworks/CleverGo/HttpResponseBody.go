// Code generated by https://github.com/gagliardetto. DO NOT EDIT.

package main

import "clevergo.tech/clevergo"

// Package clevergo.tech/clevergo@v0.5.2
func HttpResponseBody_ClevergoTechClevergoV052() {
	// Response body is set via a method call (the content-type is implicit in the method name).
	{
		// Response body is set via a method call on the clevergo.tech/clevergo.Context type (the content-type is implicit in the method name).
		{
			// func (*Context).Error(code int, msg string) error
			{
				bodyString145 := source().(string)
				var rece clevergo.Context
				rece.Error(0, bodyString145) // $contentType=text/plain $responseBody=bodyString145
			}
			// func (*Context).HTML(code int, html string) error
			{
				bodyString817 := source().(string)
				var rece clevergo.Context
				rece.HTML(0, bodyString817) // $contentType=text/html $responseBody=bodyString817
			}
			// func (*Context).HTMLBlob(code int, bs []byte) error
			{
				bodyByte474 := source().([]byte)
				var rece clevergo.Context
				rece.HTMLBlob(0, bodyByte474) // $contentType=text/html $responseBody=bodyByte474
			}
			// func (*Context).JSON(code int, data interface{}) error
			{
				bodyInterface832 := source().(interface{})
				var rece clevergo.Context
				rece.JSON(0, bodyInterface832) // $contentType=application/json $responseBody=bodyInterface832
			}
			// func (*Context).JSONBlob(code int, bs []byte) error
			{
				bodyByte378 := source().([]byte)
				var rece clevergo.Context
				rece.JSONBlob(0, bodyByte378) // $contentType=application/json $responseBody=bodyByte378
			}
			// func (*Context).JSONP(code int, data interface{}) error
			{
				bodyInterface541 := source().(interface{})
				var rece clevergo.Context
				rece.JSONP(0, bodyInterface541) // $contentType=application/javascript $responseBody=bodyInterface541
			}
			// func (*Context).JSONPBlob(code int, bs []byte) error
			{
				bodyByte139 := source().([]byte)
				var rece clevergo.Context
				rece.JSONPBlob(0, bodyByte139) // $contentType=application/javascript $responseBody=bodyByte139
			}
			// func (*Context).JSONPCallback(code int, callback string, data interface{}) error
			{
				bodyInterface814 := source().(interface{})
				var rece clevergo.Context
				rece.JSONPCallback(0, "", bodyInterface814) // $contentType=application/javascript $responseBody=bodyInterface814
			}
			// func (*Context).JSONPCallbackBlob(code int, callback string, bs []byte) (err error)
			{
				bodyByte768 := source().([]byte)
				var rece clevergo.Context
				rece.JSONPCallbackBlob(0, "", bodyByte768) // $contentType=application/javascript $responseBody=bodyByte768
			}
			// func (*Context).String(code int, s string) error
			{
				bodyString468 := source().(string)
				var rece clevergo.Context
				rece.String(0, bodyString468) // $contentType=text/plain $responseBody=bodyString468
			}
			// func (*Context).StringBlob(code int, bs []byte) error
			{
				bodyByte736 := source().([]byte)
				var rece clevergo.Context
				rece.StringBlob(0, bodyByte736) // $contentType=text/plain $responseBody=bodyByte736
			}
			// func (*Context).Stringf(code int, format string, a ...interface{}) error
			{
				bodyString516 := source().(string)
				bodyInterface246 := source().(interface{})
				var rece clevergo.Context
				rece.Stringf(0, bodyString516, bodyInterface246) // $contentType=text/plain $responseBody=bodyString516 $responseBody=bodyInterface246
			}
			// func (*Context).XML(code int, data interface{}) error
			{
				bodyInterface679 := source().(interface{})
				var rece clevergo.Context
				rece.XML(0, bodyInterface679) // $contentType=text/xml $responseBody=bodyInterface679
			}
			// func (*Context).XMLBlob(code int, bs []byte) error
			{
				bodyByte736 := source().([]byte)
				var rece clevergo.Context
				rece.XMLBlob(0, bodyByte736) // $contentType=text/xml $responseBody=bodyByte736
			}
		}
	}
	// Response body and content-type are both set via a single call of a method.
	{
		// Response body and content-type are both set via a single call of a method on the clevergo.tech/clevergo.Context type.
		{
			// func (*Context).Blob(code int, contentType string, bs []byte) (err error)
			{
				bodyByte839 := source().([]byte)
				var rece clevergo.Context
				rece.Blob(0, "application/json", bodyByte839) // $contentType=application/json $responseBody=bodyByte839
			}
			// func (*Context).Emit(code int, contentType string, body string) (err error)
			{
				bodyString273 := source().(string)
				var rece clevergo.Context
				rece.Emit(0, "application/json", bodyString273) // $contentType=application/json $responseBody=bodyString273
			}
		}
	}
	// Response body is set via a call of a type method.
	{
		// Response body is set via a call of a method on the clevergo.tech/clevergo.Context type.
		{
			// func (*Context).Write(data []byte) (int, error)
			{
				bodyByte982 := source().([]byte)
				var rece clevergo.Context
				rece.Write(bodyByte982) // $responseBody=bodyByte982
			}
			// func (*Context).WriteString(data string) (int, error)
			{
				bodyString458 := source().(string)
				var rece clevergo.Context
				rece.WriteString(bodyString458) // $responseBody=bodyString458
			}
		}
	}
}