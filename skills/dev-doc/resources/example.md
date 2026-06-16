# 命令配置 vlan batch 改造：从规范全集改为按设备实际承载的业务块生成

## 一、现状

`<业务平面vlan id>`（`biz_vlan_id`）取自 IP 地址规范的业务段定义，是**规范全集**：
规范里定义了哪些业务类型的 VLAN 规划，就全部列出来，与单台设备实际承载的业务无关，
所以每台设备渲染出来的 `vlan batch` 一模一样。

而设备实际要建哪些 VLAN，库里是有准确答案的。`plan_business_ips` 一行 = 一个业务块，
`gateway_device_uuids` 指向它的网关设备，`vlan_id` 是这个块的 VLAN。
以 ACS RDB 规划（FB7，248 块）为例：

LF1_A 名下 19 个块（M-LAG 对端 LF1_B 相同）：

```
block_no  network_type  subnet        vlan_id
1         管理平面       21.18.0.0     1000
1         业务平面       21.18.1.0     1200
1         业务平面       21.18.2.0     1201
…(连续递增)…
1         业务平面       21.18.18.0    1217
```

LF14_A 名下只有 1 个块：

```
block_no  network_type  subnet        vlan_id
14        业务平面       21.18.241.0   1200
```

Spine 名下 0 个块（不做任何块的网关）。

对比现状输出与设备实际需要：

```
vlan batch 1000 1200 to 1217 1201 to 1236 1200    ← 现状：每台都渲染规范全集
vlan batch 1000 1200 to 1217                      ← LF1~LF13 实际只需要这些
vlan batch 1200                                   ← LF14 实际只需要这一个
```

问题有三：规范里的 `1201 to 1236` 没有任何一台设备承载，纯冗余；设备间本有差异
（LF14 只该建 1200），现在全被抹平；多建的 VLAN 都是要人工核对清理的脏配置。

## 二、目标

每台设备的 `vlan batch` 只含自己名下业务块的 VLAN：按 `gateway_device_uuids`
收集该设备所有块的 `vlan_id`，去重、升序、连续段合并为 `X to Y`。

- 管理平面块的 VLAN 进 `<管理平面vlan id>`，业务平面块的进 `<业务平面vlan id>`；
  设备没挂对应块则为空（Spine 两者皆空）。
- `vlan_id = 0`（无 VLAN 形态）跳过。
- 模板语法不变；规范全集逻辑 `buildVlanTexts` 保留，仅作管理平面的回退。

预期结果：LF1~LF13 渲染 `vlan batch 1000 1200 to 1217`，LF14 渲染
`vlan batch 1200`，与各自实际承载完全一致。

## 三、开发计划

1. **VLAN 集合 → 渲染文本的纯函数**

    - **改动**

        - `internal/biz/plan_to_config.go` 新增 `vlanBatchText`

    - **验收标准**

        - 输入 `[1212, 1000, 1211, …, 1201, 0, 1201]`（乱序、含 0、含重复），输出 `1000 1201 to 1212`

        - 空集或全 0 输入，输出空串

2. **装配点改为按台计算**

    - **改动**

        - `plan_to_config.go` 修改 `buildBase`：按 `bipsByDevice` 把设备名下的块分管理／业务两组取 VLAN

        - 删除 builder 的 `bizVlanText` 字段及其加载赋值

    - **验收标准**

        - ACS RDB 规划导出，`LF1_A.cfg` 的该行为 `vlan batch 1000 1200 to 1217`

        - `LF14_A.cfg` 的该行为 `vlan batch 1200`

        - Spine 的配置不含任何业务 VLAN

        - 设备无管理平面块而规范配了管理 VLAN 时，`mgr_vlan_id` 取规范值（回退生效）

3. **测试与交付**

    - **改动**

        - `plan_to_config_test.go` 新增 `vlanBatchText` 与按台装配的用例

    - **验收标准**

        - `go build ./... && go test ./...` 全绿，提交推送到 `refactor/plan-data-model`

        - 部署新镜像后 ACS RDB 重导，三类设备的渲染结果与第 2 步标准一致
