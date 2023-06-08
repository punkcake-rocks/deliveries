### Evermile

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/punkcake-rocks/deliveries)

OAuth2 client_credentials


Shopify order created -> create evermore order if delivery required
Shopify order canceled -> cancel evermile order
 

Evermile order created -> update Shopify with tracking link
Evermile order delivered -> set Shopify order as fulfilled



Generate $TOKEN with:
```
echo -n ${CLIENT_ID}:${CLIENT_SECRET} | base64
```

Get JWT with:
```
curl --request POST 'https://auth.sandbox.evermile.io/oauth2/token' \
--header 'Authorization: Basic $TOKEN' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials'
```
