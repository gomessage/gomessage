package v3

// 导入必要的包
import (
	"encoding/json"
	"fmt"
	"gomessage/services/hijacking"
	"reflect"
	"strconv"
)

// 定义一个函数来展开JSON数据
func flattenJSON(data map[string]interface{}, prefix string, flattenedData map[string]interface{}) {
	// 循环遍历数据
	for key, value := range data {
		// 创建一个新的键，带有前缀和当前键
		newKey := prefix + key
		// 检查值的类型
		switch reflect.TypeOf(value).Kind() {
		// 如果是一个映射，使用新前缀递归调用函数
		case reflect.Map:
			flattenJSON(value.(map[string]interface{}), newKey+".", flattenedData)
		// 如果是一个切片，循环遍历切片并使用新前缀递归调用函数
		case reflect.Slice:
			for i, v := range value.([]interface{}) {
				flattenJSON(v.(map[string]interface{}), newKey+"."+strconv.Itoa(i)+".", flattenedData)
			}
		// 如果是其他类型，使用新键将其添加到展平的数据中
		default:
			flattenedData[newKey] = value
		}
	}
}

// 定义一个函数来检查值是否为空
func isEmpty(value interface{}) bool {
	// 检查值的类型
	switch reflect.TypeOf(value).Kind() {
	// 如果是一个字符串，检查它是否为空
	case reflect.String:
		return value.(string) == ""
	// 如果是一个切片，检查它是否为空
	case reflect.Slice:
		return len(value.([]interface{})) == 0
	// 如果是一个映射，检查它是否为空
	case reflect.Map:
		return len(value.(map[string]interface{})) == 0
	// 如果是其他类型，返回false
	default:
		return false
	}
}

// 定义一个函数来检查映射是否为空
func isMapEmpty(data map[string]interface{}) bool {
	// 循环遍历数据
	for _, value := range data {
		// 如果任何值不为空，返回false
		if !isEmpty(value) {
			return false
		}
	}
	// 如果所有值都为空，返回true
	return true
}

// 定义一个函数来展平和过滤JSON数据
func flattenAndFilterJSON(data string) (map[string]interface{}, error) {
	// 将JSON数据解组为映射
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return nil, err
	}
	// 展平JSON数据
	flattenedData := make(map[string]interface{})
	flattenJSON(jsonData, "", flattenedData)
	// 过滤空值
	filteredData := make(map[string]interface{})
	for key, value := range flattenedData {
		if !isEmpty(value) {
			filteredData[key] = value
		}
	}
	// 如果过滤后的数据为空，返回错误
	if isMapEmpty(filteredData) {
		return nil, fmt.Errorf("过滤后的数据为空")
	}
	// 返回过滤后的数据
	return filteredData, nil
}

func FlatteningJson(requestByte []byte) {
	flattenedData, err := flattenAndFilterJSON(string(requestByte))
	if err != nil {
		fmt.Println(err)
	} else {
		d, _ := json.MarshalIndent(flattenedData, "", "  ")
		fmt.Println(string(d))
	}

	hijacking.CacheData.KeyValueMap = make(map[string]any)
	for key, value := range flattenedData {
		tmpMap := hijacking.Mappings{
			Key:   key,
			Value: value,
		}
		hijacking.CacheData.KeyValueList = append(hijacking.CacheData.KeyValueList, tmpMap)
		hijacking.CacheData.KeyValueMap[key] = value
	}
}
