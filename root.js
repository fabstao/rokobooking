db.createUser({
  user: "root",
  pwd: "austin23",
  customData: { "level": "god" },
  roles: [
	  { role: "userAdminAnyDatabase", db: "admin" }
  ]
})
