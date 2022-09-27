CREATE TABLE "categories" (
	"id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE UNIQUE INDEX "idx_uniq_name" ON "categories" ( "name" );