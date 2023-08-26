# go-todo-api

Testing: (always include `https://` in URL)
- `curl -X POST **URL**/auth/signup -H "Content-Type: application/json" --data '{"Name": "..", "Email": "..@..", "Password": ".."}'`
- `curl -X POST **URL**/auth/login -H "Content-Type: application/json" --data '{"Email": "..@..", "Password": ".."}'`
- `curl -i **URL**/todo/list -H "Authorization: Bearer **TOKEN**"`
- `curl -i **URL**/todo/1 -H "Authorization: Bearer **TOKEN**"`
- `curl -X POST **URL**/todo/create -H "Authorization: Bearer **TOKEN**" -H "Content-Type: application/json" --data '{"Task": ".."}'`
- `curl -X PATCH **URL**/todo/1 -H "Authorization: Bearer **TOKEN**" -H "Content-Type: application/json" --data '{"Task": ".."}'`
- `curl -X PATCH **URL**/todo/1/check -H "Authorization: Bearer **TOKEN**" -H "Content-Type: application/json" --data '{"Completed": true}'`
- `curl -X DELETE **URL**/todo/1 -H "Authorization: Bearer **TOKEN**"`

