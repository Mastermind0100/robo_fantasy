define({ "api": [
  {
    "type": "post",
    "url": "/bet",
    "title": "Prediction for user",
    "group": "Bet",
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
            "type": "Number",
            "optional": false,
            "field": "matchid",
            "description": "<p>Matchid</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "allowedValues": [
              "\"red\"",
              "\"blue\""
            ],
            "optional": false,
            "field": "team",
            "description": "<p>Team color</p>"
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
            "description": "<p>Status 0:Success, 1:error</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response on updating status:",
          "content": "{\"status\":0}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Bet",
    "name": "PostBet"
  },
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
      },
      "examples": [
        {
          "title": "Response for all bots details",
          "content": "{\"status\":0,\"data\":[{\"id\":1,\"name\":\"Bot\",\"team\":\"Team\",\"category\":15,\"description\":\"test bot\",\"video\":\"\",\"img\":\"jfdslakj\"},{\"id\":2,\"name\":\"Bot1\",\"team\":\"Team1\",\"category\":151,\"description\":\"test bot1\",\"video\":\"\",\"img\":\"jfdslakj1\"}]}",
          "type": "json"
        }
      ]
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
          "content": "{ \"status\":0,\n\t\"data\":{\"id\":1,\"name\":\"bot-name\",\"team\":\"team-name\",\n\t\t\t\"category\":15, \"description\":\"bot-description\",\n\t\t\t\"video\":\"video link\",\"img\":\"image-link\"}}",
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
      },
      "examples": [
        {
          "title": "Response on Deleting a bot:",
          "content": "{\"status\":0}",
          "type": "json"
        }
      ]
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
      },
      "examples": [
        {
          "title": "Response on Adding a bot:",
          "content": "{\"status\":0}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Bots",
    "name": "PostBotNew"
  },
  {
    "type": "get",
    "url": "/match/all",
    "title": "Get All match details",
    "group": "Matches",
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
          "content": "{\"status\":0,\"data\":[{\"id\":1,\"blue\":\"BLUE\",\"red\":\"RED\",\"category\":15,\"status\":0},{\"id\":2,\"blue\":\"Red\",\"red\":\"Blue\",\"category\":30,\"status\":0}]}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Matches",
    "name": "GetMatchAll"
  },
  {
    "type": "get",
    "url": "/match/:id/delete",
    "title": "Delete a match of an id",
    "group": "Matches",
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
          "title": "Response on Deleting a match:",
          "content": "{\"status\":0}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Matches",
    "name": "GetMatchIdDelete"
  },
  {
    "type": "get",
    "url": "/match/{id}/details",
    "title": "Get Detail of a match using id",
    "group": "Matches",
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
          "content": "{ \"status\":0,\n\t\"data\":{\"id\":1,\"red\":\"red-team name\",\n\t\t\t\"blue\":\"blue-team name\", \"status\":0,\n\t\t\t\"category\":15}}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Matches",
    "name": "GetMatchIdDetails"
  },
  {
    "type": "post",
    "url": "/match/:id/status",
    "title": "Update status of the matches",
    "group": "Matches",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "allowedValues": [
              "0",
              "1",
              "2",
              "3"
            ],
            "optional": false,
            "field": "status",
            "description": "<p>0:upcoming, 1:red won, 2:blue won, 3: draw</p>"
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
            "description": "<p>Status 0:Success, 1:error</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response on updating status:",
          "content": "{\"status\":0}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Matches",
    "name": "PostMatchIdStatus"
  },
  {
    "type": "post",
    "url": "/match/new",
    "title": "Add a new Match",
    "group": "Matches",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "blue",
            "description": "<p>Blue Bot name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "red",
            "description": "<p>Red Bot name</p>"
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
            "type": "Number",
            "size": "0-4",
            "optional": true,
            "field": "status",
            "description": "<p>0:upcoming, 1:red, 2:blue, 3:draw, if no params given upcoming</p>"
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
            "description": "<p>Status 0:Success, 1:error</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response on Adding a match:",
          "content": "{\"status\":0}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "Matches",
    "name": "PostMatchNew"
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
      },
      "examples": [
        {
          "title": "Details for the user:",
          "content": "{\"status\":0,\"firstname\":\"anshu\",\"lastname\":\"anshu\",\"email\":\"anshu@test.com\",\"points\":0}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "GetUserUsernameDetails"
  },
  {
    "type": "get",
    "url": "/user/:username/leaderboard",
    "title": "Get the leaderboard for the user",
    "group": "User",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "topten",
            "description": "<p>Top Ten people (Ask me for jsonSchema)</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "rank",
            "description": "<p>Rank for the user</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "points",
            "description": "<p>Points for the user</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "LeaderBoard for the user:",
          "content": "{\"topten\":[{\"user\":{\"firstName\":\"anshu\",\"lastName\":\"anshu\",\"username\":\"anshu2\",\"email\":\"anshu2@test.com\",\"ci-username\":\"ANSHU2\"},\"points\":2},{\"user\":{\"firstName\":\"anshu\",\"lastName\":\"anshu\",\"username\":\"anshu3\",\"email\":\"anshu3@test.com\",\"ci-username\":\"ANSHU3\"},\"points\":1},{\"user\":{\"firstName\":\"anshu\",\"lastName\":\"anshu\",\"username\":\"anshu1\",\"email\":\"anshu@test.com\",\"ci-username\":\"ANSHU1\"},\"points\":0}],\"rank\":3,\"points\":0}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "GetUserUsernameLeaderboard"
  },
  {
    "type": "get",
    "url": "/user/:username/matches",
    "title": "Get results for matches for a user",
    "group": "User",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "data",
            "description": "<p>Result array of further fields define fields of object</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "red",
            "description": "<p>Red Team Name</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "blue",
            "description": "<p>Blue Team name</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "allowedValues": [
              "\"red\"",
              "\"blue\"",
              "\"none\""
            ],
            "optional": false,
            "field": "prediction",
            "description": "<p>Prediction for team color</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "result",
            "description": "<p>Result true or false</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Matches for the user:",
          "content": "{\"data\":[{\"red\":\"RED1\",\"blue\":\"BLUE1\",\"prediction\":\"blue\",\"result\":2},{\"red\":\"RED2\",\"blue\":\"BLUE2\",\"prediction\":\"blue\",\"result\":2},{\"red\":\"RED3\",\"blue\":\"BLUE3\",\"prediction\":\"none\",\"result\":0}]}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "GetUserUsernameMatches"
  },
  {
    "type": "post",
    "url": "/user/change",
    "title": "Change First and Last Name of the user (not implemented yet so no examples)",
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
    "title": "Change Password of the users (not implemented yet so no examples)",
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
      },
      "examples": [
        {
          "title": "Successful Login Response:",
          "content": "{\"status\":0,\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdE5hbWUiOiJhbnNodSIsImxhc3ROYW1lIjoiYW5zaHUiLCJ1c2VybmFtZSI6ImFjaGhhcG9saWExMCIsImVtYWlsIjoiYW5zaHVAdGVzdC5jb20iLCJjaS11c2VybmFtZSI6IkFDSEhBUE9MSUExMCJ9.Yrr__2aHEULhPaFJ9VBsYgLJR0EJ3JC-rSFB0s-Zq0M\"}",
          "type": "json"
        }
      ]
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
      },
      "examples": [
        {
          "title": "Successful Login Response:",
          "content": "{\"status\":0,\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdE5hbWUiOiJhbnNodSIsImxhc3ROYW1lIjoiYW5zaHUiLCJ1c2VybmFtZSI6ImFjaGhhcG9saWExMCIsImVtYWlsIjoiYW5zaHVAdGVzdC5jb20iLCJjaS11c2VybmFtZSI6IkFDSEhBUE9MSUExMCJ9.Yrr__2aHEULhPaFJ9VBsYgLJR0EJ3JC-rSFB0s-Zq0M\"}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./docs.go",
    "groupTitle": "User",
    "name": "PostUserNew"
  }
] });
