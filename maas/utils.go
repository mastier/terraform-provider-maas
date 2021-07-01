package maas

import (
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/gocty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/ionutbalutoiu/gomaasclient/client"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

func base64Encode(data []byte) string {
	if isBase64Encoded(data) {
		return string(data)
	}

	return base64.StdEncoding.EncodeToString(data)
}

func isBase64Encoded(data []byte) bool {
	_, err := base64.StdEncoding.DecodeString(string(data))
	return err == nil
}

func convertToStringSlice(field interface{}) []string {
	if field == nil {
		return nil
	}
	fieldSlice := field.([]interface{})
	result := make([]string, len(fieldSlice))
	for i, value := range fieldSlice {
		result[i] = value.(string)
	}
	return result
}

func isElementIPAddress(i interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	attr := p[len(p)-1].(cty.IndexStep)
	var index int
	if err := gocty.FromCtyValue(attr.Key, &index); err != nil {
		return diag.FromErr(err)
	}
	ws, es := validation.IsIPAddress(i, fmt.Sprintf("element %v", index))

	for _, w := range ws {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Warning,
			Summary:       w,
			AttributePath: p,
		})
	}
	for _, e := range es {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       e.Error(),
			AttributePath: p,
		})
	}
	return diags
}

func getNetworkInterface(client *client.Client, machineSystemID string, identifier string) (*entity.NetworkInterface, error) {
	networkInterfaces, err := client.NetworkInterfaces.Get(machineSystemID)
	if err != nil {
		return nil, err
	}
	for _, n := range networkInterfaces {
		if n.MACAddress == identifier || n.Name == identifier || fmt.Sprintf("%v", n.ID) == identifier {
			return &n, nil
		}
	}
	return nil, fmt.Errorf("network interface (%s) was not found on machine (%s)", identifier, machineSystemID)
}

func setTerraformState(d *schema.ResourceData, tfState map[string]interface{}) error {
	if val, ok := tfState["id"]; ok {
		d.SetId(val.(string))
		delete(tfState, "id")
	}
	for k, v := range tfState {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}
	return nil
}
