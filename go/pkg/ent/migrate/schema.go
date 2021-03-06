// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GooseDbVersionColumns holds the columns for the "goose_db_version" table.
	GooseDbVersionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "version_id", Type: field.TypeInt},
		{Name: "is_applied", Type: field.TypeBool},
		{Name: "tstamp", Type: field.TypeTime, Nullable: true},
	}
	// GooseDbVersionTable holds the schema information for the "goose_db_version" table.
	GooseDbVersionTable = &schema.Table{
		Name:       "goose_db_version",
		Columns:    GooseDbVersionColumns,
		PrimaryKey: []*schema.Column{GooseDbVersionColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "flags", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserBrewRecipesColumns holds the columns for the "user_brew_recipes" table.
	UserBrewRecipesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "status", Type: field.TypeInt32},
		{Name: "coffee_bean_weight", Type: field.TypeFloat64},
		{Name: "coffee_bean_grind", Type: field.TypeString},
		{Name: "liquid_weight", Type: field.TypeFloat64},
		{Name: "temperature", Type: field.TypeFloat64},
		{Name: "step_one", Type: field.TypeString},
		{Name: "step_two", Type: field.TypeString},
		{Name: "step_three", Type: field.TypeString},
		{Name: "step_four", Type: field.TypeString},
		{Name: "step_five", Type: field.TypeString},
		{Name: "memo", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeInt32, Nullable: true},
		{Name: "user_coffee_bean_id", Type: field.TypeInt32, Nullable: true},
	}
	// UserBrewRecipesTable holds the schema information for the "user_brew_recipes" table.
	UserBrewRecipesTable = &schema.Table{
		Name:       "user_brew_recipes",
		Columns:    UserBrewRecipesColumns,
		PrimaryKey: []*schema.Column{UserBrewRecipesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_brew_recipes_users_user_brew_recipes",
				Columns:    []*schema.Column{UserBrewRecipesColumns[15]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_brew_recipes_user_coffee_beans_user_brew_recipes",
				Columns:    []*schema.Column{UserBrewRecipesColumns[16]},
				RefColumns: []*schema.Column{UserCoffeeBeansColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserCoffeeBeansColumns holds the columns for the "user_coffee_beans" table.
	UserCoffeeBeansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "status", Type: field.TypeInt32},
		{Name: "name", Type: field.TypeString},
		{Name: "farm_name", Type: field.TypeString, Nullable: true},
		{Name: "country", Type: field.TypeString, Nullable: true},
		{Name: "roast_degree", Type: field.TypeString},
		{Name: "roasted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt32, Nullable: true},
	}
	// UserCoffeeBeansTable holds the schema information for the "user_coffee_beans" table.
	UserCoffeeBeansTable = &schema.Table{
		Name:       "user_coffee_beans",
		Columns:    UserCoffeeBeansColumns,
		PrimaryKey: []*schema.Column{UserCoffeeBeansColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_coffee_beans_users_user_coffee_beans",
				Columns:    []*schema.Column{UserCoffeeBeansColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GooseDbVersionTable,
		UsersTable,
		UserBrewRecipesTable,
		UserCoffeeBeansTable,
	}
)

func init() {
	GooseDbVersionTable.Annotation = &entsql.Annotation{
		Table: "goose_db_version",
	}
	UserBrewRecipesTable.ForeignKeys[0].RefTable = UsersTable
	UserBrewRecipesTable.ForeignKeys[1].RefTable = UserCoffeeBeansTable
	UserCoffeeBeansTable.ForeignKeys[0].RefTable = UsersTable
}
