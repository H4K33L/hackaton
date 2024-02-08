package Monster

import (
	"encoding/json"
	"io/ioutil"
)

func PullMonsters(filepath string) (groupMonster, error){
	monster := groupMonster{}
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return monster, err
	}

	if len(content) > 0 {
		err = json.Unmarshal(content, &monster)
		if err != nil {
			return monster, err
		}
	}
	 return monster, nil
}

func PushMonsters(filepath string, monsters groupMonster) error {
	newContent, err := json.Marshal(monsters)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, newContent, 0644)
	if err != nil {
		return err
	}

	return nil
}