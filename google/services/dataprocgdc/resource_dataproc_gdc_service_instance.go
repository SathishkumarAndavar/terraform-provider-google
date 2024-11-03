// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dataprocgdc

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceDataprocGdcServiceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataprocGdcServiceInstanceCreate,
		Read:   resourceDataprocGdcServiceInstanceRead,
		Update: resourceDataprocGdcServiceInstanceUpdate,
		Delete: resourceDataprocGdcServiceInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataprocGdcServiceInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Location of the resource.`,
			},
			"service_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Id of the service instance.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `User-provided human-readable name to be used in user interfaces.`,
			},
			"gdce_cluster": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Gdce cluster information.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gdce_cluster": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Gdce cluster resource id.`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels to associate with this service instance. Labels may be used for filtering and billing tracking. 

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"service_account": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Requested service account to associate with ServiceInstance.`,
			},
			"spark_service_instance_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Spark-specific service instance configuration.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"effective_service_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Effective service account associated with ServiceInstance. This will be the service_account if specified. Otherwise, it will be an automatically created per-resource P4SA that also automatically has Fleet Workload. Identity bindings applied.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Identifier. The name of the service instance.`,
			},
			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `Whether the service instance is currently reconciling. True if the current state of the resource does not match the intended state, and the system is working to reconcile them, whether or not the change was user initiated.`,
			},
			"requested_state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The intended state to which the service instance is reconciling. Possible values:
* 'CREATING'
* 'ACTIVE'
* 'DISCONNECTED'
* 'DELETING'
* 'STOPPING'
* 'STOPPED'
* 'STARTING'
* 'UPDATING'
* 'FAILED'`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state. Possible values:
* 'CREATING'
* 'ACTIVE'
* 'DISCONNECTED'
* 'DELETING'
* 'STOPPING'
* 'STOPPED'
* 'STARTING'
* 'UPDATING'
* 'FAILED'`,
			},
			"state_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A message explaining the current state.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `System generated unique identifier for this service instance, formatted as UUID4.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was most recently updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDataprocGdcServiceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	gdceClusterProp, err := expandDataprocGdcServiceInstanceGdceCluster(d.Get("gdce_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gdce_cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(gdceClusterProp)) && (ok || !reflect.DeepEqual(v, gdceClusterProp)) {
		obj["gdceCluster"] = gdceClusterProp
	}
	displayNameProp, err := expandDataprocGdcServiceInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	sparkServiceInstanceConfigProp, err := expandDataprocGdcServiceInstanceSparkServiceInstanceConfig(d.Get("spark_service_instance_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spark_service_instance_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkServiceInstanceConfigProp)) && (ok || !reflect.DeepEqual(v, sparkServiceInstanceConfigProp)) {
		obj["sparkServiceInstanceConfig"] = sparkServiceInstanceConfigProp
	}
	serviceAccountProp, err := expandDataprocGdcServiceInstanceServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	labelsProp, err := expandDataprocGdcServiceInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances?serviceInstanceId={{service_instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceInstance: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceInstance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating ServiceInstance: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceInstances/{{service_instance_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DataprocGdcOperationWaitTime(
		config, res, project, "Creating ServiceInstance", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServiceInstance: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServiceInstance %q: %#v", d.Id(), res)

	return resourceDataprocGdcServiceInstanceRead(d, meta)
}

func resourceDataprocGdcServiceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances/{{service_instance_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceInstance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataprocGdcServiceInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}

	if err := d.Set("gdce_cluster", flattenDataprocGdcServiceInstanceGdceCluster(res["gdceCluster"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("name", flattenDataprocGdcServiceInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("uid", flattenDataprocGdcServiceInstanceUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("display_name", flattenDataprocGdcServiceInstanceDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("create_time", flattenDataprocGdcServiceInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("update_time", flattenDataprocGdcServiceInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("requested_state", flattenDataprocGdcServiceInstanceRequestedState(res["requestedState"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("state", flattenDataprocGdcServiceInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("reconciling", flattenDataprocGdcServiceInstanceReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("labels", flattenDataprocGdcServiceInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("spark_service_instance_config", flattenDataprocGdcServiceInstanceSparkServiceInstanceConfig(res["sparkServiceInstanceConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("state_message", flattenDataprocGdcServiceInstanceStateMessage(res["stateMessage"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("service_account", flattenDataprocGdcServiceInstanceServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("effective_service_account", flattenDataprocGdcServiceInstanceEffectiveServiceAccount(res["effectiveServiceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("terraform_labels", flattenDataprocGdcServiceInstanceTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}
	if err := d.Set("effective_labels", flattenDataprocGdcServiceInstanceEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceInstance: %s", err)
	}

	return nil
}

func resourceDataprocGdcServiceInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	// Only the root field "labels" and "terraform_labels" are mutable
	return resourceDataprocGdcServiceInstanceRead(d, meta)
}

func resourceDataprocGdcServiceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceInstance: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances/{{service_instance_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ServiceInstance %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ServiceInstance")
	}

	err = DataprocGdcOperationWaitTime(
		config, res, project, "Deleting ServiceInstance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServiceInstance %q: %#v", d.Id(), res)
	return nil
}

func resourceDataprocGdcServiceInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/serviceInstances/(?P<service_instance_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<service_instance_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<service_instance_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceInstances/{{service_instance_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataprocGdcServiceInstanceGdceCluster(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gdce_cluster"] =
		flattenDataprocGdcServiceInstanceGdceClusterGdceCluster(original["gdceCluster"], d, config)
	return []interface{}{transformed}
}
func flattenDataprocGdcServiceInstanceGdceClusterGdceCluster(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceRequestedState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenDataprocGdcServiceInstanceSparkServiceInstanceConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	return []interface{}{transformed}
}

func flattenDataprocGdcServiceInstanceStateMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceEffectiveServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocGdcServiceInstanceTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenDataprocGdcServiceInstanceEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataprocGdcServiceInstanceGdceCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGdceCluster, err := expandDataprocGdcServiceInstanceGdceClusterGdceCluster(original["gdce_cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGdceCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gdceCluster"] = transformedGdceCluster
	}

	return transformed, nil
}

func expandDataprocGdcServiceInstanceGdceClusterGdceCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcServiceInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcServiceInstanceSparkServiceInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcServiceInstanceServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcServiceInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}