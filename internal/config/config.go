package config

import ("encoding/json"
		
	"os"
	 	"path/filepath")
type Config struct {
	Db_url string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}






func getConfigFilePath() (string, error){
	configFileName := ".gatorconfig.json"
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	file := filepath.Join(home, configFileName)
	return file, nil
}

func Read() (Config, error){
	file, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	reader := Config{}
	err = json.Unmarshal(data,&reader)
	if err != nil {
		return Config{}, err
	}
	return reader, nil
}


 func(a *Config) SetUser(user string)error{ 
 	
	
	file, err := getConfigFilePath()
	if err != nil {
		return err
	}

	a.CurrentUserName = user

	jason, errs  := json.Marshal(a)
	if errs != nil {
		return errs
	}
	 err =os.WriteFile(file,jason, 0600)
	 if err != nil {
	 	return err
	 }
	
	return nil
	
}
