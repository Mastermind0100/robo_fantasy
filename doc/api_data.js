define({ "api": [
  {
    "type": "get",
    "url": "/user/:username/details",
    "title": "Register a new user",
    "group": "User",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "size": "0-2",
            "optional": false,
            "field": "status",
            "description": "<p>0 : success, 2 : unknown err</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "firstname",
            "description": "<p>First Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "lastname",
            "description": "<p>Last Name</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>Email</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "points",
            "description": "<p>points</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "GetUserUsernameDetails"
  },
  {
    "type": "post",
    "url": "/user/change",
    "title": "Change First and Last Name of the user",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>Username</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "firstname",
            "description": "<p>First Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "lastname",
            "description": "<p>Last Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>JWT Auth Token</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-body Example:",
          "content": "{ \"firstname\":\"anshuman\",\n  \"lastname\":\"chhapolia\",\n  \"token\":\"jfsadhdsafh2121h4kjh213iu2hih@$rq3r12\"}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "size": "0-2",
            "optional": false,
            "field": "status",
            "description": "<p>0 : success, 2 : unknown err</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "PostUserChange"
  },
  {
    "type": "post",
    "url": "/user/change/password",
    "title": "Change Password of the users",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>Username</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "newPassword",
            "description": "<p>new Password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>JWT Auth Token</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "size": "0-2",
            "optional": false,
            "field": "status",
            "description": "<p>0 : success, 2 : unknown err</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "PostUserChangePassword"
  },
  {
    "type": "post",
    "url": "/user/login",
    "title": "login with username or email",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "email",
            "description": "<p>Email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "username",
            "description": "<p>Username</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>Password</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Login Request-Body Example:",
          "content": "{\"email\":\"achhap.10.01@gmail.com\",\n  \"username\":\"achhapolia10\",\n  \"password\":\"anshu123\"}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "size": "0-2",
            "optional": false,
            "field": "status",
            "description": "<p>0 : success, 2 : unknown err</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>JWT Auth token to be saved and pass for further requests</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "PostUserLogin"
  },
  {
    "type": "post",
    "url": "/user/new",
    "title": "Register a new user",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "firstname",
            "description": "<p>First Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "lastname",
            "description": "<p>Last Name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>Email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>Username</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>Password</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Login Request-Body Example:",
          "content": "{ \"firstname\":\"anshuman\",\n  \"lastname\":\"chhapolia\",\n  \"email\":\"achhap.10.01@gmail.com\",\n  \"username\":\"achhapolia10\",\n  \"password\":\"anshu123\"}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "size": "0-2",
            "optional": false,
            "field": "status",
            "description": "<p>0 : success, 1 : username/email already exist, 2 : unknown err</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>JWT Auth token to be saved and pass for further requests</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "PostUserNew"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "_home_anshuman_projects_robovitics_robo_fantasy_backend_doc_main_js",
    "groupTitle": "_home_anshuman_projects_robovitics_robo_fantasy_backend_doc_main_js",
    "name": ""
  }
] });
