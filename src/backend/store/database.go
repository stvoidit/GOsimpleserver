package store

import (
	"database/sql"
	"fmt"
	"os"

	// driver
	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
)

// DB - ...
var DB = Store{}

// Config - ...
var Config = config{}

func init() {
	Config.readConfig()
	DB.connect()
}

type config struct {
	host     string
	port     string
	login    string
	password string
	dbname   string
	Secret   []byte
}

func (c *config) readConfig() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	database := cfg.Section("database")
	host, _ := database.GetKey("host")
	port, _ := database.GetKey("port")
	login, _ := database.GetKey("login")
	password, _ := database.GetKey("password")
	dbname, _ := database.GetKey("dbname")
	c.host = host.String()
	c.port = port.String()
	c.login = login.String()
	c.password = password.String()
	c.dbname = dbname.String()

	application := cfg.Section("application")
	secret, err := application.GetKey("secret")
	if err != nil {
		c.Secret = []byte("VeryBadSecretKey")
	}
	c.Secret = []byte(secret.String())
}

// Store - ...
type Store struct {
	*sql.DB
}

// Connect - ....
func (s *Store) connect() {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Config.host, Config.port, Config.login, Config.password, Config.dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	s.DB = db
	s.existsTables()
}

func (s *Store) existsTables() {
	err := s.Ping()
	if err != nil {
		panic(err)
	}
	s.creatreTables()
}

func (s *Store) creatreTables() {
	tableUsers := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS public.users (
		id serial NOT NULL,
		username varchar(50) NOT NULL,
		email varchar(50) NOT NULL,
		"password" varchar(200) NOT NULL,
		"role" int4 NOT NULL,
		CONSTRAINT users_un UNIQUE (username, email)
	);

	CREATE OR REPLACE FUNCTION public.hashpassword(userpassword character varying)
	RETURNS character varying
	LANGUAGE plpgsql
	AS $function$
		declare salt varchar;
		begin
		salt := '%s';
		return md5(concat(salt, userpassword));
	END;
	$function$;

	CREATE OR REPLACE FUNCTION public.insert_user()
	RETURNS trigger
	LANGUAGE plpgsql
	AS $function$
		declare 
			_newpwd varchar;
		begin
		_newpwd := (select hashpassword(new."password"));
		if (TG_OP = 'INSERT') then
			new."password" = _newpwd;
		elseif (TG_OP = 'UPDATE') then
			if (_newpwd = old."password") then
				new."password" = old."password";
			elseif (_newpwd != (select hashpassword(old."password"))) then
				new."password" = _newpwd;
			end if;
		end if;
	return new;
	END;
	$function$;
	
	do 
	$$
	begin
	if (select count(*) from information_schema.triggers) = 0 then
	CREATE TRIGGER pwd before insert or update on public.users for each row execute procedure insert_user();
	end if;
	end;
	$$;`,
		"MySecretKey")
	_, err := s.Exec(tableUsers)
	if err != nil {
		panic(err)
	}

	tableVideos := `CREATE TABLE IF NOT EXISTS public.videos (
		id varchar NOT NULL,
		url varchar NOT NULL,
		active bool NOT NULL DEFAULT true,
		uploaddate varchar NOT NULL,
		channel varchar NULL,
		title varchar NOT NULL,
		CONSTRAINT videos_un UNIQUE (id));
		CREATE UNIQUE INDEX IF NOT EXISTS videos_un ON public.videos USING btree (id);`
	s.Exec(tableVideos)

	tableStatistic := `CREATE TABLE IF NOT EXISTS public.statistic (
		id serial NOT NULL,
		updated timestamptz NOT NULL DEFAULT now(),
		"views" int4 NOT NULL,
		likes int4 NOT NULL,
		dislikes int4 NULL,
		channel varchar NOT NULL,
		channelname varchar NOT NULL,
		followers varchar NOT NULL,
		video varchar NULL,
		CONSTRAINT statistic_fk FOREIGN KEY (video) REFERENCES videos(id) ON UPDATE CASCADE ON DELETE SET NULL);
		CREATE INDEX IF NOT EXISTS statistic_channel_idx ON public.statistic USING btree (channel);`
	s.Exec(tableStatistic)
}
