package main

//////////////////////////////////////////////////////////////////////////////////
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
                                  2 : unknown err
* @apiSuccess {String} token JWT Auth token to be saved and pass for further requests
*
*/

/**
* @api {get} 	/user/:username/details Get details of the user
* @apiGroup User
*
* @apiSuccess {Number{0-2}} status 0 : success,
                                  2 : unknown err
* @apiSuccess {String} firstname First Name
* @apiSuccess {String} lastname Last Name
* @apiSuccess {String} email Email
* @apiSuccess {Number} points points
*/

/////////////////////////////////////////////////////////////////////////
