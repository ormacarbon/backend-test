# Backend test
#### Install the dependencies
`
npm install
`
#### Create a .env file and set up the credentials as the file .env.example
`
touch .env
`
#### Run the application in dev mode on port 8000
`
npm run dev
`

## The endpoints:

### Create user
`
 post.{{ _.baseUrl }}/user/create
`
###### JSON 
`
{
	"name": "test name",
	"email": "teste@gmail.com",
	"phone": 63984789553
}
`
### Create linked user, the user and the original user will receive one point

`
 post.{{ _.baseUrl }}/user/create/linkedUser?id={{ id }}
`
###### JSON 
`
{
	"name": "test name linkedu ser",
	"email": "teste2@gmail.com",
	"phone": 63978456365
}
`
### Get top 10 winners
`
 get.{{ _.baseUrl }}/user/winners
`
