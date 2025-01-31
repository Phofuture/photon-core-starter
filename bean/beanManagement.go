package bean

import (
	"fmt"
	"reflect"
)

var (
	beanMap       = make(map[reflect.Type]reflect.Value)
	autowireArray = make([]interface{}, 0)
)

func RegisterBeanPtr(bean interface{}) {
	reflectedType := reflect.TypeOf(bean)
	reflectedValue := reflect.ValueOf(bean)
	if reflectedType.Kind() != reflect.Ptr {
		panic("bean must be a pointer, bean type" + reflectedType.String())
	}
	beanMap[reflectedType] = reflectedValue
}

func Autowire(component interface{}) {
	autowireArray = append(autowireArray, component)
}

func StartBeanManagement() {
	for _, component := range autowireArray {
		reflectedType := reflect.TypeOf(component)
		reflectedValue := reflect.ValueOf(component)
		if reflectedType.Kind() != reflect.Ptr {
			panic("component must be a pointer, type:" + reflectedType.String())
		}

		for i := 0; i < reflectedType.Elem().NumField(); i++ {
			field := reflectedType.Elem().Field(i)
			fieldValue := reflectedValue.Elem().Field(i)
			tag := field.Tag.Get("autowired")
			if tag == "" {
				continue
			}
			if field.Type.Kind() != reflect.Ptr {
				panic("autowired field must be a pointer, field:" + field.Name)
			}
			if v, ok := beanMap[field.Type]; ok {
				fieldValue.Set(v)
			} else {
				panic(fmt.Sprintf("autowired bean not found, field:%s, type:%s", field.Name, field.Type.String()))
			}
		}
	}
}
