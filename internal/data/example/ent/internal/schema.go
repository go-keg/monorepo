// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/go-keg/monorepo/internal/data/example/ent/schema\",\"Package\":\"github.com/go-keg/monorepo/internal/data/example/ent\",\"Schemas\":[{\"name\":\"Account\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"nickname\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"sensitive\":true}]},{\"name\":\"OperationLog\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"field\":\"user_id\",\"ref_name\":\"operation_logs\",\"unique\":true,\"inverse\":true,\"required\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"created_at\",\"Skip\":48}}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"updated_at\",\"Skip\":48}}},{\"name\":\"user_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"type\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"操作类型\"},{\"name\":\"context\",\"type\":{\"Type\":3,\"Ident\":\"map[string]interface {}\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"map[string]interface {}\",\"Kind\":21,\"PkgPath\":\"\",\"Methods\":{}}},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"Type\":\"Map\"}}}],\"indexes\":[{\"fields\":[\"created_at\"]}],\"annotations\":{\"Comment\":{\"Text\":\"操作日志\"},\"EntGQL\":{\"MutationInputs\":[{\"IsCreate\":true},{}],\"QueryField\":{\"Description\":\"操作日志\"},\"RelayConnection\":true},\"EntSQL\":{\"with_comments\":true}}},{\"name\":\"Permission\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\",\"ref_name\":\"permissions\",\"inverse\":true,\"annotations\":{\"EntGQL\":{\"Skip\":63}}},{\"name\":\"parent\",\"type\":\"Permission\",\"field\":\"parent_id\",\"ref\":{\"name\":\"children\",\"type\":\"Permission\"},\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"created_at\",\"Skip\":48}}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"updated_at\",\"Skip\":48}}},{\"name\":\"parent_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"Skip\":32}}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"key\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"type\",\"type\":{\"Type\":6,\"Ident\":\"permission.Type\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"enums\":[{\"N\":\"menu\",\"V\":\"menu\"},{\"N\":\"page\",\"V\":\"page\"},{\"N\":\"element\",\"V\":\"element\"}],\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"path\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"desc\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"sort\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1000,\"default_kind\":2,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"attrs\",\"type\":{\"Type\":3,\"Ident\":\"map[string]interface {}\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"map[string]interface {}\",\"Kind\":21,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"created_at\"]}],\"annotations\":{\"Comment\":{\"Text\":\"权限\"},\"EntGQL\":{\"MutationInputs\":[{\"IsCreate\":true},{}],\"QueryField\":{\"Description\":\"权限\"},\"RelayConnection\":true},\"EntSQL\":{\"with_comments\":true}}},{\"name\":\"Role\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"permissions\",\"type\":\"Permission\"},{\"name\":\"users\",\"type\":\"User\",\"ref_name\":\"roles\",\"inverse\":true,\"annotations\":{\"EntGQL\":{\"Skip\":63}}}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"created_at\",\"Skip\":48}}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"updated_at\",\"Skip\":48}}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"sort\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1000,\"default_kind\":2,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"created_at\"]}],\"annotations\":{\"Comment\":{\"Text\":\"角色\"},\"EntGQL\":{\"MutationInputs\":[{\"IsCreate\":true},{}],\"QueryField\":{\"Description\":\"角色\"},\"RelayConnection\":true},\"EntSQL\":{\"with_comments\":true}}},{\"name\":\"User\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\"},{\"name\":\"operation_logs\",\"type\":\"OperationLog\",\"annotations\":{\"EntGQL\":{\"Skip\":63}}}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"created_at\",\"Skip\":48}}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"updated_at\",\"Skip\":48}}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"annotations\":{\"EntGQL\":{\"Skip\":63}}},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntGQL\":{\"OrderField\":\"EMAIL\"}}},{\"name\":\"nickname\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"avatar\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"sensitive\":true},{\"name\":\"status\",\"type\":{\"Type\":6,\"Ident\":\"user.Status\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"enums\":[{\"N\":\"normal\",\"V\":\"normal\"},{\"N\":\"freeze\",\"V\":\"freeze\"}],\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"enums\":{\"freeze\":\"冻结\",\"normal\":\"正常\"}},\"comment\":\"状态\"},{\"name\":\"is_admin\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"created_at\"]},{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"annotations\":{\"Comment\":{\"Text\":\"用户\"},\"EntGQL\":{\"MutationInputs\":[{\"IsCreate\":true},{}],\"QueryField\":{\"Description\":\"用户\"},\"RelayConnection\":true},\"EntSQL\":{\"with_comments\":true}}}],\"Features\":[\"intercept\",\"schema/snapshot\",\"sql/modifier\",\"sql/execquery\",\"sql/upsert\",\"namedges\"]}"
