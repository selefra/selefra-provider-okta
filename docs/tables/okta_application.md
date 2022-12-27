# Table: okta_application

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | Unique key for app definition. | 
| label | string | X | √ | User-defined display name for app. | 
| status | string | X | √ | Current status of app. Valid values are ACTIVE or INACTIVE. | 
| visibility | json | X | √ | Visibility settings for app. | 
| title | string | X | √ | The title of the resource. | 
| id | string | X | √ | Unique key for app. | 
| created | timestamp | X | √ | Timestamp when user was created. | 
| last_updated | timestamp | X | √ | Timestamp when app was last updated. | 
| sign_on_mode | string | X | √ | Authentication mode of app. Can be one of AUTO_LOGIN, BASIC_AUTH, BOOKMARK, BROWSER_PLUGIN, Custom, OPENID_CONNECT, SAML_1_1, SAML_2_0, SECURE_PASSWORD_STORE and WS_FEDERATION. | 
| settings | json | X | √ | Settings for app. | 
| credentials | json | X | √ | Credentials for the specified signOnMode. | 
| accessibility | json | X | √ | Access settings for app. | 


