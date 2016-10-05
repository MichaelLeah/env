### Environment File Configuration for GO
#A very basic method to keep your sensitive credentials out of version control within your Go applications. Very loosely based of the Laravel technique. 
#Simply create a .env file in your root directory and fill it with any details you need to store, i.e. API_KEY=EXAMPLEKEYCODE
#Then simply call the function with the key to return the value

`env.Env("API_KEY")`