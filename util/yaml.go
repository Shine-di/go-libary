/**
 * @author: D-S
 * @date: 2021/1/25 5:42 下午
 */

package util

import (
	"encoding/json"
	"io/ioutil"
)

func ParseYaml(yamlFileAddr string, data interface{}) error {
	yamlFile, err := ioutil.ReadFile(yamlFileAddr)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(yamlFile, data); err != nil {
		return err
	}
	return nil
}
