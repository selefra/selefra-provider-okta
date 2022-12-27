# Table: okta_trusted_origin

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | The title of the resource. | 
| name | string | X | √ | The name of the trusted origin. | 
| created_by | string | X | √ | The ID of the user who created the trusted origin. | 
| origin | string | X | √ | The origin of the trusted origin. | 
| status | string | X | √ | Current status of the trusted origin. Valid values are 'ACTIVE' or 'INACTIVE'. | 
| scopes | json | X | √ | The scopes for the trusted origin. Valid values are 'CORS' or 'REDIRECT'. | 
| id | string | X | √ | A unique key for the trusted origin. | 
| created | timestamp | X | √ | The timestamp when the trusted origin was created. | 
| last_updated | timestamp | X | √ | The timestamp when the trusted origin was last updated. | 
| last_updated_by | string | X | √ | The ID of the user who last updated the trusted origin. | 


