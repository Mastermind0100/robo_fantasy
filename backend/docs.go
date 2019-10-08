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
*/

/**
* @api {post} 	/user/change Change First and Last Name of the user
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
* @api {post} 	/user/change/password Change Password of the users
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
*/

/**
@api {get} /bot/:id/delete Delete a bot
@apiGroup Bots

@apiSuccess {Number=0,1} status Status 0:Sucess, 1:error
*/

/**
@api {get} /bot/all get details for all the bots
@apiGroup Bots

@apiSuccess {Number=0,1} status Status 0:Success, 1:error
@apiSuccess {Object[]} data Details of all the bots structure like the single bot
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
@api {post} /match/new
@apiGroup Matches

@apiParam {String} blue Blue Bot name
@apiParam {String} red Red Bot name
@apiParam {Number=15,30,60} category Category
@apiParam {Number{0-4}} [status] 0:upcoming, 1:red, 2:blue, 3:draw, if no params given upcoming



@apiSuccess {Number=0,1} status Status 0:Success, 1:error
*/

/**
@api {get} /match/:id/delete
@apiGroup Matches

@apiSuccess {Number=0,1} status Status 0:Success, 1:error
*/

/**
@api {post} /match/:id/status
@apiGroup Matches

@apiParam {Number=0,1,2,3} status 0:upcoming, 1:red won, 2:blue won, 3: draw

@apiSuccess {Number=0,1} status Status 0:Success, 1:error
*/
