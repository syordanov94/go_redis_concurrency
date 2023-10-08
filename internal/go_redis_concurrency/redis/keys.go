package redis

import "fmt"

func BuildCompanySharesKey(companyId string) string {
	return fmt.Sprintf("shares:%s", companyId)
}
