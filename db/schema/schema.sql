create type product_type as enum (
    'pc98',
    'windows',
    'zuns_music_collection',
    'akyus_untouched_score',
    'commercial_books',
    'other'
);

create table products (
    id            text                     not null primary key,
    name          text                     not null,
    short_name    text                     not null,
    product_type  product_type             not null,
    series_number numeric(5,2)             not null,
    created_at    timestamp with time zone not null default current_timestamp,
    updated_at    timestamp with time zone not null default current_timestamp
);
comment on table  products is '原作';
comment on column products.id is '原作ID';
comment on column products.name is '名前';
comment on column products.short_name is '短い名前';
comment on column products.product_type is '原作種別';
comment on column products.series_number is 'シリーズ番号';
comment on column products.created_at is '作成日時';
comment on column products.updated_at is '更新日時';

create table original_songs (
    id           text                     not null primary key,
    product_id   text                     not null references products(id),
    name         text                     not null,
    composer     text                     not null default '',
    arranger     text                     not null default '',
    track_number integer                  not null,
    is_original  boolean                  not null default false,
    source_id    text                     not null default '',
    created_at   timestamp with time zone not null default current_timestamp,
    updated_at   timestamp with time zone not null default current_timestamp
);
comment on table  original_songs is '原曲';
comment on column original_songs.product_id is '原作ID';
comment on column original_songs.id is '原曲ID';
comment on column original_songs.name is '名前';
comment on column original_songs.composer is '作曲者';
comment on column original_songs.arranger is '編曲者';
comment on column original_songs.track_number is 'トラック番号';
comment on column original_songs.is_original is 'オリジナル有無(true: オリジナル(初出)、false: 再録など)';
comment on column original_songs.source_id is '原曲元の原曲ID';
comment on column original_songs.created_at is '作成日時';
comment on column original_songs.updated_at is '更新日時';
