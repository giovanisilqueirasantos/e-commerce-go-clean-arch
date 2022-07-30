CREATE TABLE gocleanarch.auth (
	id INT auto_increment NOT NULL,
	uuid varchar(128) NOT NULL,
	login varchar(150) NOT NULL,
	password varchar(150) NOT NULL,
	CONSTRAINT auth_id_PK PRIMARY KEY (id),
  CONSTRAINT auth_id_UN UNIQUE KEY (id),
  CONSTRAINT auth_uuid_UN UNIQUE KEY (uuid),
  CONSTRAINT auth_login_UN UNIQUE KEY (login)
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;

CREATE TABLE gocleanarch.users (
	id INT auto_increment NOT NULL,
	uuid varchar(128) NOT NULL,
	email varchar(150) NOT NULL,
	first_name varchar(100) NOT NULL,
	last_name varchar(100) NOT NULL,
	phone_number varchar(20) NOT NULL,
	address_city varchar(100) NOT NULL,
	address_state varchar(100) NOT NULL,
	address_neighborhood varchar(150) NOT NULL,
	address_street varchar(150) NOT NULL,
	address_number varchar(20) NOT NULL,
	address_zipcode varchar(100) NOT NULL,
	CONSTRAINT user_id_PK PRIMARY KEY (id),
	CONSTRAINT user_id_UN UNIQUE KEY (id),
	CONSTRAINT user_uuid_UN UNIQUE KEY (uuid),
	CONSTRAINT user_email_UN UNIQUE KEY (email),
	CONSTRAINT user_phone_number_UN UNIQUE KEY (phone_number)
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;

CREATE TABLE gocleanarch.code (
	value varchar(100) NOT NULL,
	identifier varchar(100) NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;
