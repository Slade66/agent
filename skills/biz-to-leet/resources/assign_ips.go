package problems

import "errors"

// AllocateIPs 从指定网段中，跳过已占用的 IP，按顺序分配 n 个可用地址。
//
// 参数：
//   - subnet: CIDR 格式的网段，如 "192.168.1.0/24"
//   - used:   已被占用的 IP 列表（字符串格式，如 "192.168.1.1"）
//   - n:      需要分配的 IP 数量
//
// 返回：
//   - []string: 按升序排列的 n 个可用 IP
//   - error:    若可用 IP 不足，返回具体原因
func AllocateIPs(subnet string, used []string, n int) ([]string, error) {
	panic("not implemented")
	_ = errors.New // 占位，避免 import 报错
	return nil, nil
}
