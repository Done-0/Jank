package mapper

import (
	"fmt"

	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/account"
)

// GetAccountByEmail 根据 Email 获取 Account 用户信息
func GetAccountByEmail(email string) (*model.Account, error) {
	if email == "" {
		return nil, fmt.Errorf("邮箱不能为空")
	}

	var account model.Account
	err := global.DB.Where("email = ?", email).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetAccountByUserID 根据用户ID获取账户信息
func GetAccountByUserID(userID int64) (*model.Account, error) {
	if userID <= 0 {
		return nil, fmt.Errorf("无效用户ID: %d", userID)
	}

	var account model.Account
	err := global.DB.Where("user_id = ?", userID).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// CreateAccount 创建新账户
func CreateAccount(account *model.Account) error {
	if account == nil {
		return fmt.Errorf("账户信息不能为空")
	}

	return global.DB.Create(account).Error
}

// UpdateAccount 更新账户信息
func UpdateAccount(account *model.Account) error {
	if account == nil {
		return fmt.Errorf("账户信息不能为空")
	}

	result := global.DB.Save(account)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("账户不存在或未发生更改")
	}

	return nil
}
