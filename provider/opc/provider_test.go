package opc

import (
	"os"
	"testing"

	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"opc": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	required := []string{"OPC_USERNAME", "OPC_PASSWORD", "OPC_IDENTITY_DOMAIN", "OPC_ENDPOINT"}
	for _, prop := range required {
		if os.Getenv(prop) == "" {
			t.Fatalf("%s must be set for acceptance test", prop)
		}
	}
}

type OPCResourceState struct {
	*OPCClient
	*terraform.InstanceState
}

func opcResourceCheck(resourceName string, f func(checker *OPCResourceState) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		state := &OPCResourceState{
			OPCClient:     testAccProvider.Meta().(*OPCClient),
			InstanceState: rs.Primary,
		}

		return f(state)
	}
}
