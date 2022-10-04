CREATE TABLE "categories" (
	"id"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	PRIMARY KEY("id")
);

CREATE UNIQUE INDEX "idx_uniq_name" ON "categories" ( "name" );
