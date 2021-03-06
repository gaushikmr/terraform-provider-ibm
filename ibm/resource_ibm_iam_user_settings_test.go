package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMIAMUserSettings_Basic(t *testing.T) {
	var allowedIP string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserSettingsBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserSettingsExists("ibm_iam_user_settings.user_settings", allowedIP),
					resource.TestCheckResourceAttr("ibm_iam_user_settings.user_settings", "allowed_ip_addresses.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMUserSettingsUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_settings.user_settings", "allowed_ip_addresses.#", "4"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserSettingsDestroy(s *terraform.State) error {

	userManagement, err := testAccProvider.Meta().(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	client := userManagement.UserInvite()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_user_settings" {
			continue
		}

		usermail := rs.Primary.ID

		userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		accountID := userDetails.userAccount

		accountv1Client, err := testAccProvider.Meta().(ClientSession).BluemixAcccountv1API()
		if err != nil {
			return err
		}
		accUser, err := accountv1Client.Accounts().FindAccountUserByUserId(accountID, usermail)
		if err != nil {
			return err
		} else if accUser == nil {
			return fmt.Errorf("User %s is not found under current account", usermail)
		}

		iamID := accUser.IbmUniqueId

		UserSetting, UserSettingError := client.GetUserSettings(accountID, iamID)
		if UserSettingError == nil && UserSetting.AllowedIPAddresses != "" {
			return fmt.Errorf("Allowed IP setting still exists: %s", usermail)
		}
	}

	return nil
}

func testAccCheckIBMIAMUserSettingsExists(n string, ip string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No ID is set")
		}

		userManagement, err := testAccProvider.Meta().(ClientSession).UserManagementAPI()
		if err != nil {
			return err
		}

		client := userManagement.UserInvite()

		usermail := rs.Primary.ID

		userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		accountID := userDetails.userAccount

		accountv1Client, err := testAccProvider.Meta().(ClientSession).BluemixAcccountv1API()
		if err != nil {
			return err
		}
		accUser, err := accountv1Client.Accounts().FindAccountUserByUserId(accountID, usermail)
		if err != nil {
			return err
		} else if accUser == nil {
			return fmt.Errorf("User %s is not found under current account", usermail)
		}

		iamID := accUser.IbmUniqueId

		UserSetting, UserSettingError := client.GetUserSettings(accountID, iamID)
		if UserSettingError != nil {
			return fmt.Errorf("ERROR in getting user settings: %s", rs.Primary.ID)
		}

		ip = UserSetting.AllowedIPAddresses
		return nil
	}
}

func testAccCheckIBMIAMUserSettingsBasic() string {
	return fmt.Sprintf(`

		  
	resource "ibm_iam_user_settings" "user_settings" {
		iam_id = "%s"
		allowed_ip_addresses = ["192.168.0.0","192.168.0.1"]
	  }

	`, IAMUser)
}

func testAccCheckIBMIAMUserSettingsUpdate() string {
	return fmt.Sprintf(`

		  
	resource "ibm_iam_user_settings" "user_settings" {
		iam_id = "%s"
		allowed_ip_addresses = ["192.168.0.2","192.168.0.3","192.168.0.4","192.168.0.5"]
	  }

	`, IAMUser)
}
