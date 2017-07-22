package main

// Wallet 钱包
type Wallet struct {
	uid int
}

// GetMembers 获取成员数量
func (w *Wallet) GetMembers() int {
	return 1
}

// GetCount 获取收益、已提现、冻结
func (w *Wallet) GetCount() (float32, float32, float32) {
	return 1.1, 2.1, 3.1
}
