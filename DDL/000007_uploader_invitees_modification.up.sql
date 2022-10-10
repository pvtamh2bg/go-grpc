
comment on table auth.uploader_invitees is 'アップローダー招待者';
comment on column auth.uploader_invitees.uuid is '招待ID';
comment on column auth.uploader_invitees.uploader_uuid is 'アップローダートークン';
comment on column auth.uploader_invitees.user_id is 'ユーザーID';
comment on column auth.uploader_invitees.status is 'アップローダー招待者・ステータス
  ACTIVE = 1; // 有効
  PAUSE = 2; // 停止中';
comment on column auth.uploader_invitees.created_by is '作成者ID';
comment on column auth.uploader_invitees.created_at is '作成日時';
comment on column auth.uploader_invitees.updated_by is '更新者ID';
comment on column auth.uploader_invitees.updated_at is '更新日時';
comment on column auth.uploader_invitees.deleted_by is '削除者ID';
comment on column auth.uploader_invitees.deleted_at is '削除日時';

create unique index auth_uploader_invitees_id_uindex
    on auth.uploader_invitees (uploader_uuid, user_id);
create index auth_uploader_invitees_uploader_token_index
    on auth.uploader_invitees (uploader_uuid);


