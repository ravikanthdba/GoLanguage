Cookie is a small file, which maintains the information like a Unique Identifier.
When we send a request to the client, the client then checks to see, if there is a cookie for this domain on our machine.
If yes, then we send the cookie along with the request we make.
The cookie is then checked, to see if the user is already logged in, or is the session expired.
This helps to keep the states, and hence is called a stateful system.

We can send cookies by adding the UUIDs to the routes as well.

Cookies are domain specific. If we send a request to Pepsi, and another request to Coke, we have 2 cookies:

    1 Cookie with Respect to Coke
    2 Cookie with Respect to Pepsi
    
When we send a request to Pepsi again, the cookie of Pepsi will be taken into consideration and a request would be sent.

Please Note:
    The name of the cookie should'nt have a space character 