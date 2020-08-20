# Riot API wrapper #
This simple wrapper provides a Go interface for Riot API.

## Config ##
Create a _.env_ file in the root directory of the project and set an ```api_key``` variable like this:
```html
api_key=<your_key>
```

## Usage ##
```go
import "riot-API-client/riot"
client, err := riot.New()
```

## API documentation ##
https://developer.riotgames.com/
