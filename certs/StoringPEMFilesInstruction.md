## Required steps for storing 3 .pem files in this ./certs directory

### What files do I need to store?
- certificate.pem
- key.pem
- ca-chain.pem

### [IMPORTANT] How can I decrypt key.pem file?
You need to decrypt key.pem with `cd ./certs openssl ec -in key.pem -out key.pem` from project root directory.

### Where do I need to store those 3 .pem files?
Please store at passkit-golang-members-quickstart/certs directory which is the same directory as this README file.

### Where can I get those 3 .pem files?
Those files are emailed to you right after account creation or you can find them on Settings (click gear icon in top right of the PassKit.IO web app window) > Developer Credential page (https://dev-app.passkit.io/login).

### Why do I need to store .pem files?
.pem files are used to authenticate you to connect with PassKit server.
 