package disks

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVirtualDisk(t *testing.T) {
	/*
		DeviceId int    `json:"deviceId,omitempty"`
		Id       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
		Size     int64  `json:"size,omitempty"`
		SizeGB   int    `json:"sizeGB,omitempty"`
		SizeKB   int64  `json:"sizeKB,omitempty"`
		SizeMB   int    `json:"sizeMB,omitempty"`
		Valid    bool   `json:"valid,omitempty"`
		PhysicalDisks []PhysicalDisk `json:"physicalDisks,omitempty"`
	*/
	physicaldisk := []byte(`{"deviceId":13,"Id":76,"valid":true, "physicalDisks":[{"deviceId":12,"Id":77,"valid":true},{"deviceId":14,"Id":75,"valid":true}]}`)

	vd := VirtualDisk{}

	err := json.Unmarshal(physicaldisk, &vd)
	//if err == nil {
	if err != nil {
		//t.Fatalf("%v", vd)
		t.Fatalf("%v", err.Error())
	}

	if len(vd.PhysicalDisks) != 2 {
		t.Fatalf("Disk length should have been 2  but was %d", len(vd.PhysicalDisks))
	}

	fmt.Println("Device id")
	fmt.Printf("%d", vd.DeviceId)
	fmt.Println(" ")

}
