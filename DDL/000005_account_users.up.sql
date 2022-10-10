create table auth.account_users
(
    sns_account_id integer,
    user_id integer,
    user_knd integer
);

comment on table auth.account_users is 'SNS アカウントにアサインされているEMユーザー';
comment on column auth.account_users.sns_account_id is 'SNSアカウントID（sa_id）';
comment on column auth.account_users.user_id is 'ユーザーID';
comment on column auth.account_users.user_knd is 'ユーザー権限
USER_KND_NULL     = 0 // 権限なし
USER_KND_MANAGER  = 1 // マネージャー
USER_KND_OPERATOR = 2 // オペレーター
USER_KND_WRITER   = 3 // ライター
USER_KND_ANALYST  = 4 // アナリスト';

create unique index auth_account_users_uindex
    on auth.account_users (sns_account_id, user_id);
