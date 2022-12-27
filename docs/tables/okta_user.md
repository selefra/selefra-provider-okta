# Table: okta_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| self_link | string | X | √ | A self-referential link to this user. | 
| status_changed | timestamp | X | √ | Timestamp when status last changed. | 
| profile | json | X | √ | User profile properties. | 
| type | json | X | √ | User type that determines the schema for the user's profile. | 
| assigned_roles | json | X | √ | List of roles assigned to user. | 
| activated | timestamp | X | √ | Timestamp when transition to ACTIVE status completed. | 
| last_updated | timestamp | X | √ | Timestamp when user was last updated. | 
| transitioning_to_status | string | X | √ | Target status of an in-progress asynchronous status transition. | 
| user_groups | json | X | √ | List of groups of which the user is a member. | 
| login | string | X | √ | Unique identifier for the user (username). | 
| email | string | X | √ | Primary email address of user. | 
| created | timestamp | X | √ | Timestamp when user was created. | 
| password_changed | timestamp | X | √ | Timestamp when password last changed. | 
| status | string | X | √ | Current status of user. Can be one of the STAGED, PROVISIONED, ACTIVE, RECOVERY, LOCKED_OUT, PASSWORD_EXPIRED, SUSPENDED, or DEPROVISIONED. | 
| title | string | X | √ | The title of the resource. | 
| id | string | X | √ | Unique key for user. | 
| last_login | timestamp | X | √ | Timestamp of last login. | 


