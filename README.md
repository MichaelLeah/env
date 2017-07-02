# Environment File Configuration for GO
A very basic method to keep your sensitive credentials out of version control within your Go applications. Concept based from the [DOTENV repository](https://github.com/vlucas/phpdotenv)

Simply create a `.env` file in the root directory of your project and fill it with any details you need to store as a key pair. `i.e. API_KEY=EXAMPLEKEYCODE` Then simply call the function with the key to return the value.

```go
var API_KEY = env.Get("API_KEY")
```
