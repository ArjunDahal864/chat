CREATE TABLE "user" (
	"id" serial NOT NULL,
	"full_name" VARCHAR(255) NOT NULL,
	"email" VARCHAR(255) NOT NULL UNIQUE,
	"password" VARCHAR(255) NOT NULL,
	"profile_pic" VARCHAR(255) NOT NULL,
	CONSTRAINT "user_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "chat" (
	"id" serial NOT NULL,
	"sender_id" integer NOT NULL,
	"inbox_hash" VARCHAR(255) NOT NULL,
	"msg" TEXT NOT NULL,
	"file" TEXT NOT NULL,
	"meta" VARCHAR(255) NOT NULL,
	"deleted_by" integer NOT NULL,
	CONSTRAINT "chat_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "inbox" (
	"id" serial NOT NULL,
	"user_id" integer NOT NULL,
	"sender_id" integer NOT NULL,
	"inbox_hash" VARCHAR(255) NOT NULL,
	"last_msg" VARCHAR(255) NOT NULL,
	"seen" BOOLEAN NOT NULL,
	"unseen_number" integer NOT NULL,
	"deleted" BOOLEAN NOT NULL,
	CONSTRAINT "inbox_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "friend" (
	"id" serial NOT NULL,
	"user_id" integer NOT NULL,
	"friend_id" integer NOT NULL,
	"blocked_by" integer,
	CONSTRAINT "friend_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "chat" ADD CONSTRAINT "chat_fk0" FOREIGN KEY ("sender_id") REFERENCES "user"("id");
ALTER TABLE "chat" ADD CONSTRAINT "chat_fk1" FOREIGN KEY ("deleted_by") REFERENCES "user"("id");

ALTER TABLE "inbox" ADD CONSTRAINT "inbox_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "inbox" ADD CONSTRAINT "inbox_fk1" FOREIGN KEY ("sender_id") REFERENCES "user"("id");

ALTER TABLE "friend" ADD CONSTRAINT "friend_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "friend" ADD CONSTRAINT "friend_fk1" FOREIGN KEY ("friend_id") REFERENCES "user"("id");
ALTER TABLE "friend" ADD CONSTRAINT "friend_fk2" FOREIGN KEY ("blocked_by") REFERENCES "user"("id");
