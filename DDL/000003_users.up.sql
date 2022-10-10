create table auth.users
(
    id serial PRIMARY KEY,
    em_user_id integer,
    company_id integer not null,
    name varchar(100),
    email varchar(100),
    icon_id smallint,
    user_type smallint,
    created_by integer not null,
    created_at TIMESTAMP WITH TIME ZONE not null DEFAULT now(),
    updated_by integer not null,
    updated_at TIMESTAMP WITH TIME ZONE not null DEFAULT now(),
    deleted_by integer,
    deleted_at TIMESTAMP WITH TIME ZONE default NULL
);

comment on table auth.users is 'すらぼユーザー';
comment on column auth.users.id is 'ユーザーID';
comment on column auth.users.em_user_id is 'EMユーザーID';
comment on column auth.users.company_id is '企業ID';
comment on column auth.users.name is 'ユーザー名';
comment on column auth.users.email is 'メールアドレス';
comment on column auth.users.icon_id is 'EMユーザーアイコンID';
comment on column auth.users.user_type is 'ユーザータイプ
  USER_TYPE_EM = 1;
  USER_TYPE_INVITEE_POST = 2;
  USER_TYPE_INVITEE_UPLOADER = 3;';
comment on column auth.users.created_by is '作成者ID';
comment on column auth.users.created_at is '作成日時';
comment on column auth.users.deleted_by is '削除者ID';
comment on column auth.users.deleted_at is '削除日時';

create index users_company_id_index
    on auth.users (company_id);
