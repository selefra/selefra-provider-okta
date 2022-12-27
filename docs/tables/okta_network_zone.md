# Table: okta_network_zone

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created | timestamp | X | √ | Timestamp when the network zone was created. | 
| asns | json | X | √ | Format of each array value: a string representation of an ASN numeric value. | 
| proxies | json | X | √ | IP addresses (range or CIDR form) that are allowed to forward a request from gateway addresses. These proxies are automatically trusted by Threat Insights. These proxies are used to identify the client IP of a request. | 
| name | string | X | √ | Unique name for the zone. | 
| last_updated | timestamp | X | √ | Timestamp when the network zone was last modified. | 
| title | string | X | √ | The title of the resource. | 
| system | bool | X | √ | Indicates if this is a system network zone. | 
| gateways | json | X | √ | IP addresses (range or CIDR form) of the zone. | 
| locations | json | X | √ | The geolocations of the zone. | 
| id | string | X | √ | Identifier of the network zone. | 
| proxy_type | string | X | √ | One of: '' or null (when not specified), Any (meaning any proxy), Tor, NotTorAnonymizer. | 
| status | string | X | √ | Status of the network zone: ACTIVE or INACTIVE. | 
| type | string | X | √ | The type of the network zone. | 
| usage | string | X | √ | Usage of Zone: POLICY, BLOCKLIST. | 


