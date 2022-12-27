# Table: okta_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| object_class | json | X | √ | Determines the Group's profile. | 
| title | string | X | √ | The title of the resource. | 
| name | string | X | √ | Name of the Group. | 
| id | string | X | √ | Unique key for Group. | 
| created | timestamp | X | √ | Timestamp when Group was created. | 
| last_membership_updated | timestamp | X | √ | Timestamp when Group's memberships were last updated. | 
| type | string | X | √ | Determines how a Group's Profile and memberships are managed. Can be one of OKTA_GROUP, APP_GROUP or BUILT_IN. | 
| description | string | X | √ | Description of the Group. | 
| last_updated | timestamp | X | √ | Timestamp when Group's profile was last updated. | 
| profile | json | X | √ | The Group's Profile properties. | 
| group_members | json | X | √ | List of all users that are a member of this Group. | 


