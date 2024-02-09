package Monster

import (
	"encoding/json"
	"io/ioutil"
)

func PullMonsters(filepath string) (groupMonster, error){
	/*
	The pullMonster func take a string of the path to the json file
	and return a struct fill whith information of the json file.
	---------------------------------------------------------------
	input : a string
	output : a groupMonster struct and posibly error case
	---------------------------------------------------------------
	the function can return error case
	*/

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
	/*
	The pullMonster func take a string of the path to the json file 
	and a groupMonster struct.
	The function convert the groupMonster struct in Json format and
	rewrite the monster.json file
	---------------------------------------------------------------
	input : a string and groupMonster struct
	output : posibly error case
	---------------------------------------------------------------
	the function can return error case
	*/

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