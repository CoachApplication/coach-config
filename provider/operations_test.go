package provider_test

import (
	"context"
	"github.com/CoachApplication/config"
	"github.com/CoachApplication/config/provider"
	"reflect"
	"testing"
	"time"
)

/**
 * See the test Backend provider generated in provier_backended_test - we are reusing it
 */

func TestGetOperation_Exec(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	p := testBackendConfigProvider(t)
	gOp := provider.NewGetOperation(p)

	props := gOp.Properties()

	if kp, err := props.Get(config.PROPERTY_ID_KEY); err != nil {
		t.Error("GetOperation returned an error the Key property: ", err.Error())
	} else {

		if err := kp.Set("F"); err != nil {
			t.Error("GetOperation returned an error when the Key property Set() was called: ", err.Error())
		} else {
			ctx, _ := context.WithTimeout(context.Background(), dur)
			res := gOp.Exec(props)
			select {
			case <-res.Finished():
				if !res.Success() {
					t.Error("GetOperation returned a failed result when retrieving a valid Key: ", res.Errors())
				} else if scVal, err := res.Properties().Get(config.PROPERTY_ID_SCOPEDCONFIG); err != nil {
					t.Error("GetOperation result returned an error when retrieving the ScopedConfig Property: ", err.Error())
				} else if sc, good := scVal.Get().(config.ScopedConfig); !good {
					t.Error("GetOperation ScopedConfig Property returned the wrong data type: ", reflect.TypeOf(scVal))
				} else {

					if scNo, err := sc.Get("no"); err == nil {
						t.Error("GetOperation ScopedConfigProperty invalid scope Config retrieval produced no error: ", scNo)
					} else if scNo == nil {
						t.Log("GetOperation ScopeConfig Config for an invalid key was nil")
					} else {
						res := scNo.HasValue()
						select {
						case <-res.Finished():

							if res.Success() {
								t.Error("ScopedConfig indicated that it has a value for an invalid scope")
							}

						case <-ctx.Done():
						}

						t.Error("GetOperation ScopedConfigProperty invalid scope Config indicates that it has a value: ", scNo)
					}
				}
			case <-ctx.Done():
				t.Error("GetOperation Exec result Exec")
			}

		}

		if err := kp.Set("A"); err != nil {
			t.Error("GetOperation returned an error when the Key property Set() was called: ", err.Error())
		} else {
			ctx, _ := context.WithTimeout(context.Background(), dur)
			res := gOp.Exec(props)
			select {
			case <-res.Finished():
				if !res.Success() {
					t.Error("GetOperation returned a failed result when retrieving a valid Key: ", res.Errors())
				} else if scVal, err := res.Properties().Get(config.PROPERTY_ID_SCOPEDCONFIG); err != nil {
					t.Error("GetOperation result returned an error when retrieving the ScopedConfig Property: ", err.Error())
				} else if sc, good := scVal.Get().(config.ScopedConfig); !good {
					t.Error("GetOperation ScopedConfig Property returned the wrong data type: ", reflect.TypeOf(scVal))
				} else {

					if scNo, err := sc.Get("no"); err == nil {
						t.Error("GetOperation ScopedConfigProperty invalid scope Config retrieval produced no error: ", scNo)
					} else if scNo == nil {
						t.Log("GetOperation ScopeConfig Config for an invalid key was nil")
					} else {
						res := scNo.HasValue()
						ctx, _ := context.WithTimeout(context.Background(), dur)
						select {
						case <-res.Finished():

							if res.Success() {
								t.Error("ScopedConfig indicated that it has a value for an invalid scope")
							}

						case <-ctx.Done():
						}

						t.Error("GetOperation ScopedConfigProperty invalid scope Config indicates that it has a value: ", scNo)
					}

					if scA, err := sc.Get("B"); err != nil {
						t.Error("GetOperation ScopedConfigProperty valid scope Config retrieval produced an error: ", scA)
					} else if scA == nil {
						t.Error("GetOperation ScopeConfig Config for a valid key was nil")
					} else {
						res := scA.HasValue()
						ctx, _ := context.WithTimeout(context.Background(), dur)
						select {
						case <-res.Finished():

							if !res.Success() {
								t.Error("ScopedConfig indicated that it has no value for a valid scope")
							}

						case <-ctx.Done():
							t.Error("ScopedConfig Config timed out on .HasValue() check: ", ctx.Err())
						}

						var A string
						res = scA.Get(&A)
						ctx, _ = context.WithTimeout(context.Background(), dur)
						select {
						case <-res.Finished():

							if !res.Success() {
								t.Error("ScopedConfig returned failed result on Get: ", res.Errors())
							} else if A != "AB" {
								t.Error("ScopedConfig returned wrong value on Get")
							}

						case <-ctx.Done():
							t.Error("ScopedConfig Config timed out on .HasValue() check: ", ctx.Err())
						}
					}
				}
			case <-ctx.Done():
				t.Error("GetOperation Exec result Exec")
			}

		}

	}
}

func TestListOperation_Exec(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	p := testBackendConfigProvider(t)
	lOp := provider.NewListOperation(p)

	res := lOp.Exec(lOp.Properties())
	ctx, _ := context.WithTimeout(context.Background(), dur)
	select {
	case <-res.Finished():

		if !res.Success() {
			t.Error("ListOperation returned a failed result: ", res.Errors())
		} else if ksProp, err := res.Properties().Get(config.PROPERTY_ID_KEYS); err != nil {
			t.Error("ListOperation result did not have a keys property: ", err.Error())
		} else if ks, good := ksProp.Get().([]string); !good {
			t.Error("ListOPeration kes property gave invalid data: ", ks)
		} else if len(ks) == 0 {
			t.Error("No keys returned")
		} else if len(ks) != 2 {
			t.Error("Wrong number of keys returned: ", ks)
		} else if !valIsInSlice("A", ks) {
			t.Error("Missing Key: A")
		} else if !valIsInSlice("B", ks) {
			t.Error("Missing Key: B")
		}

	case <-ctx.Done():
		t.Error("ListOperation timed out: ", ctx.Err())
	}
}

// is a value in a slice
func valIsInSlice(val string, slice []string) bool {
	for _, each := range slice {
		if each == val {
			return true
		}
	}
	return false
}
