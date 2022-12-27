# Table: okta_factor

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | Unique key for Group. | 
| user_name | string | X | √ | Unique identifier for the user (username). | 
| last_updated | timestamp | X | √ | The timestamp when the factor was last updated. | 
| provider | string | X | √ | The provider for the factor. | 
| embedded | json | X | √ | The Group's Profile properties. | 
| title | string | X | √ | The title of the resource. | 
| user_id | string | X | √ | Unique key for Group. | 
| factor_type | string | X | √ | Description of the Group. | 
| created | timestamp | X | √ | Timestamp when Group was created. | 
| status | string | X | √ | The current status of the factor. | 
| verify | json | X | √ | List of all users that are a member of this Group. | 


