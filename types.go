/**
 * @Time : 2019-09-03 17:27
 * @Author : solacowa@gmail.com
 * @File : types
 * @Software: GoLand
 */

package main

type HttpReq struct {
	HttpSessionId string `json:"http_session_id"`
	ScriptSessionId string `json:"script_session_id"`
}
