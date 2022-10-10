create table auth.uploader_invitees
(
    uuid uuid PRIMARY KEY
        DEFAULT gen_random_uuid(),
    uploader_uuid uuid not null,
    user_id integer not null,
    status smallint,
    created_by integer not null,
    created_at TIMESTAMP WITH TIME ZONE not null DEFAULT now(),
    updated_by integer not null,
    updated_at TIMESTAMP WITH TIME ZONE not null DEFAULT now(),
    deleted_by integer,
    deleted_at TIMESTAMP WITH TIME ZONE default NULL
);

-- comment on table media.uploader_invitees is 'アップローダー招待者';
-- comment on column media.uploader_invitees.uuid is '招待ID';
-- comment on column media.uploader_invitees.uploader_uuid is 'アップローダートークン';
-- comment on column media.uploader_invitees.user_id is 'ユーザーID';
-- comment on column media.uploader_invitees.status is 'アップローダー招待者・ステータス
--   ACTIVE = 1; // 有効
--   PAUSE = 2; // 停止中';
-- comment on column media.uploader_invitees.created_by is '作成者ID';
-- comment on column media.uploader_invitees.created_at is '作成日時';
-- comment on column media.uploader_invitees.updated_by is '更新者ID';
-- comment on column media.uploader_invitees.updated_at is '更新日時';
-- comment on column media.uploader_invitees.deleted_by is '削除者ID';
-- comment on column media.uploader_invitees.deleted_at is '削除日時';
--
-- create unique index media_uploader_invitees_id_uindex
--     on media.uploader_invitees (uploader_uuid, user_id);
-- create index media_uploader_invitees_uploader_token_index
--     on media.uploader_invitees (uploader_uuid);


