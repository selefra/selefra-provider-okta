# Table: okta_users

## Primary Keys

```
id
```

## Columns

| Column Name                           | Data Type | Uniq | Nullable | Description            |
| ------------------------------------- | --------- |------| -------- | ---------------------- |
| selefra_id                            | String    | √    | √        | primary keys value md5 |
| id                                    | String    | √    | √        |                        |
| last_login                            | Timestamp | X    | √        |                        |
| last_updated                          | Timestamp | X    | √        |                        |
| password_changed                      | Timestamp | X    | √        |                        |
| profile                               | JSON      | X    | √        |                        |
| status                                | String    | X    | √        |                        |
| status_changed                        | Timestamp | X    | √        |                        |
| transitioning_to_status               | String    | X    | √        |                        |
| type_created                          | Timestamp | X    | √        |                        |
| type_created_by                       | String    | X    | √        |                        |
| type_default                          | Bool      | X    | √        |                        |
| type_description                      | String    | X    | √        |                        |
| type_display_name                     | String    | X    | √        |                        |
| type_id                               | String    | X    | √        |                        |
| type_last_updated                     | Timestamp | X    | √        |                        |
| type_last_updated_by                  | String    | X    | √        |                        |
| type_name                             | String    | X    | √        |                        |
| activated                             | Timestamp | X    | √        |                        |
| created                               | Timestamp | X    | √        |                        |
| credentials_password_hash_algorithm   | String    | X    | √        |                        |
| credentials_password_hash_salt        | String    | X    | √        |                        |
| credentials_password_hash_salt_order  | String    | X    | √        |                        |
| credentials_password_hash_value       | String    | X    | √        |                        |
| credentials_password_hash_work_factor | Int       | X    | √        |                        |
| credentials_password_hook_type        | String    | X    | √        |                        |
| credentials_password_value            | String    | X    | √        |                        |
| credentials_provider_name             | String    | X    | √        |                        |
| credentials_provider_type             | String    | X    | √        |                        |
| credentials_recovery_question_answer  | String    | X    | √        |                        |
| credentials_recovery_question         | String    | X    | √        |                        |