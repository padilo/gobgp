// Code generated by "stringer -type=API_TYPE"; DO NOT EDIT.

package zebra

import "fmt"

const _API_TYPE_name = "INTERFACE_ADDINTERFACE_DELETEINTERFACE_ADDRESS_ADDINTERFACE_ADDRESS_DELETEINTERFACE_UPINTERFACE_DOWNIPV4_ROUTE_ADDIPV4_ROUTE_DELETEIPV6_ROUTE_ADDIPV6_ROUTE_DELETEREDISTRIBUTE_ADDREDISTRIBUTE_DELETEREDISTRIBUTE_DEFAULT_ADDREDISTRIBUTE_DEFAULT_DELETEIPV4_NEXTHOP_LOOKUPIPV6_NEXTHOP_LOOKUPIPV4_IMPORT_LOOKUPIPV6_IMPORT_LOOKUPINTERFACE_RENAMEROUTER_ID_ADDROUTER_ID_DELETEROUTER_ID_UPDATEHELLOIPV4_NEXTHOP_LOOKUP_MRIBVRF_UNREGISTERINTERFACE_LINK_PARAMSNEXTHOP_REGISTERNEXTHOP_UNREGISTERNEXTHOP_UPDATEMESSAGE_MAX"

var _API_TYPE_index = [...]uint16{0, 13, 29, 50, 74, 86, 100, 114, 131, 145, 162, 178, 197, 221, 248, 267, 286, 304, 322, 338, 351, 367, 383, 388, 412, 426, 447, 463, 481, 495, 506}

func (i API_TYPE) String() string {
	i -= 1
	if i >= API_TYPE(len(_API_TYPE_index)-1) {
		return fmt.Sprintf("API_TYPE(%d)", i+1)
	}
	return _API_TYPE_name[_API_TYPE_index[i]:_API_TYPE_index[i+1]]
}
