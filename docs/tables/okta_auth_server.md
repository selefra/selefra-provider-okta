# Table: okta_auth_server

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | Unique key for the authorization server. | 
| issuer | string | X | √ | The issuer URI of the authorization server. | 
| issuer_mode | string | X | √ | The issuer mode of the authorization server. | 
| status | string | X | √ | The status of the authorization server. | 
| title | string | X | √ | The title of the resource. | 
| name | string | X | √ | The name for the authorization server. | 
| created | timestamp | X | √ | Timestamp when the authorization server was created. | 
| description | string | X | √ | A human-readable description of the authorization server. | 
| last_updated | timestamp | X | √ | Timestamp when the authorization server was last updated. | 
| audiences | json | X | √ | The audiences of the authorization server. | 
| credentials | json | X | √ | The authorization server credentials. | 


