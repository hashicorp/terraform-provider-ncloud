/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:05:08Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type CreateMemberServerImageRequest struct {

	// 회원서버이미지설명
MemberServerImageDescription *string `json:"memberServerImageDescription,omitempty"`

	// 회원서버이미지명
MemberServerImageName *string `json:"memberServerImageName,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`
}
