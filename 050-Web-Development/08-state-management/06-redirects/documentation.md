Request Status Codes:
    
    - 301 - page moved permanently
    - 303 - see other
    - 307 - temporary redirect


Syntax:

    request line = method [SP] request target [SP] HTTP-version [CTLF]
    response line = HTTP version [SP] status-code [SP] reason-phase[CTLF]

    [SP] - Spaces
    [CTLF] - Character terminated Line Field

Examples:

request line:

    Example1:
    method = GET
    request target = /dog
    HTTP-version = HTTP/1.1
    
    Example2:    
    method = POST
    request target = /cat
    HTTP-version = HTTP/1.1  

response line:

    Example1:
    HTTP-version = HTTP/1.1
    status-code = 200
    reason-phase = OK

    Example2:
    HTTP-version = HTTP/1.1
    status-code = 404
    reason-phase = Not Found

More Explanation:

client / browser - sends a request to a target
server side - Option 1 the page has permanently moved
            - Option 2 the server computes the logic in the page and re-directs to a different page as an output.
            
If the existing page has been moved (i.e. Option 1), the first time we request for the page, the following happens:

1) The browser goes to the requested page
2) It figures that the page no longer exists by checking out status code 301
3) It then goes to redirected page.
4) The next time we request for this, the browser directly hits the redirected page.
 