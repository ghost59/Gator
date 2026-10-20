This is Gator, a blog agerragator tool. It was built under the guide of boot.dev.
This tool allows you to register users, add feeds, follows feeds.
It requires Postgress and Go in order to run the program.
To install gator, you will have to use go install.

You can set up a config file like using .gatorconfig.json , this will hold the current uset and the URL to the database.
``` json
{
"db_url":"postgres://example:123@localhost:3344",
"current_user_name":"cow"
}
```
Data migrations work when you set up the schema and use goose in order to get the mirgation to go up or down.
Going down drops the tables and going up creates them.
This is how we keep the tables clean and the code can be generated with sqlc to make new queries. 
