# ChatGPI CLI

This is sample application to use chatgpt api via CLI Application

## First Start

Download Binary file to execute (depend on OS)

execute with command.

```sh
chatgpt {your message}
```

or you can set path for easy to use.

## Configurations

Set Environment variable to configuration.

`GPT_CLI_APIKEY` is ChatGPI API Key that request from https://platform.openai.com/account/api-keys 

*On Windows Powershell*

```ps
$env:GPT_CLI_APIKEY='sk-xxxxxxxxxxxxxxxxxxxxxxxx'
```

*On Unix Terminal*

```sh
GPT_CLI_APIKEY=sk-xxxxxxxxxxxxxxxxxxxxxxxx
```

`GPT_CLI_MODEL` is ChatGPI Model to use (default is gpt-3.5-turbo)

*On Windows Powershell*

```ps
$env:GPT_CLI_MODEL='gpt-3.5-turbo'
```

*On Unix Terminal*

```sh
GPT_CLI_MODEL=gpt-3.5-turbo
```
