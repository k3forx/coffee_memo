// Code generated by entc, DO NOT EDIT.

package coffeebean

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// FarmName applies equality check predicate on the "farm_name" field. It's identical to FarmNameEQ.
func FarmName(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFarmName), v))
	})
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountry), v))
	})
}

// ShopID applies equality check predicate on the "shop_id" field. It's identical to ShopIDEQ.
func ShopID(v int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShopID), v))
	})
}

// RoastedDegree applies equality check predicate on the "roasted_degree" field. It's identical to RoastedDegreeEQ.
func RoastedDegree(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoastedDegree), v))
	})
}

// RoastedAt applies equality check predicate on the "roasted_at" field. It's identical to RoastedAtEQ.
func RoastedAt(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoastedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// FarmNameEQ applies the EQ predicate on the "farm_name" field.
func FarmNameEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFarmName), v))
	})
}

// FarmNameNEQ applies the NEQ predicate on the "farm_name" field.
func FarmNameNEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFarmName), v))
	})
}

// FarmNameIn applies the In predicate on the "farm_name" field.
func FarmNameIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFarmName), v...))
	})
}

// FarmNameNotIn applies the NotIn predicate on the "farm_name" field.
func FarmNameNotIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFarmName), v...))
	})
}

// FarmNameGT applies the GT predicate on the "farm_name" field.
func FarmNameGT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFarmName), v))
	})
}

// FarmNameGTE applies the GTE predicate on the "farm_name" field.
func FarmNameGTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFarmName), v))
	})
}

// FarmNameLT applies the LT predicate on the "farm_name" field.
func FarmNameLT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFarmName), v))
	})
}

// FarmNameLTE applies the LTE predicate on the "farm_name" field.
func FarmNameLTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFarmName), v))
	})
}

// FarmNameContains applies the Contains predicate on the "farm_name" field.
func FarmNameContains(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFarmName), v))
	})
}

// FarmNameHasPrefix applies the HasPrefix predicate on the "farm_name" field.
func FarmNameHasPrefix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFarmName), v))
	})
}

// FarmNameHasSuffix applies the HasSuffix predicate on the "farm_name" field.
func FarmNameHasSuffix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFarmName), v))
	})
}

// FarmNameIsNil applies the IsNil predicate on the "farm_name" field.
func FarmNameIsNil() predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFarmName)))
	})
}

// FarmNameNotNil applies the NotNil predicate on the "farm_name" field.
func FarmNameNotNil() predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFarmName)))
	})
}

// FarmNameEqualFold applies the EqualFold predicate on the "farm_name" field.
func FarmNameEqualFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFarmName), v))
	})
}

// FarmNameContainsFold applies the ContainsFold predicate on the "farm_name" field.
func FarmNameContainsFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFarmName), v))
	})
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountry), v))
	})
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCountry), v))
	})
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCountry), v...))
	})
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCountry), v...))
	})
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCountry), v))
	})
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCountry), v))
	})
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCountry), v))
	})
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCountry), v))
	})
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCountry), v))
	})
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCountry), v))
	})
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCountry), v))
	})
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCountry)))
	})
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCountry)))
	})
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCountry), v))
	})
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCountry), v))
	})
}

// ShopIDEQ applies the EQ predicate on the "shop_id" field.
func ShopIDEQ(v int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShopID), v))
	})
}

// ShopIDNEQ applies the NEQ predicate on the "shop_id" field.
func ShopIDNEQ(v int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShopID), v))
	})
}

// ShopIDIn applies the In predicate on the "shop_id" field.
func ShopIDIn(vs ...int32) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShopID), v...))
	})
}

// ShopIDNotIn applies the NotIn predicate on the "shop_id" field.
func ShopIDNotIn(vs ...int32) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShopID), v...))
	})
}

// ShopIDGT applies the GT predicate on the "shop_id" field.
func ShopIDGT(v int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShopID), v))
	})
}

// ShopIDGTE applies the GTE predicate on the "shop_id" field.
func ShopIDGTE(v int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShopID), v))
	})
}

// ShopIDLT applies the LT predicate on the "shop_id" field.
func ShopIDLT(v int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShopID), v))
	})
}

// ShopIDLTE applies the LTE predicate on the "shop_id" field.
func ShopIDLTE(v int32) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShopID), v))
	})
}

// RoastedDegreeEQ applies the EQ predicate on the "roasted_degree" field.
func RoastedDegreeEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeNEQ applies the NEQ predicate on the "roasted_degree" field.
func RoastedDegreeNEQ(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeIn applies the In predicate on the "roasted_degree" field.
func RoastedDegreeIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoastedDegree), v...))
	})
}

// RoastedDegreeNotIn applies the NotIn predicate on the "roasted_degree" field.
func RoastedDegreeNotIn(vs ...string) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoastedDegree), v...))
	})
}

// RoastedDegreeGT applies the GT predicate on the "roasted_degree" field.
func RoastedDegreeGT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeGTE applies the GTE predicate on the "roasted_degree" field.
func RoastedDegreeGTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeLT applies the LT predicate on the "roasted_degree" field.
func RoastedDegreeLT(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeLTE applies the LTE predicate on the "roasted_degree" field.
func RoastedDegreeLTE(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeContains applies the Contains predicate on the "roasted_degree" field.
func RoastedDegreeContains(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeHasPrefix applies the HasPrefix predicate on the "roasted_degree" field.
func RoastedDegreeHasPrefix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeHasSuffix applies the HasSuffix predicate on the "roasted_degree" field.
func RoastedDegreeHasSuffix(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeEqualFold applies the EqualFold predicate on the "roasted_degree" field.
func RoastedDegreeEqualFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoastedDegree), v))
	})
}

// RoastedDegreeContainsFold applies the ContainsFold predicate on the "roasted_degree" field.
func RoastedDegreeContainsFold(v string) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoastedDegree), v))
	})
}

// RoastedAtEQ applies the EQ predicate on the "roasted_at" field.
func RoastedAtEQ(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoastedAt), v))
	})
}

// RoastedAtNEQ applies the NEQ predicate on the "roasted_at" field.
func RoastedAtNEQ(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoastedAt), v))
	})
}

// RoastedAtIn applies the In predicate on the "roasted_at" field.
func RoastedAtIn(vs ...time.Time) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoastedAt), v...))
	})
}

// RoastedAtNotIn applies the NotIn predicate on the "roasted_at" field.
func RoastedAtNotIn(vs ...time.Time) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoastedAt), v...))
	})
}

// RoastedAtGT applies the GT predicate on the "roasted_at" field.
func RoastedAtGT(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoastedAt), v))
	})
}

// RoastedAtGTE applies the GTE predicate on the "roasted_at" field.
func RoastedAtGTE(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoastedAt), v))
	})
}

// RoastedAtLT applies the LT predicate on the "roasted_at" field.
func RoastedAtLT(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoastedAt), v))
	})
}

// RoastedAtLTE applies the LTE predicate on the "roasted_at" field.
func RoastedAtLTE(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoastedAt), v))
	})
}

// RoastedAtIsNil applies the IsNil predicate on the "roasted_at" field.
func RoastedAtIsNil() predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRoastedAt)))
	})
}

// RoastedAtNotNil applies the NotNil predicate on the "roasted_at" field.
func RoastedAtNotNil() predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRoastedAt)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CoffeeBean {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoffeeBean(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CoffeeBean) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CoffeeBean) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CoffeeBean) predicate.CoffeeBean {
	return predicate.CoffeeBean(func(s *sql.Selector) {
		p(s.Not())
	})
}
