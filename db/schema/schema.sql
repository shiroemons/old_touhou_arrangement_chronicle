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

create type distribution_service as enum (
    'spotify',
    'apple_music',
    'youtube_music',
    'line_music',
    'itunes',
    'youtube',
    'nicovideo',
    'sound_cloud'
);

create table product_distribution_service_urls (
    id           text                     not null primary key,
    product_id   text                     not null references products(id),
    service      distribution_service     not null,
    url          text                     not null,
    created_at   timestamp with time zone not null default current_timestamp,
    updated_at   timestamp with time zone not null default current_timestamp
);
comment on table  product_distribution_service_urls is '原作(音楽)配信サービスURL';
comment on column product_distribution_service_urls.product_id is '原作ID';
comment on column product_distribution_service_urls.service is '配信サービス';
comment on column product_distribution_service_urls.url is 'URL';
comment on column product_distribution_service_urls.created_at is '作成日時';
comment on column product_distribution_service_urls.updated_at is '更新日時';

create table original_song_distribution_service_urls (
    id               text                     not null primary key,
    original_song_id text                     not null references original_songs(id),
    service          distribution_service     not null,
    url              text                     not null,
    created_at       timestamp with time zone not null default current_timestamp,
    updated_at       timestamp with time zone not null default current_timestamp
);
comment on table  original_song_distribution_service_urls is '原曲配信サービスURL';
comment on column original_song_distribution_service_urls.original_song_id is '原曲ID';
comment on column original_song_distribution_service_urls.service is '配信サービス';
comment on column original_song_distribution_service_urls.url is 'URL';
comment on column original_song_distribution_service_urls.created_at is '作成日時';
comment on column original_song_distribution_service_urls.updated_at is '更新日時';

create table event_series (
    id           text                     not null primary key,
    name         text                     not null unique,
    created_at   timestamp with time zone not null default current_timestamp,
    updated_at   timestamp with time zone not null default current_timestamp
);
comment on table  event_series is 'イベントシリーズ';
comment on column event_series.name is '名前';
comment on column event_series.created_at is '作成日時';
comment on column event_series.updated_at is '更新日時';

create table events (
    id              text                     not null primary key,
    event_series_id text                     not null references event_series(id),
    name            text                     not null unique,
    created_at      timestamp with time zone not null default current_timestamp,
    updated_at      timestamp with time zone not null default current_timestamp
);
comment on table  events is 'イベント';
comment on column events.event_series_id is 'イベントシリーズID';
comment on column events.name is '名前';
comment on column events.created_at is '作成日時';
comment on column events.updated_at is '更新日時';

create type event_status as enum (
    'scheduled',    -- 開催済み
    'cancelled',    -- 中止
    'postpone',     -- 延期(開催日未定)
    'rescheduled',  -- 延期(開催日決定)
    'moved_online', -- オンライン開催に変更
    'other'         -- その他
);

create type event_format as enum (
    'offline', -- オフライン開催
    'online',  -- オンライン開催
    'mixed'   -- オフライン・オンライン両方開催
);

create table event_details (
    event_id     text                     not null primary key references events(id),
    event_status event_status             not null default 'scheduled'::event_status,
    format       event_format             not null default 'offline'::event_format,
    region_code  text                     not null default 'JP',
    address      text                     not null default '',
    description  text                     not null default '',
    url          text                     not null default '',
    twitter_url  text                     not null default '',
    created_at   timestamp with time zone not null default current_timestamp,
    updated_at   timestamp with time zone not null default current_timestamp
);
comment on table  event_details is 'イベント詳細';
comment on column event_details.event_id is 'イベントID';
comment on column event_details.event_status is 'ステータス/scheduled: 開催済み, cancelled: 中止, postpone: 延期(開催日未定), rescheduled: 延期(開催日決定), moved_online: オンライン開催に変更, other: その他/default: scheduled';
comment on column event_details.format is '形式/offline: オフライン開催, online: オフライン開催, mixed: 両方開催/default: offline';
comment on column event_details.region_code is 'リージョンコード/default: JP';
comment on column event_details.address is '開催場所';
comment on column event_details.description is '説明';
comment on column event_details.url is 'URL';
comment on column event_details.twitter_url is 'Twitter URL';
comment on column event_details.created_at is '作成日時';
comment on column event_details.updated_at is '更新日時';

create table sub_events (
    id         text                     not null primary key,
    event_id   text                     not null references events(id),
    name       text                     not null unique,
    event_date date                     not null unique,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp
);
comment on table  sub_events is 'サブイベント';
comment on column sub_events.event_id is 'イベントID';
comment on column sub_events.name is '名前(例: 〇〇 2日目)';
comment on column sub_events.event_date is '開催日';
comment on column sub_events.created_at is '作成日時';
comment on column sub_events.updated_at is '更新日時';

create table sub_event_details (
    sub_event_id text                     not null primary key references sub_events(id),
    event_status event_status             not null default 'scheduled'::event_status,
    description  text                     not null default '',
    created_at   timestamp with time zone not null default current_timestamp,
    updated_at   timestamp with time zone not null default current_timestamp
);
comment on table  sub_event_details is 'サブイベント詳細';
comment on column sub_event_details.sub_event_id is 'サブイベントID';
comment on column sub_event_details.event_status is 'ステータス/scheduled: 開催済み, cancelled: 中止, postpone: 延期(開催日未定), rescheduled: 延期(開催日決定), moved_online: オンライン開催に変更, other: その他/default: scheduled';
comment on column sub_event_details.description is '説明';
comment on column sub_event_details.created_at is '作成日時';
comment on column sub_event_details.updated_at is '更新日時';

create type initial_letter_type as enum (
    'symbol',   -- 記号
    'number',   -- 数字
    'alphabet', -- 英字
    'kana',     -- かな(ひらがな・カタカナ)
    'kanji',    -- 漢字
    'other'     -- その他
);

create table artists (
    id                    text                     not null primary key,
    name                  text                     not null,
    initial_letter_type   initial_letter_type      not null,
    initial_letter_detail text                     not null,
    created_at            timestamp with time zone not null default current_timestamp,
    updated_at            timestamp with time zone not null default current_timestamp
);
comment on table  artists is 'アーティスト';
comment on column artists.name is '名前';
comment on column artists.initial_letter_type is '頭文字の文字種別(symbol,number,alphabet,kana,kanji,other)';
comment on column artists.initial_letter_detail is '開催日';
comment on column artists.created_at is '作成日時';
comment on column artists.updated_at is '更新日時';
