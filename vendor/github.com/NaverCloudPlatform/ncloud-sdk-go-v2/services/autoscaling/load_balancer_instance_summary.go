/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-08-07T06:47:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type LoadBalancerInstanceSummary struct {

	// 로드밸런서인스턴스번호
LoadBalancerInstanceNo *string `json:"loadBalancerInstanceNo,omitempty"`

	// 로드밸런서명
LoadBalancerName *string `json:"loadBalancerName,omitempty"`
}
