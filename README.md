# gormLearning
gorm 练习项目

Usage

1. run the sql in schema.sql to create the database
2. fill the conf/config-example.yaml and rename it config.yaml
3. `go run main.go` and input `help` to check the manual

Structure

When the project running, It first conducts some common things like load the config; link to the database .etc
and then prepare the api list by register all `Service` and then complete the api list, link all command to a specific function
After the initial phase. It can take command and call the function by reflection

All sub-service interface were composed in the `Service` struct, and the `NewService()` from all sub-services' imp package
return an implementation

All models locate at the folder _model_