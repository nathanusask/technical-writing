# Foundation is Key: Introduction


Ever since I joined Vendasta, I have been working on different projects with different teams. In a lot of times after a whole day of working I didn't feel the energy to consolodate the foundation of the technologies I used. From time to time I just felt that I was being overly occupied with stories and bugs (and meetings for sure). Other times I would even put up considerable time after regular working hours to deliver values. Nothing seems to be terribly wrong so far because delivering values is the most important thing right? Well, that depends.

On one hand, I did find that I had gained substantial experience in development and to the date I'm confident to say that I've developed a great skill in navigating and driving; on the other hand, I might not be able to tell why we should do certain things in certain ways. Oftentimes I just managed to convince myself and teammates that there was some existing code that solved a similar problem and it worked. But that didn't answer the question why they worked; also, no technology is perfect and there must be a reason why we chose to use some in the first place. To me personally, technical skills and knowledge make who we are as developers, not our searching skills. There's a joke about stackoverflow saying that when stackoverflow crashes then the stackoverflow engineers will have trouble, because they can't use stackoverflow to troubleshoot what causes the crash. LMAO!

In this post, I'll summarize what I've done for my first ever professional development, my plan, as well as my motivaion, Finally and hopefully, the materials I've read and the videos I've watched can help us to consolidate our foundations so that we don't get *stackoverflowed* in future.


## Plan & Motivation

1. Understand terminologies and consolicate basic computer science concepts
	- middleware (to make our codebase more structured and efficient)
	- dockerization, containerization, orchestration, cgroups and namesppaces (dive deeper into how CI/CD works behind the scenes)
	- [coroutine](https://www.youtube.com/watch?v=VwBnPZQl5m8) vs thread vs process (know how programs run in different settings and units)
	- [API gateway](https://www.youtube.com/watch?v=vHQqQBYJtLI)
	- [database transaction priciples: ACID](https://www.geeksforgeeks.org/acid-properties-in-dbms/)
	- observables vs promises
	- [http vs https]()
	- [http 1.1 vs http 2.0](HTTP/1.1 vs HTTP/2: What's the Difference?)
1. Learn different web development frameworks
	- Django (remember Graham did a hackathon project to protect our partner forms from CSRF attacks, Django may offer such a capability)
	- Angular vs React vs Vue (to have a broader view over different frameworks to facilitate our future decision/migration)
1. Learn Redis, [SQL vs NoSQL](https://redislabs.com/ebook/part-1-getting-started/chapter-1-getting-to-know-redis/1-1-what-is-redis/1-1-1-redis-compared-to-other-databases-and-software/) (redis is best known for its caching, but it also provides features such as querying, locking, throttling and Pub/Sub; data is always in memory so Redis is fast; redis is single-threaded so you may consider its master-slave scaling to serve more requests.)
1. Consolidate golang microservice development essence (learn prior to the cadence workflow, how did we do to make use of the golang concurrency)
	- goroutines and channels
	- gRPC

- What next?
----------------
Learn rust: rust is comparatively new to the industry as its first stable release was in May 2015. The language, as you may heard, has multiple advantages over other languages such as it is fast, memory-safe, concurrency-safe and etc. The main downside of the language is that its learning curve is long and the code isn't as readable as golang or your familiar mainstream languages. However, to get ahead of the game, we may start learning it sooner rather than later.


## Materials

### Frontend
1. [JavaScript Theory: Promise vs Observable](https://medium.com/javascript-everyday/javascript-theory-promise-vs-observable-d3087bc1239a)
2. [Inheritance in JavaScript](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/Inheritance)
3. [You Dont Know JS](https://github.com/getify/You-Dont-Know-JS)

### HTTP
1. [HTTP/1.1 vs HTTP/2: What's the Difference?](https://www.digitalocean.com/community/tutorials/http-1-1-vs-http-2-what-s-the-difference)
2. [Why is HTTP not secure? | HTTP vs. HTTPS](https://www.cloudflare.com/learning/ssl/why-is-http-not-secure/)

### Middleware
1. [What is middleware?](https://www.redhat.com/en/topics/middleware/what-is-middleware#:~:text=Middleware%20is%20software%20that%20provides,all%20commonly%20handled%20by%20middleware.)
2. [Securing Your Go REST APIs with JSON Web Tokens(JWTS)] (https://www.youtube.com/watch?v=-Scg9INymBs)
3. [Middleware](https://www.youtube.com/watch?v=HOlklLaFgfM)
4. [Angular route guards](https://www.youtube.com/watch?v=WveRq-tlb6I)
5. [My own practice with middleware](https://github.com/nathanusask/backend-studies/tree/main/middleware)
6. [PR: use middleware to parse JWT tokens for cloud-blue asset request](https://github.com/vendasta/cloud-blue/pull/18)

### Django
1. [Python Django Web Framework - Full Course for Beginners](https://www.youtube.com/watch?v=F5mRW0jo-U4)
1. [Django Rest Framework & React Tutorial: Learning Management System (Blackboard / Moodle Clone)](https://www.youtube.com/watch?v=JIFqqdRxmVo)
2. [CSRF Token with Django web application](https://www.codementor.io/@venkatingen/csrf-token-with-django-web-application-mxn85nm8x)
3. [Django middleware documentation](https://docs.djangoproject.com/en/2.2/ref/middleware/)
3. [Django security middleware codebase](https://github.com/django/django/tree/master/django/middleware)


### Redis
1. [Redis in Action](https://redislabs.com/ebook/redis-in-action/)
2. [My own practice with redis in golang to implement a voting system](https://github.com/nathanusask/backend-studies/tree/main/redis/voting_go)
3. [The official documentation has certain errors in using redis python, this is a refined version of the voting system in the textbook example](https://github.com/nathanusask/backend-studies/tree/main/redis/voting_py)


### Golang microservice
1. [What is the use of an atomic package in golang?](https://www.quora.com/What-is-the-use-of-an-atomic-package-in-golang)
2. [Concurrency Patterns In Go](https://www.youtube.com/watch?v=YEKjSzIwAdA)
3. [GopherCon 2017: Kavya Joshi - Understanding Channels](https://www.youtube.com/watch?v=KBZlN0izeiY)
4. [Remote Procedure Call (RPC)](https://searchapparchitecture.techtarget.com/definition/Remote-Procedure-Call-RPC)
5. [gRPC Core Concepts, Architechtures and lifecycle](https://grpc.io/docs/what-is-grpc/core-concepts/)

### Dockerization, Containerization, Orchestration
1. [Microservices + Events + Docker = A Perfect Trio](https://www.youtube.com/watch?v=sSm2dRarhPo)
2. [This Is How Docker Works, The Fun Way!](https://www.youtube.com/watch?v=-NzfOhSAZpA)
3. [Kubernetes vs. Docker: It's Not an Either/Or Question](https://www.youtube.com/watch?v=2vMEQ5zs1ko)
4. [Should You Use Kubernetes and Docker In Your Next Project?](https://www.youtube.com/watch?v=u8dW8DrcSmo)

### Databases
1. [SQL vs NoSQL or MySQL vs MongoDB](https://www.youtube.com/watch?v=ZS_kXvOeQ5Y)
