# Table: okta_signon_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | Name of the Policy. | 
| description | string | X | √ | Description of the Policy. | 
| status | string | X | √ | Status of the Policy: ACTIVE or INACTIVE. | 
| system | bool | X | √ | This is set to true on system policies, which cannot be deleted. | 
| conditions | json | X | √ | Conditions for Policy. | 
| id | string | X | √ | Identifier of the Policy. | 
| created | timestamp | X | √ | Timestamp when the Policy was created. | 
| last_updated | timestamp | X | √ | Timestamp when the Policy was last modified. | 
| priority | int | X | √ | Priority of the Policy. | 
| rules | json | X | √ | Each Policy may contain one or more Rules. Rules, like Policies, contain conditions that must be satisfied for the Rule to be applied. | 
| title | string | X | √ | The title of the resource. | 


