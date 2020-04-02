# Bootstrapping New Frontend UI-- Sharing my experience from creating the web crawler UI

You know what Graham, bootstrapping isn't just a term referring to building up
a brand new microservice
and here's my supportive [statement](#https://www.techopedia.com/definition/3328/bootstrap):
"A bootstrap is the program that initializes the operating system (OS) during startup.
The term bootstrap or bootstrapping originated in the early 1950s.
It referred to a bootstrap load button that was used to initiate a hardwired bootstrap program,
or smaller program that executed a larger program such as the OS."

## Setting up github code repository and cloudbuild

## Authentication

## Redirection between auth and app pages with different environments
The app should operate in a given environment (e.g., local, demo, prod, test, etc.) and
make network calls based on the environment. The environment would then make different
for things such as authentication tokens so that everything is aligned;
e.g., if I'm running the frontend locally or accessing the demo app,
the app should make API calls to the demo
endpoints and get data from vStore demo with a demo bearer token.
So how do we do that?
In your `index.html` file, *never ever* forget about the following commented out lines:
```html
<!-- vStaticInject:environment -->
<!-- vStaticInject:deployment -->
<!-- vStaticInject:iam -->
```
why? According to Graham:
When you run `fecli deploy` you are sending off all of your build artifacts to vstatic;
then when you go to your app in the browser, vstatic serves those files.
But right before it serves `index.html` it checks for those html comments to see
if they're there and if they are it just replaces them with the pertinent information.

## Server setup to serve the frontend app

## Setup service provider
1. Register service provider (_repeat it for prod as the example is for demo_)
    ```shell script
    curl --location --request POST 'https://sso-api-demo.vendasta-internal.com/sso.v1.ServiceProviderAdmin/CreateServiceProvider' \
    --header 'content-type: application/json' \
    --header 'authorization: Bearer <bear token>' \
    --data-raw '{
        "service_provider": {
        "entry_url": "https://<your_app-name>-demo.apigatway.co/",
        "owned_and_operated": true,
        "primary_domain": "https://<your_app_name>-demo.apigatway.co/",
        "service_provider_id": "<your_service_provider_id>",
        "session_url": "",
        "logout_url": ""
        }
    }'
    ```

1. Obtain client IDs (_repeat for prod as the example is for demo_)
    ```shell script
    curl --location --request POST 'https://sso-api-demo.apigateway.co/sso.v1.ServiceProviderAdmin/CreateConfig' \
    --header 'Content-Type: application/json' \
    --header 'authorization: Bearer <bear token>' \
    --data-raw '{
        "name": "web-crawler-client",
        "service_provider_id": "<your_service_provider_id>",
        "redirect_uris": ["https://<your_app_name>-demo.vendasta-internal.com"]
    }'
    ```
    Grab the client ID from here and use it in the next step.

1. Update redirect URLs for service provider (_repeat for prod
and get rid of_ `localhost` _as the example is for demo_)
    ```shell script
    curl --location --request POST 'https://sso-api-demo.vendasta-internal.com/sso.v1.ServiceProviderAdmin/UpdateConfig' \
    --header 'content-type: application/json' \
    --data-raw '{
        "mutations": [
            {
                "redirect_uris": {
                    "repeated_string": ["http://localhost:4200", "https://<your_app_name>-demo.vendasta-internal.com"]
                }
            }
        ],
        "client_id": "<client_id>"
    }'
    ```

## Dependencies and troubleshooting for missing service providers

## Code format

### Prettier

### Code smells
1. Nested `pipe`'s
1. Nested `subscribe`'s

## Asking good questions


## Acknowledgements
Special thanks go out to Taylor Wiebe (for helping me go through
the auth process),
Graham Holtslander (for expert frontend opinions),
Jonathan Baxter (for helping me with everything)
as well as the remainder of the data team for welcoming me 
and for walking me through what we do in this team.