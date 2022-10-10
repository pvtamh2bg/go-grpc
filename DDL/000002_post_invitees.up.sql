create table auth.post_invitees
(
    uuid uuid PRIMARY KEY
        DEFAULT gen_random_uuid(),
    post_id integer not null,
    user_id integer not null,
    "comment" text,
    passcode char(100),
    created_by integer not null,
    created_at TIMESTAMP WITH TIME ZONE not null DEFAULT now(),
    updated_by integer not null,
    updated_at TIMESTAMP WITH TIME ZONE not null DEFAULT now(),
    deleted_by integer,
    deleted_at TIMESTAMP WITH TIME ZONE default NULL
);

comment on table auth.post_invitees is '投稿作成画面・招待者';
comment on column auth.post_invitees.uuid is '招待ID';
comment on column auth.post_invitees.post_id is '投稿ID';
comment on column auth.post_invitees.user_id is 'ユーザーID';
comment on column auth.post_invitees."comment" is 'コメント';
comment on column auth.post_invitees.passcode is 'パスコード';
comment on column auth.post_invitees.created_by is '作成者ID';
comment on column auth.post_invitees.created_at is '作成日時';
comment on column auth.post_invitees.updated_by is '更新者ID';
comment on column auth.post_invitees.updated_at is '更新日時';
comment on column auth.post_invitees.deleted_by is '削除者ID';
comment on column auth.post_invitees.deleted_at is '削除日時';

create unique index post_invitees_id_uindex
    on auth.post_invitees (post_id, user_id, COALESCE(deleted_at, 'infinity'));
create index invitees_company_id_index
    on auth.post_invitees (post_id);

