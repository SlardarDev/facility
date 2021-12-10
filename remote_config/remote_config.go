package remote_config

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"sync"
	"time"

	"github.com/SlardarDev/facility/conv"
	"github.com/SlardarDev/facility/fjson"
	"github.com/SlardarDev/facility/logs"
)

type LockableStruct interface {
	RLock()
	RUnlock()
	Lock()
	Unlock()
}

var (
	ErrTypeOfModel = errors.New("type of model must be a ptr of struct")
	ErrNilModel    = errors.New("value of model must not be nil")
)

var (
	configMap     map[string]string = make(map[string]string)
	configMapLock sync.RWMutex
)

type RemoteConfigClient interface {
	Get(ctx context.Context, key string) (string, error)
}

// rwlock为保护修改用，业务访问参数的时候，需要根据rw mutex来访问
func UnmashalRemoteConfig(tccClient RemoteConfigClient, refreshInterval time.Duration, model LockableStruct, tagName, modelName string) error {

	if tagName == "default" {
		panic("tag name can not be default")
	}
	if tagName == "" {
		tagName = "tcc"
	}

	err := internalUnmarshalTccField(tccClient, model, true, tagName, modelName)
	go func() {
		ticker := time.NewTicker(refreshInterval)
		for range ticker.C {
			_ = internalUnmarshalTccField(tccClient, model, false, tagName, modelName)
		}
	}()
	return err
}

