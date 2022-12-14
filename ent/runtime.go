// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"entgo.io/bug/ent/hoge"
	"entgo.io/bug/ent/hogeadministrator"
	"entgo.io/bug/ent/schema"
	"entgo.io/bug/ulid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	hogeMixin := schema.Hoge{}.Mixin()
	hogeMixinFields0 := hogeMixin[0].Fields()
	_ = hogeMixinFields0
	hogeMixinFields1 := hogeMixin[1].Fields()
	_ = hogeMixinFields1
	hogeFields := schema.Hoge{}.Fields()
	_ = hogeFields
	// hogeDescCreatedAt is the schema descriptor for created_at field.
	hogeDescCreatedAt := hogeMixinFields1[0].Descriptor()
	// hoge.DefaultCreatedAt holds the default value on creation for the created_at field.
	hoge.DefaultCreatedAt = hogeDescCreatedAt.Default.(func() time.Time)
	// hogeDescUpdatedAt is the schema descriptor for updated_at field.
	hogeDescUpdatedAt := hogeMixinFields1[1].Descriptor()
	// hoge.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	hoge.DefaultUpdatedAt = hogeDescUpdatedAt.Default.(func() time.Time)
	// hoge.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	hoge.UpdateDefaultUpdatedAt = hogeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// hogeDescName is the schema descriptor for name field.
	hogeDescName := hogeFields[0].Descriptor()
	// hoge.NameValidator is a validator for the "name" field. It is called by the builders before save.
	hoge.NameValidator = func() func(string) error {
		validators := hogeDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// hogeDescID is the schema descriptor for id field.
	hogeDescID := hogeMixinFields0[0].Descriptor()
	// hoge.DefaultID holds the default value on creation for the id field.
	hoge.DefaultID = hogeDescID.Default.(func() ulid.ID)
	hogeadministratorMixin := schema.HogeAdministrator{}.Mixin()
	hogeadministratorMixinFields0 := hogeadministratorMixin[0].Fields()
	_ = hogeadministratorMixinFields0
	hogeadministratorMixinFields1 := hogeadministratorMixin[1].Fields()
	_ = hogeadministratorMixinFields1
	hogeadministratorFields := schema.HogeAdministrator{}.Fields()
	_ = hogeadministratorFields
	// hogeadministratorDescCreatedAt is the schema descriptor for created_at field.
	hogeadministratorDescCreatedAt := hogeadministratorMixinFields1[0].Descriptor()
	// hogeadministrator.DefaultCreatedAt holds the default value on creation for the created_at field.
	hogeadministrator.DefaultCreatedAt = hogeadministratorDescCreatedAt.Default.(func() time.Time)
	// hogeadministratorDescUpdatedAt is the schema descriptor for updated_at field.
	hogeadministratorDescUpdatedAt := hogeadministratorMixinFields1[1].Descriptor()
	// hogeadministrator.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	hogeadministrator.DefaultUpdatedAt = hogeadministratorDescUpdatedAt.Default.(func() time.Time)
	// hogeadministrator.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	hogeadministrator.UpdateDefaultUpdatedAt = hogeadministratorDescUpdatedAt.UpdateDefault.(func() time.Time)
	// hogeadministratorDescFirstName is the schema descriptor for first_name field.
	hogeadministratorDescFirstName := hogeadministratorFields[0].Descriptor()
	// hogeadministrator.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	hogeadministrator.FirstNameValidator = hogeadministratorDescFirstName.Validators[0].(func(string) error)
	// hogeadministratorDescLastName is the schema descriptor for last_name field.
	hogeadministratorDescLastName := hogeadministratorFields[1].Descriptor()
	// hogeadministrator.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	hogeadministrator.LastNameValidator = hogeadministratorDescLastName.Validators[0].(func(string) error)
	// hogeadministratorDescEmail is the schema descriptor for email field.
	hogeadministratorDescEmail := hogeadministratorFields[2].Descriptor()
	// hogeadministrator.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	hogeadministrator.EmailValidator = func() func(string) error {
		validators := hogeadministratorDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// hogeadministratorDescIsActive is the schema descriptor for is_active field.
	hogeadministratorDescIsActive := hogeadministratorFields[3].Descriptor()
	// hogeadministrator.DefaultIsActive holds the default value on creation for the is_active field.
	hogeadministrator.DefaultIsActive = hogeadministratorDescIsActive.Default.(bool)
	// hogeadministratorDescID is the schema descriptor for id field.
	hogeadministratorDescID := hogeadministratorMixinFields0[0].Descriptor()
	// hogeadministrator.DefaultID holds the default value on creation for the id field.
	hogeadministrator.DefaultID = hogeadministratorDescID.Default.(func() ulid.ID)
}
