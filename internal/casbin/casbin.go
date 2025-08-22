// Package casbin 提供Casbin权限系统初始化和管理功能
// 创建者：Done-0
// 创建时间：2025-08-05
package casbin

import (
	"github.com/casbin/casbin/v2"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/model/rbac"

	gormAdapter "github.com/casbin/gorm-adapter/v3"
)

// New 初始化 Casbin 权限管理系统
// 参数：
//
//	config: 配置信息
func New(config *configs.Config) {
	var enforcer *casbin.Enforcer
	var err error

	// 根据适配器类型创建 Enforcer
	if config.CasbinConfig.DBAdapter {
		adapter, err := gormAdapter.NewAdapterByDBWithCustomTable(global.DB, &rbac.Policy{})
		if err != nil {
			global.SysLog.Errorf("failed to create Casbin database adapter: %v", err)
			return
		}

		// 创建 Enforcer
		enforcer, err = casbin.NewEnforcer(config.CasbinConfig.ModelPath, adapter)
		if err != nil {
			global.SysLog.Errorf("failed to create Casbin Enforcer with model %s: %v", config.CasbinConfig.ModelPath, err)
			return
		}

		// 如果数据库中没有策略，从文件同步到数据库
		if policies, _ := enforcer.GetPolicy(); len(policies) == 0 {
			if fileEnforcer, err := casbin.NewEnforcer(config.CasbinConfig.ModelPath, config.CasbinConfig.PolicyPath); err == nil {
				var policyRecords []rbac.Policy

				// 收集普通策略
				if policies, err := fileEnforcer.GetPolicy(); err == nil {
					for _, policy := range policies {
						if len(policy) >= 3 {
							policyRecord := rbac.Policy{
								Ptype: "p",
								V0:    policy[0],
								V1:    policy[1],
								V2:    policy[2],
							}
							// 如果有额外字段（name 和 description），也要保存
							if len(policy) >= 4 {
								policyRecord.V3 = policy[3]
							}
							if len(policy) >= 5 {
								policyRecord.V4 = policy[4]
							}
							policyRecords = append(policyRecords, policyRecord)
						}
					}
				}

				// 收集角色继承关系
				if groupPolicies, err := fileEnforcer.GetGroupingPolicy(); err == nil {
					for _, groupPolicy := range groupPolicies {
						if len(groupPolicy) >= 2 {
							policyRecord := rbac.Policy{
								Ptype: "g",
								V0:    groupPolicy[0],
								V1:    groupPolicy[1],
							}
							policyRecords = append(policyRecords, policyRecord)
						}
					}
				}

				if len(policyRecords) > 0 {
					if err := global.DB.Create(&policyRecords).Error; err != nil {
						global.SysLog.Errorf("Failed to create policy records: %v", err)
					} else {
						enforcer.LoadPolicy()
						global.SysLog.Infof("Casbin policy synced from file to database")
					}
				}
			}
		}
	} else {
		// 使用文件适配器
		enforcer, err = casbin.NewEnforcer(config.CasbinConfig.ModelPath, config.CasbinConfig.PolicyPath)
		if err != nil {
			global.SysLog.Errorf("创建 Casbin Enforcer 失败: %v", err)
			return
		}
	}

	// 启用自动保存
	enforcer.EnableAutoSave(true)
	global.Enforcer = enforcer
	global.SysLog.Infof("Casbin enforcer initialized successfully")
}