func internalUnmarshalTccField(tccClient RemoteConfigClient, model LockableStruct, isFirst bool, tagName, modelName string) error {
	ptrTypeOfModel := reflect.TypeOf(model)
	if ptrTypeOfModel.Kind() != reflect.Ptr {
		return ErrTypeOfModel
	}
	typeOfModel := ptrTypeOfModel.Elem()
	valueOfModel := reflect.ValueOf(model)
	if valueOfModel.IsNil() {
		return ErrNilModel
	}
	if valueOfModel.IsZero() {
		valueOfModel = reflect.New(typeOfModel)
	}
	valueOfModel = valueOfModel.Elem()
	numField := typeOfModel.NumField()

	for i := 0; i < numField; i++ {
		field := typeOfModel.Field(i)
		vField := valueOfModel.Field(i)

		tccKey := field.Tag.Get(tagName)
		defaultValue, defaultValueOk := field.Tag.Lookup("default")
		if tccKey == "" {
			continue
		}
		tccValue, tccErr := tccClient.Get(context.Background(), tccKey)
		/*如果命中直接跳过吧*/
		configMapLock.RLock()
		if v, ok := configMap[modelName+":"+tccKey]; ok {
			if v == tccValue {
				configMapLock.RUnlock()
				continue
			}
		}
		configMapLock.RUnlock()
		configMapLock.Lock()
		configMap[modelName+":"+tccKey] = tccValue
		configMapLock.Unlock()
		if tccErr != nil {
			logs.Error("get tcc value error for field %s, error: %s", field.Name, tccErr.Error())
			if !defaultValueOk {
				continue
			}
		}
		switch field.Type.Kind() {
		case reflect.Ptr:
			switch field.Type.Elem().Kind() {
			case reflect.Uint64:
				if v, err := conv.Uint64(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Uint64(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Uint32:
				if v, err := conv.Uint32(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Uint32(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Uint16:
				if v, err := conv.Uint16(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Uint16(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Uint8:
				if v, err := conv.Uint8(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Uint8(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Uint:
				if v, err := conv.Uint(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Uint(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Int64:
				if v, err := conv.Int64(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Int64(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Int32:
				if v, err := conv.Int32(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Int32(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Int16:
				if v, err := conv.Int16(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Int16(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Int8:
				if v, err := conv.Int8(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Int8(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Int:
				if v, err := conv.Int(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Int(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.String:
				if tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(tccValue))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if isFirst {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(defaultValue))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				}
			case reflect.Bool:
				if v, err := conv.Bool(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Bool(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Float64:
				if v, err := conv.Float64(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Float64(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Float32:
				if v, err := conv.Float32(tccValue); err == nil && tccErr == nil {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(v))
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				} else if defaultValueOk && isFirst {
					if v, err := conv.Float32(defaultValue); err == nil {
						nField := reflect.New(vField.Type().Elem())
						nField.Elem().Set(reflect.ValueOf(v))
						model.Lock()
						vField.Set(nField)
						model.Unlock()
					}
				}
			case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
				value := ""
				if tccErr != nil && !isFirst {
					// 出错直接反馈
					continue
				} else if tccErr != nil && isFirst {
					value = defaultValue
				} else {
					value = tccValue
				}
				if pv, err := getStructFieldValueForPtr(field, value); err != nil {
					logs.Error("can not unmarshal tcc field %s , cause value of field is %s", field.Name, value)
					continue
				} else {
					nField := reflect.New(vField.Type().Elem())
					nField.Elem().Set(reflect.ValueOf(pv).Elem())
					model.Lock()
					vField.Set(nField)
					model.Unlock()
				}
			case reflect.Interface:
				panic("can not define field as interface")
			}
		case reflect.Uint64:
			if v, err := conv.Uint64(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Uint64(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Uint32:
			if v, err := conv.Uint32(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Uint32(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Uint16:
			if v, err := conv.Uint16(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Uint16(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Uint8:
			if v, err := conv.Uint8(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Uint8(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Uint:
			if v, err := conv.Uint(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Uint(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Int64:
			if v, err := conv.Int64(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Int64(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Int32:
			if v, err := conv.Int32(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Int32(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Int16:
			if v, err := conv.Int16(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Int16(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Int8:
			if v, err := conv.Int8(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Int8(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Int:
			if v, err := conv.Int(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Int(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.String:
			if tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(tccValue))
				model.Unlock()
			} else if isFirst {
				model.Lock()
				vField.Set(reflect.ValueOf(defaultValue))
				model.Unlock()
			}
		case reflect.Bool:
			if v, err := conv.Bool(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Bool(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Float64:
			if v, err := conv.Float64(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Float64(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Float32:
			if v, err := conv.Float32(tccValue); err == nil && tccErr == nil {
				model.Lock()
				vField.Set(reflect.ValueOf(v))
				model.Unlock()
			} else if defaultValueOk && isFirst {
				if v, err := conv.Float32(defaultValue); err == nil {
					model.Lock()
					vField.Set(reflect.ValueOf(v))
					model.Unlock()
				}
			}
		case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
			value := ""
			if tccErr != nil && !isFirst {
				// 出错直接反馈
				continue
			} else if tccErr != nil && isFirst {
				value = defaultValue
			} else {
				value = tccValue
			}
			if pv, err := getStructFieldValue(field, value); err != nil {
				logs.Error("can not unmarshal tcc field %s , cause value of field is %s", field.Name, value)
				continue
			} else {
				model.Lock()
				vField.Set(reflect.ValueOf(pv).Elem())
				model.Unlock()
			}
		case reflect.Interface:
			panic("can not define field as interface")
		}

	}
	return nil
}

func getStructFieldValueForPtr(field reflect.StructField, tccValue string) (interface{}, error) {
	v := reflect.New(field.Type.Elem())
	pv := v.Interface()
	if m, ok := v.Interface().(json.Unmarshaler); ok {
		err := m.UnmarshalJSON([]byte(tccValue))
		if err != nil {
			return nil, err
		}
		return m, nil
	} else {
		err := fjson.ConfigCompatibleWithStandardLibrary.UnmarshalFromString(tccValue, pv)
		if err != nil {
			return nil, err
		}
		return pv, nil
	}
}

func getStructFieldValue(field reflect.StructField, tccValue string) (interface{}, error) {
	v := reflect.New(field.Type)
	pv := v.Interface()
	if m, ok := v.Interface().(json.Unmarshaler); ok {
		err := m.UnmarshalJSON([]byte(tccValue))
		if err != nil {
			return nil, err
		}
		return m, nil
	} else {
		err := fjson.ConfigCompatibleWithStandardLibrary.UnmarshalFromString(tccValue, pv)
		if err != nil {
			return nil, err
		}
		return pv, nil
	}
}
