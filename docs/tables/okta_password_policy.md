# Table: okta_password_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_updated | timestamp | X | √ | Timestamp when the Policy was last modified. | 
| priority | int | X | √ | Priority of the Policy. | 
| status | string | X | √ | Status of the Policy: ACTIVE or INACTIVE. | 
| settings | json | X | √ | Settings of the password policy. | 
| title | string | X | √ | The title of the resource. | 
| description | string | X | √ | Description of the Policy. | 
| created | timestamp | X | √ | Timestamp when the Policy was created. | 
| system | bool | X | √ | This is set to true on system policies, which cannot be deleted. | 
| conditions | json | X | √ | Conditions for Policy. | 
| rules | json | X | √ | Each Policy may contain one or more Rules. Rules, like Policies, contain conditions that must be satisfied for the Rule to be applied. | 
| name | string | X | √ | Name of the Policy. | 
| id | string | X | √ | Identifier of the Policy. | 


