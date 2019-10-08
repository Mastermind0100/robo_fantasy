define({ "api": [
  {
    "type": "get",
    "url": "/bot/all",
    "title": "get details for all the bots",
    "group": "Bots",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "allowedValues": [
              "0",
              "1"
            ],
            "optional": false,
            "field": "status",
            "description": "<p>Status 0:Success, 1:error</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "data",
            "description": "<p>Details of all the bots structure like the single bot</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Bots",
    "name": "GetBotAll"
  },
  {
    "type": "get",
    "url": "/bot/:id",
    "title": "get detail for a bot id",
    "group": "Bots",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "allowedValues": [
              "0",
              "1"
            ],
            "optional": false,
            "field": "status",
            "description": "<p>Status 0:Success, 1:error</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Single bot response:",
          "content": "{ \"status\":0,\n\t\"data\":{\"name\":\"bot-name\",\"team\":\"team-name\",\n\t\t\t\"category\":15, \"description\":\"bot-description\",\n\t\t\t\"video\":\"video link\",\"img\":\"image-link\"}}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Bots",
    "name": "GetBotId"
  },
  {
    "type": "get",
    "url": "/bot/:id/delete",
    "title": "Delete a bot",
    "group": "Bots",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "allowedValues": [
              "0",
              "1"
            ],
            "optional": false,
            "field": "status",
            "description": "<p>Status 0:Sucess, 1:error</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Bots",
    "name": "GetBotIdDelete"
  },
  {
    "type": "post",
    "url": "/bot/new",
    "title": "Create a new bot",
    "group": "Bots",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Bot name</p>"
          },
          {
            "group": "Parameter",
            "type": "Team",
            "optional": false,
            "field": "team",
            "description": "<p>Team name</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "allowedValues": [
              "15",
              "30",
              "60"
            ],
            "optional": false,
            "field": "category",
            "description": "<p>Category</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "description",
            "description": "<p>Description</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "video",
            "description": "<p>Youtube link</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "image",
            "description": "<p>Image Link</p>"
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
            "allowedValues": [
              "0",
              "1"
            ],
            "optional": false,
            "field": "status",
            "description": "<p>Status 0:Sucess, 1:error</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Bots",
    "name": "PostBotNew"
  },
  {
    "type": "get",
    "url": "/user/:username/details",
    "title": "Get details of the user",
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
            "description": "<p>0 : success, 1: user not found 2 : unknown err</p>"
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
            "description": "<p>0 : success, 1: username/email password mismatch, 2 : unknown err</p>"
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
  }
] });
