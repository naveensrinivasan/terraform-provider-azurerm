package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/features"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

func schemaFeatures(supportLegacyTestSuite bool) *pluginsdk.Schema {
	// NOTE: if there's only one nested field these want to be Required (since there's no point
	//       specifying the block otherwise) - however for 2+ they should be optional
	featuresMap := map[string]*pluginsdk.Schema{
		//lintignore:XS003
		"api_management": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"purge_soft_delete_on_destroy": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						Default:  true,
					},

					"recover_soft_deleted": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						Default:  true,
					},
				},
			},
		},

		"cognitive_account": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"purge_soft_delete_on_destroy": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
						Default:  true,
					},
				},
			},
		},

		"key_vault": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"recover_soft_deleted_key_vaults": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
					"purge_soft_delete_on_destroy": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
				},
			},
		},

		"log_analytics_workspace": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"permanently_delete_on_destroy": {
						Type:     pluginsdk.TypeBool,
						Required: true,
					},
				},
			},
		},

		"network": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"relaxed_locking": {
						Type:     pluginsdk.TypeBool,
						Required: true,
					},
				},
			},
		},

		"template_deployment": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"delete_nested_items_during_deletion": {
						Type:     pluginsdk.TypeBool,
						Required: true,
					},
				},
			},
		},

		//lintignore:XS003
		"virtual_machine": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"delete_os_disk_on_deletion": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
					"graceful_shutdown": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
					"skip_shutdown_and_force_delete": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},

		"virtual_machine_scale_set": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"force_delete": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
					"roll_instances_when_required": {
						Type:     pluginsdk.TypeBool,
						Required: true,
					},
					"scale_to_zero_before_deletion": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
				},
			},
		},

		"resource_group": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"prevent_deletion_if_contains_resources": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
				},
			},
		},
	}

	if features.ThreePointOhBeta() {
		f := featuresMap["key_vault"].Elem.(*pluginsdk.Resource)
		// TODO: Add this to 3.0 Upgrade guide
		// `recover_soft_deleted_keys` - (Default: true) when enabled soft-deleted `azurerm_key_vault_key` resources will be restored, instead of creating new ones.
		f.Schema["recover_soft_deleted_keys"] = &pluginsdk.Schema{
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		}

		// TODO: Add this to 3.0 Upgrade guide
		// `purge_soft_deleted_keys_on_destroy` - (Default: true) when enabled soft-deleted `azurerm_key_vault_key` resources will be permanently deleted (e.g purged), when destroyed.
		f.Schema["purge_soft_deleted_keys_on_destroy"] = &pluginsdk.Schema{
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		}

		// TODO: Add this to 3.0 Upgrade guide
		// `recover_soft_deleted_certificates` - (Default: true) when enabled soft-deleted `azurerm_key_vault_certificate` resources will be restored, instead of creating new ones.
		f.Schema["recover_soft_deleted_certificates"] = &pluginsdk.Schema{
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		}

		// TODO: Add this to 3.0 Upgrade guide
		// `purge_soft_deleted_certificates_on_destroy` - (Default: true) when enabled soft-deleted `azurerm_key_vault_certificate` resources will be permanently deleted (e.g purged), when destroyed.
		f.Schema["purge_soft_deleted_certificates_on_destroy"] = &pluginsdk.Schema{
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		}

		// TODO: Add this to 3.0 Upgrade guide
		// `recover_soft_deleted_secrets` - (Default: true) when enabled soft-deleted `azurerm_key_vault_secret` resources will be restored, instead of creating new ones.
		f.Schema["recover_soft_deleted_secrets"] = &pluginsdk.Schema{
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		}

		// TODO: Add this to 3.0 Upgrade guide
		// `purge_soft_deleted_secrets_on_destroy` - (Default: true) when enabled soft-deleted `azurerm_key_vault_secret` resources will be permanently deleted (e.g purged), when destroyed.
		f.Schema["purge_soft_deleted_secrets_on_destroy"] = &pluginsdk.Schema{
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		}
	}

	// this is a temporary hack to enable us to gradually add provider blocks to test configurations
	// rather than doing it as a big-bang and breaking all open PR's
	if supportLegacyTestSuite {
		return &pluginsdk.Schema{
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: featuresMap,
			},
		}
	}

	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: featuresMap,
		},
	}
}

