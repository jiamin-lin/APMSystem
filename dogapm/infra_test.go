package dogapm

import "testing"

func TestInfra_InitInfra(t *testing.T) {
	Infra.InitInfra(InfraDbOption("root:1234567@tcp(134.175.127.240:3308)/ordersvc"),
		InfraRDBOption("134.175.127.240:6380"))
}
