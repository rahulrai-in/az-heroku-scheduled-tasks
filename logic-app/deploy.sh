az group create -l australiaeast -n picletter-rg
az deployment group create --resource-group picletter-rg --template-file picletter-app.definition.json --parameters heroku_token= token <heroku >workflows_picletter_app_name=picletter-logic-app
