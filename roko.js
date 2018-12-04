use rokobookdb;
db.createUser({
  user: "roko",
  pwd: "rokoroko",
  customData: { "level": "god" },
  roles: [
	  { role: "readWrite", db: "rokobookdb" }
  ]
});