func expandFeatures(input []interface{}) features.UserFeatures {
	// these are the defaults if omitted from the config
	featuresMap := features.Default()

	if len(input) == 0 || input[0] == nil {
		return featuresMap
	}

	val := input[0].(map[string]interface{})

	if raw, ok := val["api_management"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 && items[0] != nil {
			apimRaw := items[0].(map[string]interface{})
			if v, ok := apimRaw["purge_soft_delete_on_destroy"]; ok {
				featuresMap.ApiManagement.PurgeSoftDeleteOnDestroy = v.(bool)
			}
			if v, ok := apimRaw["recover_soft_deleted"]; ok {
				featuresMap.ApiManagement.RecoverSoftDeleted = v.(bool)
			}
		}
	}

	if raw, ok := val["cognitive_account"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 && items[0] != nil {
			cognitiveRaw := items[0].(map[string]interface{})
			if v, ok := cognitiveRaw["purge_soft_delete_on_destroy"]; ok {
				featuresMap.CognitiveAccount.PurgeSoftDeleteOnDestroy = v.(bool)
			}
		}
	}

	if raw, ok := val["key_vault"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 && items[0] != nil {
			keyVaultRaw := items[0].(map[string]interface{})
			if v, ok := keyVaultRaw["purge_soft_delete_on_destroy"]; ok {
				featuresMap.KeyVault.PurgeSoftDeleteOnDestroy = v.(bool)
			}
			if v, ok := keyVaultRaw["recover_soft_deleted_key_vaults"]; ok {
				featuresMap.KeyVault.RecoverSoftDeletedKeyVaults = v.(bool)
			}

			if !features.ThreePointOhBeta() {
				// Inherit Key Vault recovery setting by default. If we're on 3.0 then the code below will overwrite
				// these values as needed.
				featuresMap.KeyVault.RecoverSoftDeletedCerts = featuresMap.KeyVault.RecoverSoftDeletedKeyVaults
				featuresMap.KeyVault.RecoverSoftDeletedSecrets = featuresMap.KeyVault.RecoverSoftDeletedKeyVaults
				featuresMap.KeyVault.RecoverSoftDeletedKeys = featuresMap.KeyVault.RecoverSoftDeletedKeyVaults
				featuresMap.KeyVault.PurgeSoftDeletedKeysOnDestroy = featuresMap.KeyVault.PurgeSoftDeleteOnDestroy
				featuresMap.KeyVault.PurgeSoftDeletedCertsOnDestroy = featuresMap.KeyVault.PurgeSoftDeleteOnDestroy
				featuresMap.KeyVault.PurgeSoftDeletedSecretsOnDestroy = featuresMap.KeyVault.PurgeSoftDeleteOnDestroy
			} else {
				if v, ok := keyVaultRaw["recover_soft_deleted_certificates"]; ok {
					featuresMap.KeyVault.RecoverSoftDeletedCerts = v.(bool)
				}
				if v, ok := keyVaultRaw["purge_soft_deleted_certificates_on_destroy"]; ok {
					featuresMap.KeyVault.PurgeSoftDeletedCertsOnDestroy = v.(bool)
				}
				if v, ok := keyVaultRaw["recover_soft_deleted_secrets"]; ok {
					featuresMap.KeyVault.RecoverSoftDeletedSecrets = v.(bool)
				}
				if v, ok := keyVaultRaw["purge_soft_deleted_secrets_on_destroy"]; ok {
					featuresMap.KeyVault.PurgeSoftDeletedSecretsOnDestroy = v.(bool)
				}
				if v, ok := keyVaultRaw["recover_soft_deleted_keys"]; ok {
					featuresMap.KeyVault.RecoverSoftDeletedKeys = v.(bool)
				}
				if v, ok := keyVaultRaw["purge_soft_deleted_keys_on_destroy"]; ok {
					featuresMap.KeyVault.PurgeSoftDeletedKeysOnDestroy = v.(bool)
				}
			}
		}
	}

	if raw, ok := val["log_analytics_workspace"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 {
			logAnalyticsWorkspaceRaw := items[0].(map[string]interface{})
			if v, ok := logAnalyticsWorkspaceRaw["permanently_delete_on_destroy"]; ok {
				featuresMap.LogAnalyticsWorkspace.PermanentlyDeleteOnDestroy = v.(bool)
			}
		}
	}

	if raw, ok := val["template_deployment"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 {
			templateRaw := items[0].(map[string]interface{})
			if v, ok := templateRaw["delete_nested_items_during_deletion"]; ok {
				featuresMap.TemplateDeployment.DeleteNestedItemsDuringDeletion = v.(bool)
			}
		}
	}

	if raw, ok := val["virtual_machine"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 && items[0] != nil {
			virtualMachinesRaw := items[0].(map[string]interface{})
			if v, ok := virtualMachinesRaw["delete_os_disk_on_deletion"]; ok {
				featuresMap.VirtualMachine.DeleteOSDiskOnDeletion = v.(bool)
			}
			if v, ok := virtualMachinesRaw["graceful_shutdown"]; ok {
				featuresMap.VirtualMachine.GracefulShutdown = v.(bool)
			}
			if v, ok := virtualMachinesRaw["skip_shutdown_and_force_delete"]; ok {
				featuresMap.VirtualMachine.SkipShutdownAndForceDelete = v.(bool)
			}
		}
	}

	if raw, ok := val["virtual_machine_scale_set"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 {
			scaleSetRaw := items[0].(map[string]interface{})
			if v, ok := scaleSetRaw["roll_instances_when_required"]; ok {
				featuresMap.VirtualMachineScaleSet.RollInstancesWhenRequired = v.(bool)
			}
			if v, ok := scaleSetRaw["force_delete"]; ok {
				featuresMap.VirtualMachineScaleSet.ForceDelete = v.(bool)
			}
			if v, ok := scaleSetRaw["scale_to_zero_before_deletion"]; ok {
				featuresMap.VirtualMachineScaleSet.ScaleToZeroOnDelete = v.(bool)
			}
		}
	}

	if raw, ok := val["resource_group"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 {
			resourceGroupRaw := items[0].(map[string]interface{})
			if v, ok := resourceGroupRaw["prevent_deletion_if_contains_resources"]; ok {
				featuresMap.ResourceGroup.PreventDeletionIfContainsResources = v.(bool)
			}
		}
	}

	return featuresMap
}
