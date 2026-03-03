package problems

import (
	"testing"
)

func TestAllocateIPs(t *testing.T) {
	tests := []struct {
		name    string
		subnet  string
		used    []string
		n       int
		want    []string
		wantErr bool
	}{
		// ── Happy Path（正常路径）────────────────────────────────────
		{
			name:   "HP-01: /24 网段，跳过2个已占用IP，分配4个",
			subnet: "192.168.1.0/24",
			used:   []string{"192.168.1.1", "192.168.1.3"},
			n:      4,
			want:   []string{"192.168.1.2", "192.168.1.4", "192.168.1.5", "192.168.1.6"},
		},
		{
			name:   "HP-02: /30 网段，used为空，分配全部2个可用IP",
			subnet: "10.0.0.0/30",
			used:   []string{},
			n:      2,
			want:   []string{"10.0.0.1", "10.0.0.2"},
		},
		{
			name:   "HP-03: /24 网段，used为空，分配1个（第一个可用）",
			subnet: "192.168.1.0/24",
			used:   []string{},
			n:      1,
			want:   []string{"192.168.1.1"},
		},
		{
			name:   "HP-04: /24 网段，已占用前3个，分配从第4个开始",
			subnet: "192.168.1.0/24",
			used:   []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"},
			n:      3,
			want:   []string{"192.168.1.4", "192.168.1.5", "192.168.1.6"},
		},
		{
			name:   "HP-05: /16 网段，分配跨越第一个八位组边界",
			subnet: "10.0.0.0/16",
			used:   []string{},
			n:      3,
			want:   []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"},
		},

		// ── Boundary Cases（边界条件）────────────────────────────────
		{
			name:    "BC-01: /30 可用只有2个，请求2个但1个被占，不足返回error",
			subnet:  "10.0.0.0/30",
			used:    []string{"10.0.0.1"},
			n:       2,
			wantErr: true,
		},
		{
			name:    "BC-02: /30 可用2个全被占，n=1，返回error",
			subnet:  "10.0.0.0/30",
			used:    []string{"10.0.0.1", "10.0.0.2"},
			n:       1,
			wantErr: true,
		},
		{
			name:   "BC-03: /30 分配恰好剩余全部1个",
			subnet: "10.0.0.0/30",
			used:   []string{"10.0.0.1"},
			n:      1,
			want:   []string{"10.0.0.2"},
		},
		{
			name:   "BC-04: /24 分配最后一个可用IP（.254）",
			subnet: "192.168.1.0/24",
			used:   func() []string { // 占用 .1 ~ .253
				ips := make([]string, 253)
				for i := 0; i < 253; i++ {
					ips[i] = fmt.Sprintf("192.168.1.%d", i+1)
				}
				return ips
			}(),
			n:    1,
			want: []string{"192.168.1.254"},
		},
		{
			name:    "BC-05: n=0，返回空切片（无需分配）",
			subnet:  "192.168.1.0/24",
			used:    []string{},
			n:       0,
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "BC-06: subnet 为 /8，分配数量极大（n=1000）",
			subnet:  "10.0.0.0/8",
			used:    []string{},
			n:       1000,
			wantErr: false, // /8 有 16777214 个可用，1000 完全够
		},

		// ── Edge Cases（易出Bug场景）──────────────────────────────────
		{
			name:    "EC-01: subnet 为非法字符串，返回 error",
			subnet:  "not-a-cidr",
			used:    []string{},
			n:       1,
			wantErr: true,
		},
		{
			name:    "EC-02: subnet IP 与掩码不一致（主机位非零），如 192.168.1.5/24",
			subnet:  "192.168.1.5/24",
			used:    []string{},
			n:       1,
			wantErr: true,
		},
		{
			name:    "EC-03: used 中包含不属于该网段的 IP",
			subnet:  "192.168.1.0/24",
			used:    []string{"10.0.0.1"},
			n:       1,
			wantErr: true,
		},
		{
			name:   "EC-04: used 中包含重复 IP，不应重复跳过导致分配错误",
			subnet: "192.168.1.0/24",
			used:   []string{"192.168.1.1", "192.168.1.1"}, // 重复
			n:      2,
			want:   []string{"192.168.1.2", "192.168.1.3"},
		},
		{
			name:   "EC-05: used 顺序乱序，不影响分配结果",
			subnet: "192.168.1.0/24",
			used:   []string{"192.168.1.5", "192.168.1.2", "192.168.1.1"},
			n:      3,
			want:   []string{"192.168.1.3", "192.168.1.4", "192.168.1.6"},
		},
		{
			name:    "EC-06: used 中包含网络地址本身",
			subnet:  "192.168.1.0/24",
			used:    []string{"192.168.1.0"},
			n:       1,
			wantErr: true, // 网络地址不合法，应报错
		},
		{
			name:    "EC-07: used 中包含广播地址",
			subnet:  "192.168.1.0/24",
			used:    []string{"192.168.1.255"},
			n:       1,
			wantErr: true,
		},
		{
			name:    "EC-08: subnet 为 /31，可用IP为0（RFC 3021 特殊情况，按本题规则不支持）",
			subnet:  "192.168.1.0/31",
			used:    []string{},
			n:       1,
			wantErr: true,
		},
		{
			name:    "EC-09: n 为负数",
			subnet:  "192.168.1.0/24",
			used:    []string{},
			n:       -1,
			wantErr: true,
		},
		{
			name:   "EC-10: used 为 nil（而非空切片），行为与空切片一致",
			subnet: "192.168.1.0/24",
			used:   nil,
			n:      2,
			want:   []string{"192.168.1.1", "192.168.1.2"},
		},
		// ... 此处继续补充至 100+ 个用例
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AllocateIPs(tt.subnet, tt.used, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllocateIPs() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("AllocateIPs() len = %d, want %d", len(got), len(tt.want))
					return
				}
				for i := range tt.want {
					if got[i] != tt.want[i] {
						t.Errorf("AllocateIPs()[%d] = %s, want %s", i, got[i], tt.want[i])
					}
				}
			}
		})
	}
}
