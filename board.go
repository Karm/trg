package main

// Board struct
type Board struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`

	/*"id": "5cf12faef016e615c7a8c945",
	"name": "JWS/JBCS",
	"desc": "",
	"descData": null,
	"closed": false,
	"idOrganization": "5600ea9f2ab78b83fe63e68b",
	"url": "https://trello.com/b/kBoVdK8D/jws-jbcs",
	"labelNames": {
		"green": "OpenShift",
		"yellow": "OneOff",
		"orange": "Debug",
		"red": "JWS",
		"purple": "Infrastructure",
		"blue": "JBCS",
		"sky": "Upstream",
		"lime": "Docs",
		"pink": "R&D",
		"black": "Paperwork"
	  },
	"shortUrl": "https://trello.com/b/kBoVdK8D",
	"labels": [
		{
		  "id": "5cf12fae91d0c2ddc5f069e8",
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "name": "Infrastructure",
		  "color": "purple"
		},
		{
		  "id": "5cf15b38323ccf2d910a074c",
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "name": "NOE",
		  "color": null
		}
	  ],
	"lists": [
		{
		  "id": "5cf15622d8921f5b658bdaf8",
		  "name": "JBQA JIRA Backlog",
		  "closed": true,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 32767.5,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		},
		{
		  "id": "5cf2d70bb4c381303a08c215",
		  "name": "JBQA Backlog",
		  "closed": false,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 49151.25,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		},

		{
		  "id": "5cf1347908853c5313faa0e0",
		  "name": "R&D",
		  "closed": true,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 262143,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		},
		{
		  "id": "5cf13483c54ad4560729f9b1",
		  "name": "Paperwork",
		  "closed": true,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 327679,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		},
		{
		  "id": "5cf134bc3dbdb1158cf69120",
		  "name": "Doing Week 23 2019",
		  "closed": false,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 393215,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		},
		{
		  "id": "5cf134c16520d4431712397f",
		  "name": "Done Week 23 2019",
		  "closed": false,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 458751,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		},
		{
		  "id": "5cf4cd76479597506b1a9e97",
		  "name": "Doing Week 24 2019",
		  "closed": false,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 524287,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		},
		{
		  "id": "5cf4e7481f87f384994afb68",
		  "name": "Done Week 24 2019",
		  "closed": false,
		  "idBoard": "5cf12faef016e615c7a8c945",
		  "pos": 589823,
		  "subscribed": false,
		  "softLimit": null,
		  "limits": {
			"cards": {
			  "openPerList": {
				"status": "ok",
				"disableAt": 5000,
				"warnAt": 4500
			  },
			  "totalPerList": {
				"status": "ok",
				"disableAt": 1000000,
				"warnAt": 900000
			  }
			}
		  },
		  "creationMethod": null
		}
	],
	"members": [
		{
		  "id": "57f2517d0e61ff9e8f42dbb7",
		  "avatarHash": null,
		  "avatarUrl": null,
		  "bio": "",
		  "bioData": {
			"emoji": {}
		  },
		  "confirmed": true,
		  "fullName": "Jan Onderka",
		  "idEnterprise": null,
		  "idEnterprisesDeactivated": null,
		  "idMemberReferrer": null,
		  "idPremOrgsAdmin": [],
		  "initials": "JO",
		  "memberType": "normal",
		  "nonPublic": {},
		  "nonPublicAvailable": false,
		  "products": [],
		  "url": "https://trello.com/janonderka1",
		  "username": "janonderka1",
		  "status": "disconnected"
		},
		{
		  "id": "559d4dff02dc3087b9800dfc",
		  "avatarHash": "a55fe89a2429bccadb2daef1e1dc667e",
		  "avatarUrl": "https://trello-avatars.s3.amazonaws.com/a55fe89a2429bccadb2daef1e1dc667e",
		  "bio": "",
		  "bioData": null,
		  "confirmed": true,
		  "fullName": "Karm",
		  "idEnterprise": "5aebf7d0a8dff42c393d28c0",
		  "idEnterprisesDeactivated": [],
		  "idMemberReferrer": null,
		  "idPremOrgsAdmin": [
			"5600ea9f2ab78b83fe63e68b"
		  ],
		  "initials": "K",
		  "memberType": "normal",
		  "nonPublic": {},
		  "nonPublicAvailable": false,
		  "products": [
			10,
			37
		  ],
		  "url": "https://trello.com/karm2",
		  "username": "karm2",
		  "status": "disconnected"
		},
		{
		  "id": "59ca1601ece19a35a850a836",
		  "avatarHash": null,
		  "avatarUrl": null,
		  "bio": "",
		  "bioData": null,
		  "confirmed": true,
		  "fullName": "Matus Madzin",
		  "idEnterprise": null,
		  "idEnterprisesDeactivated": null,
		  "idMemberReferrer": null,
		  "idPremOrgsAdmin": [],
		  "initials": "MM",
		  "memberType": "normal",
		  "nonPublic": {},
		  "nonPublicAvailable": false,
		  "products": [],
		  "url": "https://trello.com/matusmadzin1",
		  "username": "matusmadzin1",
		  "status": "disconnected"
		},
		{
		  "id": "5cf50b2868c4844f361b542e",
		  "avatarHash": null,
		  "avatarUrl": null,
		  "bio": "",
		  "bioData": null,
		  "confirmed": true,
		  "fullName": "Paul Lodge",
		  "idEnterprise": null,
		  "idEnterprisesDeactivated": null,
		  "idMemberReferrer": null,
		  "idPremOrgsAdmin": [],
		  "initials": "PL",
		  "memberType": "normal",
		  "nonPublic": {},
		  "nonPublicAvailable": false,
		  "products": [],
		  "url": "https://trello.com/paullodge2",
		  "username": "paullodge2",
		  "status": "disconnected"
		},
		{
		  "id": "57ea23286c1658e696abcd9a",
		  "avatarHash": null,
		  "avatarUrl": null,
		  "bio": "",
		  "bioData": null,
		  "confirmed": true,
		  "fullName": "jstefl",
		  "idEnterprise": null,
		  "idEnterprisesDeactivated": null,
		  "idMemberReferrer": null,
		  "idPremOrgsAdmin": [
			"5600ea9f2ab78b83fe63e68b"
		  ],
		  "initials": "J",
		  "memberType": "normal",
		  "nonPublic": {},
		  "nonPublicAvailable": false,
		  "products": [],
		  "url": "https://trello.com/jstefl",
		  "username": "jstefl",
		  "status": "disconnected"
		}
	  ],









	*/
}
