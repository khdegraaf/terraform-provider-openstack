package openstack

import "github.com/hashicorp/terraform/helper/schema"

func ResourceComputeInstanceV2Create(d *schema.ResourceData, meta interface{}) error {
	return resourceComputeInstanceV2Create(d, meta)
}

func ResourceComputeInstanceV2Read(d *schema.ResourceData, meta interface{}) error {
	return resourceComputeInstanceV2Read(d, meta)
}

func ResourceComputeInstanceV2Delete(d *schema.ResourceData, meta interface{}) error {
	return resourceComputeInstanceV2Delete(d, meta)
}

func ResourceComputeInstanceV2Update(d *schema.ResourceData, meta interface{}) error {
	return resourceComputeInstanceV2Update(d, meta)
}

func ResourceComputeSchedulerHintsHash(v interface{}) int {
	return resourceComputeSchedulerHintsHash(v)
}

func ResourceComputeInstancePersonalityHash(v interface{}) int {
	return resourceComputeInstancePersonalityHash(v)
}
