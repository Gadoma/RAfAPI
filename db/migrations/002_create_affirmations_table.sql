CREATE TABLE "affirmations" (
	"id"	        INTEGER NOT NULL,
	"text"	        TEXT NOT NULL,
	"category_id"	INTEGER NOT NULL,
	"created_at"	TEXT NOT NULL,
	"updated_at"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("category_id") REFERENCES "categories"("id") ON DELETE RESTRICT
);

CREATE INDEX "idx_category_id" ON "affirmations" ( "category_id" ASC );
CREATE UNIQUE INDEX "idx_uniq_text" ON "affirmations" ( "text" );