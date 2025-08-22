package impl

// PluginInfo 插件元数据和运行时信息
type PluginInfo struct {
	// 基本信息
	ID          string `json:"id"`                    // 插件唯一标识
	Name        string `json:"name"`                  // 插件显示名称
	Version     string `json:"version"`               // 插件版本号
	Author      string `json:"author"`                // 插件作者
	Description string `json:"description,omitempty"` // 插件描述
	Repository  string `json:"repository"`            // 插件仓库地址
	Binary      string `json:"binary"`                // 二进制文件路径
	Type        string `json:"type"`                  // 插件类型（provider/filter/handler/notifier）

	// 配置信息
	AutoStart    bool  `json:"auto_start"`              // 是否自动启动
	StartTimeout int64 `json:"start_timeout,omitempty"` // 启动超时(毫秒)
	MinPort      uint  `json:"min_port,omitempty"`      // 最小端口
	MaxPort      uint  `json:"max_port,omitempty"`      // 最大端口
	AutoMTLS     bool  `json:"auto_mtls,omitempty"`     // 是否启用自动 MTLS
	Managed      bool  `json:"managed,omitempty"`       // 是否为托管模式

	// 运行时信息
	Status            string `json:"status"`                       // 当前状态
	StartedAt         int64  `json:"started_at,omitempty"`         // 启动时间戳
	ProcessID         string `json:"process_id,omitempty"`         // 进程标识
	Protocol          string `json:"protocol,omitempty"`           // 通信协议
	IsExited          bool   `json:"is_exited,omitempty"`          // 是否已退出
	NegotiatedVersion int    `json:"negotiated_version,omitempty"` // 协商的协议版本
	ProcessPID        int    `json:"process_pid,omitempty"`        // 系统进程 PID
	ProtocolVersion   int    `json:"protocol_version,omitempty"`   // 协议版本
	NetworkAddr       string `json:"network_addr,omitempty"`       // 网络地址
}
