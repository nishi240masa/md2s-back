create table if not exists user(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name varchar(50),
    icon_url varchar(255),
    google_id int,
);

create table if not exists like(
    id serial primary key,
    user_id uuid,
    article_id int
);


create table if not exists article(
    id serial primary key,
    user_id uuid,
    title varchar(50),
    main_md text,
    slide_md text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
    like_count int default 0
    public boolean default false
    qiita_article boolean default false
);

create table if not exists article_tag_relation(
    id serial primary key,
    article_id int,
    tag_id int
);

create table if not exists tags(
    id serial primary key,
    word varchar(20)
);



