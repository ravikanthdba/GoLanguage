Session:

	- We store the unique ID in a cookie
	- On our server, we associate each user with a cookie
	- This allows us to identify each and every user visiting the website

Security:

	- Uniqueness of every ID
	- Encryption in transit with HTTPS


Unique ID can be created in 2 ways:

	- Use UUID
	- Using a database. If using a DB, then make sure its not key to the user, but a separate session table
