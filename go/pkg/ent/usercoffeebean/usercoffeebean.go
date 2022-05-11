// Code generated by entc, DO NOT EDIT.

package usercoffeebean

const (
	// Label holds the string label denoting the usercoffeebean type in the database.
	Label = "user_coffee_bean"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFarmName holds the string denoting the farm_name field in the database.
	FieldFarmName = "farm_name"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldRoastDegree holds the string denoting the roast_degree field in the database.
	FieldRoastDegree = "roast_degree"
	// FieldRoastedAt holds the string denoting the roasted_at field in the database.
	FieldRoastedAt = "roasted_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeUserDripRecipes holds the string denoting the user_drip_recipes edge name in mutations.
	EdgeUserDripRecipes = "user_drip_recipes"
	// Table holds the table name of the usercoffeebean in the database.
	Table = "user_coffee_beans"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_coffee_beans"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// UserDripRecipesTable is the table that holds the user_drip_recipes relation/edge.
	UserDripRecipesTable = "user_drip_recipes"
	// UserDripRecipesInverseTable is the table name for the UserDripRecipe entity.
	// It exists in this package in order to avoid circular dependency with the "userdriprecipe" package.
	UserDripRecipesInverseTable = "user_drip_recipes"
	// UserDripRecipesColumn is the table column denoting the user_drip_recipes relation/edge.
	UserDripRecipesColumn = "coffee_bean_id"
)

// Columns holds all SQL columns for usercoffeebean fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldUserID,
	FieldName,
	FieldFarmName,
	FieldCountry,
	FieldRoastDegree,
	FieldRoastedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
