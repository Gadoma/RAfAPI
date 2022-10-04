CREATE TABLE "affirmations" (
	"id"	        TEXT NOT NULL,
	"text"	        TEXT NOT NULL,
	"category_id"	INTEGER NOT NULL,
	"created_at"	TEXT NOT NULL,
	"updated_at"	TEXT NOT NULL,
	PRIMARY KEY("id"),
	FOREIGN KEY("category_id") REFERENCES "categories"("id") ON DELETE RESTRICT
);

CREATE INDEX "idx_category_id" ON "affirmations" ( "category_id" ASC );
CREATE UNIQUE INDEX "idx_uniq_text" ON "affirmations" ( "text" );

INSERT INTO "main"."categories" ("id", "name") VALUES ('01GEJ0CR9DWN7SA1QBSJE4DVKF', 'Gratitude');
INSERT INTO "main"."categories" ("id", "name") VALUES ('01GEJ0CRM2JW0KY2Z4R5CH4349', 'Health');
INSERT INTO "main"."categories" ("id", "name") VALUES ('01GEJ0CRYJ1AAGQZDS9BR13AKS', 'Money');
INSERT INTO "main"."categories" ("id", "name") VALUES ('01GEJ0CS926M3GV1V1HXQY13AX', 'Intelligence');
INSERT INTO "main"."categories" ("id", "name") VALUES ('01GEJ0CSKKVPY3PR8VXJMPDQAY', 'Empty');
