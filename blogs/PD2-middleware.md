


# Foundation is Key: Middleware

We've probably heard this word numerous times, but what is it and why are they important? Before you google the word itself first and get cloudy after, I'd like to give you an example.

Say in our golang microservice, we have an endpoint to list all the bad habits of Nathan but only Nathan himself can see it (Typical right? because it's privacy. Let's not talk about the purpose of the endpoint itself because you already think this is ridiculous, because it is). So authorization is needed to verify if the request is sent by Nathan. Easy piece, we can just write the following code and call it a day,

```go
package main

import (
	"http"
...
)

func main() {
	// whatever setup
	mux := http.NewServeMux()
	
	mux. HandleFunc("/v1/list-nathan-bad-habits", HandleListNathanBadHabits())
	
	// whatever after
}

func HandleListNathanBadHabits() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if it's Nathan who sends the request
	
		// whatever logic to list all Nathan's bad habits
	}
}
```

So far so good, right? Well, now that Nathan has hit this endpoint and then learns how many bad habits he has, he decides to improve himself by creating another endpoint to update the list (which is cheating by the way). So now the code becomes

```go
package main

import (
	"http"
...
)

func main() {
	// whatever setup
	mux := http.NewServeMux()
	
	mux.HandleFunc("/v1/list-nathan-bad-habits", HandleListNathanBadHabits())
	mux.HandleFunc("/v1/update-nathan-bad-habits", HandleUpdateNathanBadHabits())
	
	// whatever after
}

func HandleListNathanBadHabits() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if it's Nathan who sends the request
	
		// whatever logic to list all Nathan's bad habits
	}
}

func HandleUpdateNathanBadHabits() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if it's Nathan who sends the request
	
		// whatever logic to update all Nathan's bad habits
	}
}
```

As you may have noticed, the same authorization has been used twice already; and as you add more endpoints using the same authorization, the code duplication keeps growing. So you may think, easy piece, I'll wrap up the auth logic into a function. Not too bad a solution, but that way the handler is not decoupled from the auth logic; i.e., we want the handler to only have the list/update logic and when the auth fails we don't want to enter the handler logic. So here a middleware comes into play,


```go
package main

import (
	"http"
...
)

func main() {
	// whatever setup
	mux := http.NewServeMux()
	
	// now we can use the auth middleware
	mux.Handle("/v1/list-nathan-bad-habits", isAuthed(HandleListNathanBadHabits()))
	mux.Handle("/v1/update-nathan-bad-habits", isAuthed(HandleUpdateNathanBadHabits()))
	
	// whatever after
}

func HandleListNathanBadHabits() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// whatever logic to list all Nathan's bad habits
	}
}

func HandleUpdateNathanBadHabits() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// whatever logic to update all Nathan's bad habits
	}
}

func isAuthed(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(w http.ResponseWriter, r *http.Request) {
		// check if it's Nathan who sends the request
		
		if err { // say you have error from previous steps
			http.Error(w, err.Error(), http.StatusBadRequest)
			
			// In some tutorials you would also see some code below
			// fmt.Fprintf(w, err.Error())
			// the reason why we don't use this approach can be that the http error code is unclear
			return
		}
		
		// serve the request by redirecting to the endpoint handler
		endpoint(w, r)
	}
}
```

(The middleware concept is now adopted in the `cloud-blue` microservice where we want to parse JWT tokens for new asset requests. The PR can be found [here](https://github.com/vendasta/cloud-blue/pull/18))

A pretty simple demo, isn't it? So now back to the beginning, what is middleware and why is it important? According to [this article](https://www.redhat.com/en/topics/middleware/what-is-middleware#:~:text=Middleware%20is%20software%20that%20provides,all%20commonly%20handled%20by%20middleware.), middleware provides common services and capabilities to applications where the common services and capabilities are not provided by the operating systems. In a cloud-native development setting, middleware is an effective way to bring down complexity resulted from discrepancies such as infrastructures, frameworks, operating systems, and languages.

At Vendasta, you probably have seen multiple use cases of middleware in our golang microservice codebases to handle auth and JWT token parsing; additionally, in our Angular projects, you can also find a use case of middleware and that is our router-guard service. Briefly speaking, when you have a form component to update some information, upon clicking a button to redirect to another page without having saved your current form, you may consider adding a router-guard service so that users don't get redirected immediately or their entered data will be lost. Here's [a short and nice video](https://www.youtube.com/watch?v=WveRq-tlb6I) for you to get what I'm saying.

