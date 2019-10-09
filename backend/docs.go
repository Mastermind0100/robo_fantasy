package main

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/**
* @api {post} 	/user/new Register a new user
* @apiGroup User
*
* @apiParam {String} firstname First Name
* @apiParam {String} lastname Last Name
* @apiParam {String} email Email
* @apiParam {String} username Username
* @apiParam {String} password Password
*
* @apiParamExample {json} Login Request-Body Example:
*		{ "firstname":"anshuman",
*		  "lastname":"chhapolia",
*		  "email":"achhap.10.01@gmail.com",
*		  "username":"achhapolia10",
*		  "password":"anshu123"}
*
* @apiSuccess {Number{0-2}} status 0 : success,
     							    1 : username/email already exist,
                                  2 : unknown err
* @apiSuccess {String} token JWT Auth token to be saved and pass for further requests
*
*@apiSuccessExample {json} Successful Login Response:
{"status":0,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdE5hbWUiOiJhbnNodSIsImxhc3ROYW1lIjoiYW5zaHUiLCJ1c2VybmFtZSI6ImFjaGhhcG9saWExMCIsImVtYWlsIjoiYW5zaHVAdGVzdC5jb20iLCJjaS11c2VybmFtZSI6IkFDSEhBUE9MSUExMCJ9.Yrr__2aHEULhPaFJ9VBsYgLJR0EJ3JC-rSFB0s-Zq0M"}
*/

/**
* @api {post} 	/user/change Change First and Last Name of the user (not implemented yet so no examples)
* @apiGroup User
*
* @apiParam {String} username Username
* @apiParam {String} firstname First Name
* @apiParam {String} lastname] Last Name
* @apiParam {String} token JWT Auth Token
*
* @apiParamExample {json} Request-body Example:
*		{ "firstname":"anshuman",
*		  "lastname":"chhapolia",
*		  "token":"jfsadhdsafh2121h4kjh213iu2hih@$rq3r12"}
*
*
* @apiSuccess {Number{0-2}} status 0 : success,
                                  2 : unknown err
*
*/

/**
* @api {post} 	/user/change/password Change Password of the users (not implemented yet so no examples)
* @apiGroup User
*
* @apiParam {String} username Username
* @apiParam {String} newPassword new Password
* @apiParam {String} token JWT Auth Token
*
@apiSuccess {Number{0-2}} status 0 : success,
                                 2 : unknown err
*/

/**
* @api {post} 	/user/login login with username or email
* @apiGroup User
*
* @apiParam {String} [email] Email
* @apiParam {String} [username] Username
* @apiParam {String} password Password
*
* @apiParamExample {json} Login Request-Body Example:
*		{"email":"achhap.10.01@gmail.com",
*		  "username":"achhapolia10",
*		  "password":"anshu123"}
*
* @apiSuccess {Number{0-2}} status 0 : success,
									1: username/email password mismatch,
                                  2 : unknown err
* @apiSuccess {String} token JWT Auth token to be saved and pass for further requests

@apiSuccessExample {json} Successful Login Response:
{"status":0,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdE5hbWUiOiJhbnNodSIsImxhc3ROYW1lIjoiYW5zaHUiLCJ1c2VybmFtZSI6ImFjaGhhcG9saWExMCIsImVtYWlsIjoiYW5zaHVAdGVzdC5jb20iLCJjaS11c2VybmFtZSI6IkFDSEhBUE9MSUExMCJ9.Yrr__2aHEULhPaFJ9VBsYgLJR0EJ3JC-rSFB0s-Zq0M"}

*
*/

/**
* @api {get} 	/user/:username/details Get details of the user
* @apiGroup User
*
* @apiSuccess {Number{0-2}} status 0 : success,
									1: user not found
                                  2 : unknown err
* @apiSuccess {String} firstname First Name
* @apiSuccess {String} lastname Last Name
* @apiSuccess {String} email Email
* @apiSuccess {Number} points points

@apiSuccessExample {json} Details for the user:
{"status":0,"firstname":"anshu","lastname":"anshu","email":"anshu@test.com","points":0}
*/

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
@api {post} /bot/new Create a new bot
@apiGroup Bots

@apiParam  {String} name Bot name
@apiParam  {Team} team Team name
@apiParam {Number=15,30,60} category Category
@apiParam {String} description Description
@apiParam {String} video Youtube link
@apiParam {String} image Image Link

@apiSuccess {Number=0,1} status Status 0:Sucess, 1:error

@apiSuccessExample {json} Response on Adding a bot:
{"status":0}
*/

/**
@api {get} /bot/:id/delete Delete a bot
@apiGroup Bots

@apiSuccess {Number=0,1} status Status 0:Sucess, 1:error
@apiSuccessExample {json} Response on Deleting a bot:
{"status":0}
*/

/**
@api {get} /bot/all get details for all the bots
@apiGroup Bots

@apiSuccess {Number=0,1} status Status 0:Success, 1:error
@apiSuccess {Object[]} data Details of all the bots structure like the single bot

@apiSuccessExample {json} Response for all bots details
{"status":0,"data":[{"id":1,"name":"Bot","team":"Team","category":15,"description":"test bot","video":"","img":"jfdslakj"},{"id":2,"name":"Bot1","team":"Team1","category":151,"description":"test bot1","video":"","img":"jfdslakj1"}]}
*/

/**
@api {get} /bot/:id get detail for a bot id
@apiGroup Bots

@apiSuccess {Number=0,1} status Status 0:Success, 1:error

@apiSuccessExample {json} Single bot response:
	{ "status":0,
		"data":{"id":1,"name":"bot-name","team":"team-name",
				"category":15, "description":"bot-description",
				"video":"video link","img":"image-link"}}
*/

/////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
@api {post} /match/new Add a new Match
@apiGroup Matches

@apiParam {String} blue Blue Bot name
@apiParam {String} red Red Bot name
@apiParam {Number=15,30,60} category Category
@apiParam {Number{0-4}} [status] 0:upcoming, 1:red, 2:blue, 3:draw, if no params given upcoming



@apiSuccess {Number=0,1} status Status 0:Success, 1:error
@apiSuccessExample {json} Response on Adding a match:
{"status":0}
*/

/**
@api {get} /match/:id/delete Delete a match of an id
@apiGroup Matches

@apiSuccess {Number=0,1} status Status 0:Success, 1:error
@apiSuccessExample {json} Response on Deleting a match:
{"status":0}
*/

/**
@api {post} /match/:id/status Update status of the matches
@apiGroup Matches

@apiParam {Number=0,1,2,3} status 0:upcoming, 1:red won, 2:blue won, 3: draw

@apiSuccess {Number=0,1} status Status 0:Success, 1:error

@apiSuccessExample {json} Response on updating status:
{"status":0}
*/

/**
@api {get} /match/{id}/details Get Detail of a match using id
@apiGroup Matches
@apiSuccess {Number=0,1} status Status 0:Success, 1:error

@apiSuccessParam {Number} data-status 0:upcoming, 1:red won, 2:blue won, 3: draw
@apiSuccessExample {json} Single bot response:
	{ "status":0,
		"data":{"id":1,"red":"red-team name",
				"blue":"blue-team name", "status":0,
				"category":15}}
*/

/**
@api {get} /match/all Get All match details
@apiGroup Matches

@apiSuccess {Number=0,1} status Status 0:Success, 1:error
@apiSuccessParam {Object[]} data
@apiSuccessExample {json} Single bot response:
	{"status":0,"data":[{"id":1,"blue":"BLUE","red":"RED","category":15,"status":0},{"id":2,"blue":"Red","red":"Blue","category":30,"status":0}]}
*/
