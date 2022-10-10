-- #234 投稿作成画面招待者のDDL修正
comment on column auth.users.user_type is 'ユーザータイプ
  USER_TYPE_EM = 1;
  USER_TYPE_INVITEE = 2;';
